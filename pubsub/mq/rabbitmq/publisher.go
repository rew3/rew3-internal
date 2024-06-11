package rabbitmq

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rew3/rew3-pkg/pubsub/mq/config"
	"github.com/rew3/rew3-pkg/pubsub/mq/message"
	"github.com/rew3/rew3-pkg/pubsub/mq/types"
	"github.com/rew3/rew3-pkg/utils/logger"
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
	isConfiguring  bool
	isConfigured   bool
	isStopped      bool
}

/**
 * Create new publisher.
 */
func NewPublisher(connection *MQConnection, props config.PublisherProps) types.Publisher {
	channel := NewChannel(string(props.ExchangeProps.Name), true, connection)
	publisher := &MQPublisher{
		mutex:          sync.Mutex{},
		channel:        channel,
		props:          props,
		codec:          message.DefaultCodec[message.Message](),
		ackCallbacks:   make(map[types.DeliveryTag]func()),
		errorCallbacks: make(map[string]func()),
		processedDT:    0,
		confirmedDT:    0,
		isConfiguring:  false,
		isConfigured:   false,
		isStopped:      false,
	}
	publisher.initializePublisher()
	return publisher
}

/**
 * Publish message.
 * Note: make sure to provide routing key, if distribution type require routing key.
 * Please call publish() only in same goroutine where publisher is created, otherwise, there can be
 * possible chance of Data race and deadlock when multiple publish() is called from different goroutines.
 */
func (p *MQPublisher) Publish(key config.RoutingKey, message message.Message) (types.DeliveryTag, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.isStopped {
		return 0, errors.New("publisher is already stopped")
	}
	if !p.isConfigured {
		return 0, errors.New("publisher is is not yet configured, please try again in a while")
	}
	body, err := p.codec.Serializer.Serialize(message)
	if err != nil {
		logger.Log().Errorln(p.logMsg("Unable to serialize message: "), err)
		return 0, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
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
		logger.Log().Errorln(p.logMsg("Error while publishing message: "), err)
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
 * If publisher stopped due to error or manual invocation, this callback will not be invoked.
 */
func (p *MQPublisher) OnAck(tag types.DeliveryTag, callback func()) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.ackCallbacks[tag] = callback
}

/**
 * Add listener to invoke when given message id message publish is failed.
 * Note: Call back will be executed in another go-routine.
 * If publisher stopped due to error or manual invocation, this callback will not be invoked.
 */
func (p *MQPublisher) OnError(messageId string, callback func()) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.errorCallbacks[messageId] = callback
}

/**
 * Stop the publisher.
 */
func (p *MQPublisher) Stop() {
	p.mutex.Lock()
	p.channel.Close()
	p.isStopped = true
	p.mutex.Unlock()
	logger.Log().Infoln(p.logMsg("Publisher is manually cancelled."))
}

/**
 * Reset the publisher, clear all ack and error callbacks, reset delivery tags.
 */
func (p *MQPublisher) reset() {
	p.mutex.Lock()
	p.isConfiguring = false
	p.isConfigured = false
	p.ackCallbacks = make(map[types.DeliveryTag]func())
	p.errorCallbacks = make(map[string]func())
	p.processedDT = 0
	p.confirmedDT = 0
	p.mutex.Unlock()
}

/**
 * Initialize publisher.
 */
func (p *MQPublisher) initializePublisher() {
	logger.Log().Infoln(p.logMsg("Initializing publisher..."))
	go func() {
		// when error in channel, reset publisher.
		for range p.channel.NotifyErrorClose(make(chan bool, 1)) {
			logger.Log().Infoln(p.logMsg("Channel Closed on error. Resetting publisher..."))
			p.reset()
		}
	}()
	configuringFn := func() {
		p.mutex.Lock()
		if p.isConfiguring || p.isConfigured {
			logger.Log().Infoln(p.logMsg("Publisher is configuring or already configured."))
			p.mutex.Unlock()
			return
		}
		p.isConfiguring = true
		p.mutex.Unlock()
		logger.Log().Infoln(p.logMsg("Configuring settings..."))
		notifyChannel := make(chan bool, 1)
		p.configureMqSetting(false, notifyChannel)
		go func() {
			logger.Log().Infoln(p.logMsg("Enabling publish failed messages handling..."))
			// Handle message publish failed case.
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
		go func() {
			logger.Log().Infoln(p.logMsg("Enabling published messages ack handling. Note: Make sure EnableConfirm() is called in Channel)"))
			// handle acknowledgement confirmation. make sure channel.ConfirmAck() is called.
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
		<-notifyChannel
		close(notifyChannel)
		p.mutex.Lock()
		p.isConfiguring = false
		p.isConfigured = true
		p.mutex.Unlock()
		logger.Log().Infoln(p.logMsg("Settings configured."))
	}
	go func() {
		for range p.channel.NotifyReady(make(chan bool, 1)) {
			logger.Log().Infoln(p.logMsg("Channel opened. Configuring publisher..."))
			go configuringFn()
		}
	}()
	if p.channel.IsOpened() {
		go configuringFn()
	}
}

/**
 * Configure message queue settings.
 * Note: if one of the configuration failed, it will try to reattempt in every 15 seconds untill
 * setting is configured. once completed, isConfigured is set true.
 */
func (p *MQPublisher) configureMqSetting(isExDeclared bool, notifyConfigured chan bool) {
	logger.Log().Infoln(p.logMsg("Configuring message queue settings for publisher..."))
	if !p.channel.IsOpened() {
		logger.Log().Infoln(p.logMsg("Channel is not opened. Cancelling publisher configuration..."))
		return
	}
	err := p.declareExchange()
	if err != nil {
		logger.Log().Errorln(p.logMsg("Redeclaring exchange in 15 seconds..."))
		time.AfterFunc(time.Duration(15)*time.Second, func() {
			p.configureMqSetting(false, notifyConfigured)
		})
		return
	}
	p.channel.EnableConfirm()
	logger.Log().Infoln(p.logMsg("Configuration of message queue setting completed."))
	notifyConfigured <- true
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
		logger.Log().Errorln(p.logMsg("Unable to declare exchange for publisher: "), err)
		return err
	}
	return nil
}

// create log message
func (p *MQPublisher) logMsg(msg string) string {
	return "[MQ Publisher: " + string(p.props.ExchangeProps.Name) + "] " + msg
}
