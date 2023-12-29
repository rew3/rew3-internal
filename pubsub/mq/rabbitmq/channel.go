package rabbitmq

import (
	"sync"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rew3/rew3-internal/pkg/utils/logger"
)

/**
 * MQChannel for Message Queue.
 * MQChannel is used by publisher and consumer to process messages.
 * Note: if auto reconnect enabled, in case of disgraceful close, it
 * will try to re-open in every 10 seconds.
 * channel is fault tolreant????
 */
type MQChannel struct {
	mutex                    sync.Mutex
	connection               *MQConnection
	channel                  *amqp091.Channel
	isClosed                 bool
	autoReconnect            bool                  // in case of connection error, reopen channel automatically
	NotifyOpened             chan *amqp091.Channel // go channel to notify for new opened channel (e.g re-opened)
	NotifyErrorClose         chan bool             // go channel to notify when channel is closed by error.
	registeredCloseNotifiers []chan bool
}

/**
 * Create new channel.
 */
func NewChannel(autoReconnect bool, connection *MQConnection) *MQChannel {
	channel := &MQChannel{
		mutex:                    sync.Mutex{},
		connection:               connection,
		channel:                  nil,
		autoReconnect:            autoReconnect,
		NotifyOpened:             make(chan *amqp091.Channel, 1),
		NotifyErrorClose:         make(chan bool, 1),
		registeredCloseNotifiers: []chan bool{},
		isClosed:                 false,
	}
	// enable auto channel reconnect.
	go channel.enableAutoReconnect()
	go func() {
		// When connection is ready, open channel. useful for re open channel
		// when connection is reconnected.
		for range connection.NotifyConnected() {
			logger.Log().Infoln("Received Connection Connected event.")
			if channel.isClosed {
				logger.Log().Infoln("Channel is already closed, stop listening to connection notification")
				return
			}
			channel.openChannel()
		}
	}()
	// initially, try to open channel is connection is opened.
	if connection.IsConnected() {
		channel.openChannel()
	}
	return channel
}

/*
 * Open Channel.
 * Note: if error occured while opening channel, it will be retried after 10 seconds
 * and new channel is opened. In addition, when closed, you cannot reopen the channel,
 * instead create new.
 */
func (c *MQChannel) openChannel() {
	logger.Log().Infoln("Opening channel...")
	if c.IsOpened() || !c.connection.IsConnected() {
		// Channel already opened and running
		// OR connection is not connected or ready yet.
		logger.Log().Infoln("Channel already opened")
		return
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	ch, err := c.connection.Connection().Channel()
	if err != nil {
		logger.Log().Errorln("Unable to open channel: ", err)
		logger.Log().Errorln("Retrying in 15 seconds...")
		time.AfterFunc(time.Duration(10)*time.Second, func() {
			c.openChannel()
		})
		return
	}
	c.channel = ch
	logger.Log().Infoln("Channel opened")
	c.NotifyOpened <- c.channel
}

/**
 * Enable auto reconnect to channel in case of channel is closed unexpectedly.
 */
func (c *MQChannel) enableAutoReconnect() {
	if c.autoReconnect {
		go func() {
			// when channel is opened, handle auto channel open.
			for ch := range c.NotifyOpened {
				c.mutex.Lock()
				errorChannel := ch.NotifyClose(make(chan *amqp091.Error, 1))
				c.mutex.Unlock()
				for err := range errorChannel {
					logger.Log().Error("Channel closed unexpectedly: ", err)
					logger.Log().Info("Re-opening channel...")
					c.NotifyErrorClose <- true
					close(errorChannel)
					if !c.connection.IsConnected() {
						// seems like main connection is closed, dont do anything.
						// if auto connect enabled in connection, this channel will be initialized again.
						return
					} else {
						time.AfterFunc(time.Duration(10)*time.Second, func() {
							c.openChannel()
						})
					}
				}
			}
		}()
	}
}

/**
 * Use this to enable confirmation ack by server so that, client will be notified
 * of successful message is received by sever. Used by publisher.
 */
func (c *MQChannel) EnableConfirm() {
	err := c.channel.Confirm(false)
	if err != nil {
		logger.Log().Errorln("Unable to enable confirmation ack")
	}
}

/**
 * Get opened channel.
 * Note: if channel is not opened yet, nil is returned.
 * Alternatively, You can listen to NotifyOpened channel to know when channel is opened and is
 * ready to process messages. Also, if there are errors, channel will be retried continously to
 * be opened and notified by NotifyOpened when opened.
 */
func (c *MQChannel) GetChannel() *amqp091.Channel {
	return c.channel
}

/**
 * Check is channel is opened or not.
 */
func (c *MQChannel) IsOpened() bool {
	return c.channel != nil && !c.channel.IsClosed()
}

/**
 * Close channel.
 * Note: when closed, you cannot/dont reopen the channel.
 */
func (c *MQChannel) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
	for _, ch := range c.registeredCloseNotifiers {
		ch <- true // notify that channel is closed to each registered close notifier.
	}
	c.mutex.Lock()
	c.isClosed = true
	c.mutex.Unlock()
	defer close(c.NotifyOpened)
	defer close(c.NotifyErrorClose)
	logger.Log().Errorln("Channel is manually closed")
}

/**
 * Notify when channel is manually is closed by calling Close().
 * Close this returned channel if required.
 */
func (c *MQChannel) NotifyWhenClose() chan bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	ch := make(chan bool, 1)
	c.registeredCloseNotifiers = append(c.registeredCloseNotifiers, ch)
	return ch
}