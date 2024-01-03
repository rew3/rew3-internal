package rabbitmq

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rew3/rew3-internal/pkg/utils/logger"
	"github.com/rew3/rew3-internal/pubsub/mq/config"
	"github.com/rew3/rew3-internal/pubsub/mq/message"
	"github.com/rew3/rew3-internal/pubsub/mq/types"
)

/**
 * MQConsumer for rabbitmq Message Queue.
 * This consumer is asynchronous and is self error handling with auto connect and retries
 * in case of channel/consumer error. Only, calling Cancel() will stop this consumer completely.
 */
type MQConsumer struct {
	types.Consumer
	mutex         sync.Mutex
	channel       *MQChannel
	props         config.ConsumerProps
	codec         *message.Codec[message.Message, []byte]
	subscribers   []*subscriber
	isConfigured  bool
	isConfiguring bool
	isCancelled   bool
	isConsuming   bool
	queueName     string
}

/**
 * Message Subscriber
 */
type subscriber struct {
	isAutoAck  bool                 // Flag if message should be auto acknowledge or manually.
	msgChannel chan message.Message // Channel to receive message.
	ackChannel chan bool            // If manual ack, acknowledge via this channel by publishing true or false.
}

/**
 * Consumer groups.
 * Used to manage multiple consumers using same channel and connection. This is useful to allow
 * consuming multiple messages, using same channel connection. In addition, using this, we can gracefully
 * close multiple consumers at once.
 */
type MQConsumerGroup struct {
	types.ConsumerGroup
	mutex     sync.Mutex
	channel   *MQChannel
	consumers []*MQConsumer
}

/*
 * Create new consumer group.
 */
func NewConsumerGroup(connection *MQConnection) *MQConsumerGroup {
	channel := NewChannel(true, connection)
	return &MQConsumerGroup{
		mutex:     sync.Mutex{},
		channel:   channel,
		consumers: []*MQConsumer{},
	}
}

/**
 * Create new consumer using consumer group.
 */
func (cg *MQConsumerGroup) NewConsumer(props config.ConsumerProps) types.Consumer {
	consumer := &MQConsumer{
		mutex:         sync.Mutex{},
		channel:       cg.channel,
		props:         props,
		codec:         message.DefaultCodec[message.Message](),
		subscribers:   []*subscriber{},
		isConfigured:  false,
		isConfiguring: false,
		isCancelled:   false,
		isConsuming:   false,
		queueName:     props.Name + "_" + uuid.NewString(),
	}
	cg.mutex.Lock()
	cg.consumers = append(cg.consumers, consumer)
	cg.mutex.Unlock()
	consumer.asyncInit()
	return consumer
}

/**
 * Cancel all consumers of this consumer group.
 */
func (cg *MQConsumerGroup) Cancel() {
	logger.Log().Infoln("Cancelling all consumers of consuemr group...")
	for _, c := range cg.consumers {
		c.Cancel()
	}
	logger.Log().Infoln("All consumer group's consumers are cancelled.")
}

/*
 * Create new consumer.
 */
func NewConsumer(connection *MQConnection, props config.ConsumerProps) types.Consumer {
	channel := NewChannel(true, connection)
	consumer := &MQConsumer{
		mutex:         sync.Mutex{},
		channel:       channel,
		props:         props,
		codec:         message.DefaultCodec[message.Message](),
		subscribers:   []*subscriber{},
		isConfigured:  false,
		isConfiguring: false,
		isCancelled:   false,
		isConsuming:   false,
		queueName:     props.Name + "_" + uuid.NewString(),
	}
	consumer.asyncInit()
	return consumer
}

/**
 * Subscribe to new message from queue.
 * If consumer is already cancelled, subscribe will have no effect (do nothing).
 *
 * bufSize: Default buffer size for channel. provide valid size, to avoid deadlock.
 * isAutoAck: if true, message acknowledge automatically.
 *
 * Returns: (messageChannel, ackChannel)
 * messageChannel:  receive main message from consumer. if subscriber is cancelled, it is nil.
 * ackChannel: channel to confirm acknowledgement of message (use only if isAutoAck is true).
 *
 * Note: you can subscribe multiple times using same consumer.
 */
func (c *MQConsumer) Subscribe(bufSize int, isAutoAck bool) (<-chan message.Message, chan<- bool) {
	if c.isCancelled {
		return nil, nil
	}
	c.mutex.Lock()
	ch := make(chan message.Message, bufSize)
	ackCh := make(chan bool, 20)
	subscriber := &subscriber{
		isAutoAck:  isAutoAck,
		msgChannel: ch,
		ackChannel: ackCh,
	}
	c.subscribers = append(c.subscribers, subscriber)
	c.mutex.Unlock()
	return ch, ackCh
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
		close(s.msgChannel)
		close(s.ackChannel)
	}
	c.subscribers = []*subscriber{}
	c.mutex.Unlock()
	logger.Log().Infoln("Consumer is manually cancelled.")
}

/**
 * Initialize the consumer asynchrously.
 */
func (c *MQConsumer) asyncInit() {
	logger.Log().Infoln("Initializing Consumer...")
	if c.channel.IsOpened() {
		c.initListeners()
		go func() { // mocking. for first initialization.
			c.channel.NotifyOpened <- c.channel.channel
		}()
	} else {
		c.initListeners()
	}
}

/**
 * Initialize listeners for consumer configuration and startup.
 */
