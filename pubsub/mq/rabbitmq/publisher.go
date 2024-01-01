package rabbitmq

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rew3/rew3-internal/pkg/utils/logger"
	"github.com/rew3/rew3-internal/pubsub/mq/config"
	"github.com/rew3/rew3-internal/pubsub/mq/message"
	"github.com/rew3/rew3-internal/pubsub/mq/types"
)

/**
 * MQPublisher for rabbitmq Message Queue.
 */
type MQPublisher struct {
	types.Publisher
	mutex          sync.Mutex
	channel        *MQChannel
	props          config.PublisherProps
	codec          *message.Codec[message.Message, []byte]
	ackCallbacks   map[types.DeliveryTag]func()
	errorCallbacks map[string]func()
	processedDT    types.DeliveryTag // delivery tag computed for last published message.
	confirmedDT    types.DeliveryTag // delivery tag for message that has been confirmed and successfully published.
	isConfigured   bool
	isStopped      bool
}

/**
 * Create new publisher.
 */
func NewPublisher(connection *MQConnection, props config.PublisherProps) types.Publisher {
	channel := NewChannel(true, connection)
	publisher := &MQPublisher{
		mutex:          sync.Mutex{},
		channel:        channel,
		props:          props,
		codec:          message.DefaultCodec[message.Message](),
		ackCallbacks:   make(map[types.DeliveryTag]func()),
		errorCallbacks: make(map[string]func()),
		processedDT:    0,
		confirmedDT:    0,
		isConfigured:   false,
		isStopped:      false,
	}
	publisher.asyncInit()
	return publisher
}

/**
 * Publish message.
 * Note: make sure to provide routing key, if distribution type require routing key.
 * Please call publish() only in same goroutine where publisher is created, otherwise, there can be
 * possible chance of Data race and deadlock when multiple publish() is called from different goroutines.
 */
func (p *MQPublisher) Publish(key config.RoutingKey, message message.Message) (types.DeliveryTag, error) {
	if p.isStopped {
		return 0, errors.New("publisher is already stopped")
	}
	if !p.isConfigured {
		return 0, errors.New("publisher is is not yet configured, please try again in a while")
	}
	body, err := p.codec.Serializer.Serialize(message)
	if err != nil {
		logger.Log().Errorln("Unable to serialize message: ", err)
		return 0, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	p.mutex.Lock()
	defer p.mutex.Unlock()
	err = p.channel.GetChannel().PublishWithContext(
		ctx,
		string(p.props.ExchangeProps.Name), // exchange
		string(key),                        // routing key
		false,                              // mandatory
		false,                              // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	if err != nil {
		logger.Log().Errorln("Error while publishing message: ", err)
		return 0, err
	}
	p.processedDT = p.processedDT + 1
	return p.processedDT, nil
}

/**
 * Check is message has been delivered and successfully acknowledged or not.
 */
func (p *MQPublisher) IsAck(tag types.DeliveryTag) bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return tag <= p.confirmedDT
}

/**
 * Add listener to invoke when message has been received/acknowledge by server.
 * A message delivery id is provided for the confirmation which will be returned by Publish method.
 * Note: Call back will be executed in another go-routine.
 */
func (p *MQPublisher) OnAck(tag types.DeliveryTag, callback func()) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.ackCallbacks[tag] = callback
}

/**
 * Add listener to invoke when given message id message publish is failed.
 * Note: Call back will be executed in another go-routine.
 */
func (p *MQPublisher) OnError(messageId string, callback func()) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.errorCallbacks[messageId] = callback
}

/**
 * Stop the publisher.
 */
func (c *MQPublisher) Stop() {
	c.mutex.Lock()
	c.channel.Close()
	c.isStopped = true
	c.mutex.Unlock()
	logger.Log().Infoln("Publisher is manually cancelled.")
}

/**
 * Initialize the publisher.
 */
func (p *MQPublisher) asyncInit() {
	if p.channel.IsOpened() {
		p.initListeners()
		go func() { // mocking. for first initialization.
			p.channel.NotifyOpened <- p.channel.channel
		}()
	} else {
		p.initListeners()
	}
}

