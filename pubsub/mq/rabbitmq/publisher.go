package rabbitmq

import (
	"context"
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
	pendingDT      types.DeliveryTag // delivery tag computed for last published message.
	confirmedDT    types.DeliveryTag // delivery tag for message that has been confirmed and successfully published.
}

/**
 * Create new publisher.
 */
func NewPublisher(connection *MQConnection, props config.PublisherProps) (types.Publisher, error) {
	channel := NewChannel(true, connection)
	publisher := &MQPublisher{
		mutex:          sync.Mutex{},
		channel:        channel,
		props:          props,
		codec:          message.DefaultCodec[message.Message](),
		ackCallbacks:   make(map[types.DeliveryTag]func()),
		errorCallbacks: make(map[string]func()),
		pendingDT:      0,
		confirmedDT:    0,
	}
	publisher.checkConnection()
	err := publisher.declareExchange()
	if err != nil {
		logger.Log().Errorln("Unable to create publisher: exchange declare failed")
		channel.Close()
		return nil, err
	}
	return publisher, nil
}

/**
 * Publish message.
 * Note: make sure to provide routing key, if distribution type require routing key.
 */
func (p *MQPublisher) Publish(key config.RoutingKey, message message.Message) (types.DeliveryTag, error) {
	p.checkConnection()
	body, err := p.codec.Serializer.Serialize(message)
	if err != nil {
		logger.Log().Errorln("Unable to serialize message: ", err)
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
		logger.Log().Errorln("Error while publishing message: ", err)
		return 0, err
	}
	p.mutex.Lock()
	p.pendingDT = p.pendingDT + 1
	p.mutex.Unlock()
	return types.DeliveryTag(p.pendingDT), nil
}

/**
 * Check is message has been delivered and successfully acknowledged or not.
 */
func (p *MQPublisher) IsAck(tag types.DeliveryTag) bool {
	return tag <= p.confirmedDT
}

/**
 * Add listener to invoke when message has been received/acknowledge by server.
 * A message delivery id is provided for the confirmation which will be returned by Publish method.
 * Note: Call back will be executed in another go-routine.
 */
func (p *MQPublisher) OnAck(tag types.DeliveryTag, callback func()) {
	p.ackCallbacks[tag] = callback
}

/**
 * Add listener to invoke when given message id message publish is failed.
 * Note: Call back will be executed in another go-routine.
 */
func (p *MQPublisher) OnError(messageId string, callback func()) {
	p.errorCallbacks[messageId] = callback
}

/**
 * Check and wait till channel connection is opened.
 */
func (p *MQPublisher) checkConnection() {
	if !p.channel.IsOpened() && p.channel.AutoReconnect {
		logger.Log().Infoln("Waiting till connection is opened...")
		<-p.channel.NotifyOpened
		logger.Log().Infoln("Channel opened.")
		p.mutex.Lock()
		p.pendingDT = 0
		p.confirmedDT = 0
		p.mutex.Unlock()
		p.initListeners() // will have to reinitialize after channel/connection is closed.
	}
}

/**
 * Initialize event listeners.
 */
func (p *MQPublisher) initListeners() {
	logger.Log().Infoln("initializing listeners...")
	go func() {
		// Handle publish failed case.
		cn := p.channel.GetChannel().NotifyReturn(make(chan amqp091.Return, 1000))
		for returnErr := range cn {
			msg, _ := p.codec.Deserializer.Deserialize(returnErr.Body)
			cb := p.errorCallbacks[msg.MessageId]
			if cb != nil {
				cb()
			}
			p.mutex.Lock()
			delete(p.errorCallbacks, msg.MessageId)
			p.mutex.Unlock()
		}
	}()
	if p.channel.PublisherConfirmMode {
		// handle acknowledgement confirmation.
		go func() {
			cn := p.channel.GetChannel().NotifyPublish(make(chan amqp091.Confirmation, 1000))
			for confirmation := range cn {
				tag := types.DeliveryTag(confirmation.DeliveryTag)
				cb := p.ackCallbacks[tag]
				if cb != nil {
					cb()
				}
				p.mutex.Lock()
				p.confirmedDT = tag
				delete(p.ackCallbacks, tag)
				p.mutex.Unlock()
			}
		}()
	}
}

/**
 * Declare rabbitmq exchange.
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