func (c *MQConsumer) initListeners() {
	logger.Log().Infoln("Setting up Consumer event listeners...")
	go func() {
		for range c.channel.NotifyOpened {
			logger.Log().Infoln("Channel opened event received by Consumer.")
			c.mutex.Lock()
			isConfigured := c.isConfigured
			c.mutex.Unlock()
			if !isConfigured {
				logger.Log().Infoln("Configuring consumer...")
				ch := make(chan bool, 1)
				c.configureMqSetting(false, false, false, ch)
				go func() {
					<-ch // configured.
					close(ch)
					logger.Log().Infoln("Consumer configured.")
					logger.Log().Infoln("Starting consumer...")

					c.mutex.Lock() // To avoid: WARNING: DATA RACE
					isConsuming := c.isConsuming
					c.mutex.Unlock()
					if !isConsuming {
						c.startConsume()
					}
				}()
			} else {
				logger.Log().Infoln("Starting consumer...")
				c.mutex.Lock() // To avoid: WARNING: DATA RACE
				isConsuming := c.isConsuming
				c.mutex.Unlock()
				if !isConsuming {
					c.startConsume()
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
	logger.Log().Infoln("Configuring message queue settings for consumer...")
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.isConfiguring {
		logger.Log().Infoln("Consumer is being already being configured...")
		return
	}
	if c.channel.isClosed {
		logger.Log().Infoln("Channel is not already closed. Cancelling consumer configuration...")
		return
	}
	if !c.channel.IsOpened() {
		logger.Log().Infoln("Channel is not opened. Cancelling consumer configuration...")
		return
	}
	c.isConfiguring = true
	err := c.declareExchange()
	if err != nil {
		logger.Log().Errorln("Redeclaring exchange in 15 seconds...")
		c.isConfiguring = false
		time.AfterFunc(time.Duration(15)*time.Second, func() {
			c.configureMqSetting(false, isQueueDec, isRKBinded, ch)
		})
		return
	}
	err = c.declareQueue()
	if err != nil {
		logger.Log().Errorln("Redeclaring queue in 15 seconds...")
		c.isConfiguring = false
		time.AfterFunc(time.Duration(15)*time.Second, func() {
			c.configureMqSetting(isExDeclared, false, isRKBinded, ch)
		})
		return
	}
	err = c.bindRoutingKeys()
	if err != nil {
		logger.Log().Errorln("Binding routing keys in 15 seconds...")
		c.isConfiguring = false
		time.AfterFunc(time.Duration(15)*time.Second, func() {
			c.configureMqSetting(isExDeclared, isQueueDec, false, ch)
		})
		return
	}
	c.isConfigured = true
	c.isConfiguring = false
	logger.Log().Infoln("Configuration of message queue setting completed.")
	ch <- true // all configured, notifying.
}

/**
 * Declare rabbitmq exchange.
 * Note: make sure channel is already opened.
 */
func (c *MQConsumer) declareExchange() error {
	logger.Log().Infoln("Declaring exchange...")
	err := c.channel.GetChannel().ExchangeDeclare(
		string(c.props.ExchangeProps.Name), // name
		string(c.props.ExchangeProps.Type), // type
		c.props.ExchangeProps.IsDurable,    // durable
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
 * Note: make sure channel is already opened.
 */
func (c *MQConsumer) declareQueue() error {
	logger.Log().Infoln("Declaring queue...")
	_, err := c.channel.GetChannel().QueueDeclare(
		c.queueName,            // name
		c.props.IsDurableQueue, // durable
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
 * Note: make sure channel is already opened.
 */
func (c *MQConsumer) bindRoutingKeys() error {
	logger.Log().Infoln("Binding routing keys...")
	for _, key := range c.props.RoutingKeys {
		err := c.channel.GetChannel().QueueBind(
			c.queueName,                        // queue name
			string(key),                        // routing key
			string(c.props.ExchangeProps.Name), // exchange
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
	msgs, err := c.channel.GetChannel().Consume(
		c.queueName, // queue
		"",          // consumer
		false,       // auto ack
		false,       // exclusive
		false,       // no local
		false,       // no wait
		nil,         // args
	)
	if err != nil {
		logger.Log().Errorln("Error, Unable start consuming: ", err)
		logger.Log().Errorln("Restarting consumer in 15 seconds...")
		c.mutex.Unlock()
		time.AfterFunc(time.Duration(15)*time.Second, func() {
			c.startConsume()
		})
		return
	}
	c.isConsuming = true
	c.mutex.Unlock()
	go func() {
		for d := range msgs {
			data := d.Body
			msg, err := c.codec.Deserializer.Deserialize(data)
			if err != nil {
				// no need to handle, ensure, correct message is published.
				logger.Log().Errorln("Unable to de-serialize message: ", err)
			}
			for _, sc := range c.subscribers {
				go func(s *subscriber, dv amqp091.Delivery) {
					s.msgChannel <- msg // publishing to each registered subscriber.
					if s.isAutoAck {
						dv.Ack(false) // acknowledge message.
					} else {
						if <-s.ackChannel {
							dv.Ack(false) // acknowledge message.
						}
					}
				}(sc, d)
			}
		}
		// if this is reached, that means, consumer is stopped.
		logger.Log().Infoln("Consumer stopped receiving message.")
		c.mutex.Lock()
		c.isConsuming = false
		c.mutex.Unlock()
	}()
	logger.Log().Infoln("Consumer Started.")
}