/**
 * Reset the publisher, clear all ack and error callbacks, reset delivery tags.
 */
func (p *MQPublisher) reset() {
	p.mutex.Lock()
	p.ackCallbacks = make(map[types.DeliveryTag]func())
	p.errorCallbacks = make(map[string]func())
	p.processedDT = 0
	p.confirmedDT = 0
	p.mutex.Unlock()
}

/**
 * Initialize event listeners.
 */
func (p *MQPublisher) initListeners() {
	logger.Log().Infoln("Initializing listeners...")
	go func() {
		// when channel connection opened - new or reconnected.
		for range p.channel.NotifyOpened {
			logger.Log().Infoln("Channel opened. configuring publisher...")
			p.reset()
			p.mutex.Lock()
			isConfigured := p.isConfigured
			p.mutex.Unlock()
			if !isConfigured {
				go p.configureMqSetting(false)
			}

			// Handle publish failed case.
			go func() {
				cn := p.channel.GetChannel().NotifyReturn(make(chan amqp091.Return, 1000))
				for returnErr := range cn {
					msg, _ := p.codec.Deserializer.Deserialize(returnErr.Body)
					cb := p.errorCallbacks[msg.MessageId]
					if cb != nil {
						go cb()
					}
					p.mutex.Lock()
					delete(p.errorCallbacks, msg.MessageId)
					p.mutex.Unlock()
				}
			}()

			// handle acknowledgement confirmation. make sure channel.ConfirmAck() is called.
			go func() {
				cn := p.channel.GetChannel().NotifyPublish(make(chan amqp091.Confirmation, 1000))
				for confirmation := range cn {
					tag := types.DeliveryTag(confirmation.DeliveryTag)
					cb := p.ackCallbacks[tag]
					if cb != nil {
						go cb()
					}
					p.mutex.Lock()
					p.confirmedDT = tag
					delete(p.ackCallbacks, tag)
					p.mutex.Unlock()
				}
			}()
		}
	}()
	go func() {
		// when error in channel, reset publisher.
		for range p.channel.NotifyErrorClose {
			logger.Log().Infoln("Channel Closed on error. Resetting publisher...")
			p.reset()
		}
	}()
	logger.Log().Infoln("Listener initialized.")
}

/**
 * Configure message queue settings.
 * Note: if one of the configuration failed, it will try to reattempt in every 15 seconds untill
 * setting is configured. once completed, isConfigured is set true.
 */
func (c *MQPublisher) configureMqSetting(isExDeclared bool) {
	logger.Log().Infoln("Configuring message queue settings for publisher...")
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.channel.isClosed {
		logger.Log().Infoln("Channel is not already closed. Cancelling publisher configuration...")
		return
	}
	if !c.channel.IsOpened() {
		logger.Log().Infoln("Channel is not opened. Cancelling publisher configuration...")
		return
	}
	err := c.declareExchange()
	if err != nil {
		logger.Log().Errorln("Redeclaring exchange in 15 seconds...")
		time.AfterFunc(time.Duration(15)*time.Second, func() {
			c.configureMqSetting(false)
		})
		return
	}
	c.channel.EnableConfirm()
	c.isConfigured = true
	logger.Log().Infoln("Message queue settings configured.")
	logger.Log().Infoln("Ready to Publish Messages.")
}

/**
 * Declare rabbitmq exchange.
 * Note: make sure channel is already opened.
 */
func (p *MQPublisher) declareExchange() error {
	err := p.channel.GetChannel().ExchangeDeclare(
		string(p.props.ExchangeProps.Name), // name
		string(p.props.ExchangeProps.Type), // type
		p.props.ExchangeProps.IsDurable,    // durable
		false,                              // auto-deleted
		false,                              // internal
		false,                              // no-wait
		nil,                                // arguments
	)
	if err != nil {
		logger.Log().Errorln("Unable to declare exchange for publisher: ", err)
		return err
	}
	return nil
}
