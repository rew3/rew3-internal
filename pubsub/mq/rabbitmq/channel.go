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
 * Note: This channel is async by nature and is self handling in case of error/exception and
 * re-connection automatically in scheduled time of every 15 seconds.
 *
 * By nature, channel should not be accessed from multiple threads/goroutine, but this channel
 * can be access from another goroutine as well, but the access is synced while publishing.
 */
type MQChannel struct {
	mutex                    sync.Mutex
	name                     string
	connection               *MQConnection
	channel                  *amqp091.Channel
	isClosed                 bool
	isConnecting             bool
	autoReconnect            bool                  // in case of connection error, reopen channel automatically
	NotifyOpened             chan *amqp091.Channel // go channel to notify for new opened channel (e.g re-opened)
	internalNotifyOpened     chan *amqp091.Channel // go channel to notify for new opened channel (e.g re-opened)
	NotifyErrorClose         chan bool             // go channel to notify when channel is closed by error.
	registeredCloseNotifiers []chan bool
}

/**
 * Create new channel.
 */
func NewChannel(name string, autoReconnect bool, connection *MQConnection) *MQChannel {
	channel := &MQChannel{
		mutex:                    sync.Mutex{},
		name:                     name,
		connection:               connection,
		channel:                  nil,
		autoReconnect:            autoReconnect,
		isConnecting:             false,
		NotifyOpened:             make(chan *amqp091.Channel, 1),
		internalNotifyOpened:     make(chan *amqp091.Channel, 1),
		NotifyErrorClose:         make(chan bool, 1),
		registeredCloseNotifiers: []chan bool{},
		isClosed:                 false,
	}
	channel.enableAutoReconnect()
	go func() {
		// When connection is ready, open channel. useful for re open the
		// channel when connection is reconnected.
		for range connection.NotifyConnected() {
			channel.mutex.Lock()
			isConnecting := channel.isConnecting
			channel.mutex.Unlock()
			if isConnecting {
				return
			}
			logger.Log().Infoln("[", name, "]", "Connection established, opening the channel...")
			if channel.isClosed {
				logger.Log().Infoln("[", name, "]", "Channel is already closed, exitting.")
				return
			}
			if !channel.IsOpened() {
				logger.Log().Infoln("[", name, "]", "Channel was closed, opening...")
				channel.openChannel()
			} else {
				logger.Log().Infoln("[", name, "]", "Channel is already open.")
			}
		}
	}()
	// initially, try to open channel is connection is opened.
	if connection.IsConnected() && !channel.isConnecting {
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
	c.mutex.Lock()
	c.isConnecting = true
	c.mutex.Unlock()
	logger.Log().Infoln("[", c.name, "]", "Opening channel...")
	if c.IsOpened() || !c.connection.IsConnected() {
		// Channel already opened and running
		// OR connection is not connected or ready yet.
		logger.Log().Infoln("[", c.name, "]", "Channel already opened.")
		c.mutex.Lock()
		c.isConnecting = false
		c.mutex.Unlock()
		return
	}
	c.mutex.Lock()
	if c.channel != nil && !c.channel.IsClosed() {
		c.channel.Close()
	}
	ch, err := c.connection.Connection().Channel()
	if err != nil {
		logger.Log().Errorln("[", c.name, "]", "Unable to open channel: ", err)
		logger.Log().Errorln("[", c.name, "]", "Retrying in 15 seconds...")
		c.isConnecting = false
		c.mutex.Unlock()
		time.AfterFunc(time.Duration(15)*time.Second, func() {
			c.openChannel()
		})
		return
	}
	c.channel = ch
	c.isConnecting = false
	c.mutex.Unlock()
	logger.Log().Infoln("[", c.name, "]", "Channel opened.")
	c.internalNotifyOpened <- c.channel // using internal, so autoconnect will listen to this event first.
	c.NotifyOpened <- c.channel
}

/**
 * Enable auto reconnect to channel in case of channel is closed unexpectedly.
 */
func (c *MQChannel) enableAutoReconnect() {
	if c.autoReconnect {
		go func() {
			// when channel is opened, handle auto channel open.
			for ch := range c.internalNotifyOpened {
				logger.Log().Info("[", c.name, "]", "Channel Opened. Listening to channel error...")
				errorChannel := ch.NotifyClose(make(chan *amqp091.Error, 1))
				err := <-errorChannel
				logger.Log().Error("[", c.name, "]", "Channel closed unexpectedly: ", err)
				logger.Log().Info("[", c.name, "]", "Re-opening channel in 15 seconds...")
				go func() {
					c.NotifyErrorClose <- true
				}()
				if !c.connection.IsConnected() {
					// seems like main connection is closed, dont do anything.
					// if auto connect enabled in connection, this channel will be initialized again.
					logger.Log().Errorln("[", c.name, "]", "Connection is not available. Unable to open channel.")
				} else {
					time.AfterFunc(time.Duration(15)*time.Second, func() {
						c.openChannel()
					})
				}
			}

			logger.Log().Info("[", c.name, "]", "Notify open channel closed...")
		}()
	}
}

/**
 * Use this to enable confirmation ack by server so that, client will be notified
 * of successful message is received by sever. Used by publisher.
 */
func (c *MQChannel) EnableConfirm() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	err := c.channel.Confirm(false)
	if err != nil {
		logger.Log().Errorln("[", c.name, "]", "Unable to enable confirmation ack, retrying")
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
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.channel != nil && !c.channel.IsClosed()
}

/**
 * Close channel.
 * Note: when closed, you cannot/dont reopen the channel.
 */
func (c *MQChannel) Close() {
	if c.IsOpened() {
		c.channel.Close()
	}
	for _, ch := range c.registeredCloseNotifiers {
		ch <- true // notify that channel is closed to each registered close notifier.
	}
	c.mutex.Lock()
	c.isClosed = true
	c.mutex.Unlock()
	defer close(c.NotifyOpened)
	defer close(c.internalNotifyOpened)
	defer close(c.NotifyErrorClose)
	logger.Log().Errorln("[", c.name, "]", "Channel is manually closed")
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
