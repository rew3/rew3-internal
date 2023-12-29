package rabbitmq

import (
	"sync"
	"time"

	"github.com/rew3/rew3-internal/pkg/utils/logger"
	"github.com/rew3/rew3-internal/pubsub/mq/config"
	"github.com/rew3/rew3-internal/pubsub/mq/message"
	"github.com/rew3/rew3-internal/pubsub/mq/types"
)

/**
 * MQConsumer for rabbitmq Message Queue.
 * This consumer is asynchronous and is self error handling with auto connect and retries
 * in case of channel/consumer error. Only, calling Cancel() will stop this consumer completely.
 *
 * Note: Messages are automatically acknowledged when received and serialized.
 */
type MQConsumer struct {
	types.Consumer
	mutex        sync.Mutex
	channel      *MQChannel
	props        config.ConsumerProps
	codec        *message.Codec[message.Message, []byte]
	subscribers  []chan<- message.Message
	isConfigured bool
	isCancelled  bool
	isConsuming  bool
}

/*
 * Create new consumer.
 */
func NewConsumer(connection *MQConnection, props config.ConsumerProps) types.Consumer {
	channel := NewChannel(true, connection)
	consumer := &MQConsumer{
		mutex:        sync.Mutex{},
		channel:      channel,
		props:        props,
		codec:        message.DefaultCodec[message.Message](),
		subscribers:  []chan<- message.Message{},
		isConfigured: false,
		isCancelled:  false,
		isConsuming:  false,
	}
	consumer.asyncInit()
	return consumer
}

/**
 * Subscribe to new message from queue.
 * If consumer is already cancelled, subscribe will have no effect (do nothing).
 * bufSize: Default buffer size for channel. so make sure when receiving message, it is handled
 * properly. If buffer size, exceed, deadlock will occur.
 * Note: you can subscribe multiple times using same consumer.
 */
func (c *MQConsumer) Subscribe(bufSize int) <-chan message.Message {
	if c.isCancelled {
		return nil
	}
	ch := make(chan message.Message, 100)
	c.mutex.Lock()
	c.subscribers = append(c.subscribers, ch)
	c.mutex.Unlock()
	return ch
}

/**
 * Cancel all the subscriptions in this consumer.
 * Cancelled consumer cannot be re-subscribed again. In such case, re-start new consumer and subscribe.
 */
func (c *MQConsumer) Cancel() {
	c.mutex.Lock()
	c.channel.Close()
	c.isCancelled = true
	c.isConsuming = false
	for _, s := range c.subscribers {
		close(s)
	}
	c.mutex.Unlock()
	logger.Log().Infoln("Consumer is manually cancelled.")
}

/**
 * Initialize the consumer asynchrously.
 */
func (p *MQConsumer) asyncInit() {
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
 * Initialize listeners for consumer configuration and startup.
 */
func (p *MQConsumer) initListeners() {
	go func() {
		for range p.channel.NotifyOpened {
			if !p.isConfigured {
				logger.Log().Infoln("Configuring consumer...")
				ch := make(chan bool, 1)
				p.configureMqSetting(false, false, false, ch)
				<-ch // configured.
				close(ch)
				logger.Log().Infoln("Consumer configured.")
				logger.Log().Infoln("Starting consumer...")
				if !p.isConsuming {
					p.startConsume()
				}
			} else {
				logger.Log().Infoln("Starting consumer...")
				if !p.isConsuming {
					p.startConsume()
				}
			}
		}
	}()
}

/**
 * Configure message queue settings.
 * Declare exchange, queue and bind routing keys.
 * Note: if one of the configuration failed, it will try to reattempt in every 15 seconds untill all
 * settings are configured. once completed, isConfigured is set true.
 */
func (c *MQConsumer) configureMqSetting(isExDeclared, isQueueDec, isRKBinded bool, ch chan<- bool) {
	logger.Log().Infoln("Configuring message queue settings...")
	c.mutex.Lock()
	defer c.mutex.Unlock()
	err := c.declareExchange()
	if err != nil {
		logger.Log().Errorln("Redeclaring exchange in 15 seconds...")
		time.AfterFunc(time.Duration(15)*time.Second, func() {
			c.configureMqSetting(false, isQueueDec, isRKBinded, ch)
		})
		return
	}
	err = c.declareQueue()
	if err != nil {
		logger.Log().Errorln("Redeclaring queue in 15 seconds...")
		time.AfterFunc(time.Duration(15)*time.Second, func() {
			c.configureMqSetting(isExDeclared, false, isRKBinded, ch)
		})
		return
	}
	err = c.bindRoutingKeys()
	if err != nil {
		logger.Log().Errorln("Binding routing keys in 15 seconds...")
		time.AfterFunc(time.Duration(15)*time.Second, func() {
			c.configureMqSetting(isExDeclared, isQueueDec, false, ch)
		})
		return
	}
	c.isConfigured = true
	logger.Log().Infoln("Configuration of message queue setting completed.")
	ch <- true // all configured, notifying.
}

/**
 * Declare rabbitmq exchange.
 */
func (p *MQConsumer) declareExchange() error {
	logger.Log().Infoln("Declaring exchange...")
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
	logger.Log().Infoln("Declaring queue...")
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
	logger.Log().Infoln("Binding routing keys...")
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

/**
 * Start consuming new messages.
 * If there is error will auto start consuming in 15 seconds.
 */
func (c *MQConsumer) startConsume() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
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
		logger.Log().Errorln("Restarting consumer in 15 seconds...")
		time.AfterFunc(time.Duration(15)*time.Second, func() {
			c.startConsume()
		})
		return
	}
	c.isConsuming = true
	go func() {
		for d := range msgs {
			data := d.Body
			d.Ack(false) // acknowledge message.
			msg, err := c.codec.Deserializer.Deserialize(data)
			if err != nil {
				// no need to handle, ensure, correct message is published.
				logger.Log().Errorln("Unable to de-serialize message: ", err)
			}
			for _, susbcriber := range c.subscribers {
				susbcriber <- msg // publishing to each registered subscriber.
			}
		}
		// consumer stopped.
		c.mutex.Lock()
		c.isConsuming = false
		c.mutex.Unlock()
	}()
}
