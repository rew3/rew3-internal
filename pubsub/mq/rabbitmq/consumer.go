package rabbitmq

import (
	"github.com/rew3/rew3-internal/pkg/utils/logger"
	"github.com/rew3/rew3-internal/pubsub/mq/config"
	"github.com/rew3/rew3-internal/pubsub/mq/message"
	"github.com/rew3/rew3-internal/pubsub/mq/types"
)

/**
 * MQConsumer for rabbitmq Message Queue.
 */
type MQConsumer struct {
	types.Publisher
	channel *MQChannel
	props   config.ConsumerProps
	codec   *message.Codec[message.Message, []byte]
}

/**
 * Create new consumer.
 * Initialization will fail in either of one condition: connection error, exchange declare
 * fail, queue declare fail, routing key binding fail.
 *
 * Note: if channel/connection is closed, consumer will be closed. you will have to create
 * new consumer, in order to continue consuming new message from message queue server.
 */
func NewConsumer(connection *MQConnection, props config.ConsumerProps) (types.Consumer, error) {
	channel := NewChannel(false, connection)
	consumer := &MQConsumer{
		channel: channel,
		props:   props,
		codec:   message.DefaultCodec[message.Message](),
	}
	consumer.checkConnection()
	err := consumer.declareExchange()
	if err != nil {
		logger.Log().Errorln("Unable to create consumer: exchange declare failed")
		channel.Close()
		return nil, err
	}
	err = consumer.declareQueue()
	if err != nil {
		logger.Log().Errorln("Unable to create consumer: queue declare failed")
		channel.Close()
		return nil, err
	}
	err = consumer.bindRoutingKeys()
	if err != nil {
		logger.Log().Errorln("Unable to create consumer: routing key binding failed")
		channel.Close()
		return nil, err
	}
	return consumer, nil
}

/**
 * Subscribe to new message from queue.
 */
func (c *MQConsumer) Subscribe() (<-chan message.Message, error) {
	msgs, err := c.channel.GetChannel().Consume(
		"",    // queue
		"",    // consumer
		false, // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // args
	)
	if err != nil {
		logger.Log().Errorln("Error, Unable start consuming: ", err)
		return nil, err
	}
	channel := make(chan message.Message)
	go func() {
		for d := range msgs {
			data := d.Body
			d.Ack(false) // acknowledge message.
			msg, err := c.codec.Deserializer.Deserialize(data)
			if err != nil {
				// no need to handle, ensure, correct message is published.
				logger.Log().Errorln("Unable to de-serialize message: ", err)
			}
			channel <- msg
		}
		close(channel) // close if message receive stop.
	}()
	go func() {
		// when error in message queue channel, close data channel.
		for range c.channel.NotifyErrorClose {
			close(channel)
			return
		}
	}()
	return channel, nil
}

/**
 * Check and wait till channel connection is opened.
 */
func (p *MQConsumer) checkConnection() {
	if !p.channel.IsOpened() && p.channel.AutoReconnect {
		logger.Log().Infoln("Waiting till connection is opened...")
		<-p.channel.NotifyOpened
		logger.Log().Infoln("Channel opened.")
	}
	logger.Log().Infoln("Channel is connected")
}

/**
 * Declare rabbitmq exchange.
 */
func (p *MQConsumer) declareExchange() error {
	if p.channel.GetChannel() == nil {
		logger.Log().Println("Nil channel.....")
	}
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
		logger.Log().Errorln("Unable to declare exchange for consumer: ", err)
		return err
	}
	return nil
}

/**
 * Declare rabbitmq queue.
 */
func (p *MQConsumer) declareQueue() error {
	_, err := p.channel.GetChannel().QueueDeclare(
		"",                     // name
		p.props.IsDurableQueue, // durable
		false,                  // delete when unused
		true,                   // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	if err != nil {
		logger.Log().Errorln("Unable to declare queue: ", err)
		return err
	}
	return nil
}

/**
 * Bind routing keys to queue.
 */
func (p *MQConsumer) bindRoutingKeys() error {
	for _, key := range p.props.RoutingKeys {
		err := p.channel.GetChannel().QueueBind(
			"",                                 // queue name
			string(key),                        // routing key
			string(p.props.ExchangeProps.Name), // exchange
			false,
			nil,
		)
		if err != nil {
			logger.Log().Errorln("Unable to bind keys with queue: ", err)
			return err
		}
	}
	return nil
}
