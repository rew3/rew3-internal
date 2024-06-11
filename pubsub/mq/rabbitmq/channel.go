package rabbitmq

import (
	"sync"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rew3/rew3-pkg/utils/logger"
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
	mutex         sync.Mutex
	retryTime     int // retry time in seconds.
	name          string
	connection    *MQConnection
	channel       *amqp091.Channel
	isClosed      bool
	isConnecting  bool
	autoReconnect bool        // in case of connection error, reopen channel automatically
	onReady       []chan bool // list of callback, when channel is opened.
	onErrorClose  []chan bool // list of callback, when channel stop with error.
}

/**
 * Create new channel.
 */
func NewChannel(name string, autoReconnect bool, connection *MQConnection) *MQChannel {
	channel := &MQChannel{
		mutex:         sync.Mutex{},
		retryTime:     15,
		name:          name,
		connection:    connection,
		channel:       nil,
		autoReconnect: autoReconnect,
		isConnecting:  false,
		isClosed:      false,
		onReady:       []chan bool{},
		onErrorClose:  []chan bool{},
	}
	go func() {
		// When connection is ready, open channel. useful for re open the
		// channel when connection is reconnected.
		for range connection.NotifyReady(make(chan bool, 1)) {
			logger.Log().Infoln("[MQ Channel: ", name, "]", "Connection is Ready. Initializing Channel...")
			channel.mutex.Lock()
			isConnecting := channel.isConnecting
			isClosed := channel.isClosed
			channel.mutex.Unlock()
			if isClosed {
				logger.Log().Infoln("[MQ Channel: ", name, "]", "Cancel Init. Channel has been gracefully closed.")
				return
			}
			if isConnecting {
				return
			}
			if !channel.IsOpened() {
				channel.openChannel()
			} else {
				logger.Log().Infoln("[MQ Channel: ", name, "]", "Channel is already opened.")
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
 * Note: if error occured while opening channel, it will be retried after 15 seconds
 * and new channel is opened. In addition, when closed, you cannot reopen the channel,
 * instead create new.
 */
func (c *MQChannel) openChannel() {
	c.mutex.Lock()
	c.isConnecting = true
	isClosed := c.isClosed
	c.mutex.Unlock()
	logger.Log().Infoln("[MQ Channel: ", c.name, "]", "Opening Channel...")
	if isClosed {
		logger.Log().Infoln("[MQ Channel: ", c.name, "]", "Cancel open. Channel has been gracefully closed.")
		return
	}
	if c.IsOpened() || !c.connection.IsConnected() {
		// Channel already opened and running
		// OR connection is not connected or ready yet.
		logger.Log().Infoln("[MQ Channel: ", c.name, "]", "Channel already opened or Connection is not available.")
		c.mutex.Lock()
		c.isConnecting = false
		c.mutex.Unlock()
		return
	}
	ch, err := c.connection.Connection().Channel()
	if err != nil {
		logger.Log().Errorln("[MQ Channel: ", c.name, "]", "Unable to open channel: ", err)
		logger.Log().Errorln("[MQ Channel: ", c.name, "]", "Re-opening channel in ", c.retryTime, " seconds...")
		c.mutex.Lock()
		c.isConnecting = false
		c.mutex.Unlock()
		time.AfterFunc(time.Duration(c.retryTime)*time.Second, func() {
			c.openChannel()
		})
		return
	}
	c.mutex.Lock()
	c.channel = ch
	c.isConnecting = false
	c.mutex.Unlock()
	logger.Log().Infoln("[MQ Channel: ", c.name, "]", "Channel is opened.")
	go c.listenToChannelError(c.channel)
	go func() {
		logger.Log().Info("[MQ Channel: ", c.name, "]", "Notifying channel opened event to registered callbacks...")
		for _, cb := range c.onReady {
			go func(callback chan<- bool) {
				callback <- true
			}(cb)
		}
	}()
}

/**
 * Listen to channel errors.
 */
func (c *MQChannel) listenToChannelError(ch *amqp091.Channel) {
	logger.Log().Info("[MQ Channel: ", c.name, "]", "Start listening to this channel errors...")
	errorChannel := ch.NotifyClose(make(chan *amqp091.Error, 1))
	err := <-errorChannel
	logger.Log().Error("[MQ Channel: ", c.name, "]", "Channel closed unexpectedly: ", err)
	go func() {
		logger.Log().Info("[MQ Channel: ", c.name, "]", "Notifying error to registered callbacks...")
		for _, cb := range c.onErrorClose {
			go func(callback chan<- bool) {
				callback <- true
			}(cb)
		}
	}()
	if c.autoReconnect {
		logger.Log().Info("[MQ Channel: ", c.name, "]", "Auto re-connect is enabled.")
		if !c.connection.IsConnected() {
			// seems like main connection is closed, dont do anything.
			// if auto connect enabled in connection, this channel will be initialized again.
			logger.Log().Errorln("[MQ Channel: ", c.name, "]", "Unable to re-connect, Connection is not available.")
		} else {
			logger.Log().Errorln("[MQ Channel: ", c.name, "]", "Re-opening channel in ", c.retryTime, " seconds...")
			time.AfterFunc(time.Duration(c.retryTime)*time.Second, func() {
				c.openChannel()
			})
		}
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
		logger.Log().Errorln("[MQ Channel: ", c.name, "]", "Unable to enable confirmation ack, retrying")
	}
}

/**
 * Get opened channel.
 * Note: if channel is not opened yet, nil is returned.
 * Alternatively, You can listen to NotifyReady channel to know when channel is opened and is
 * ready to process messages and then call GetChannel() again.
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
	return c.channel != nil && !c.channel.IsClosed() && !c.isClosed
}

/**
 * Close channel.
 * Note: when closed, you cannot/dont reopen the channel.
 */
func (c *MQChannel) Close() {
	logger.Log().Errorln("[MQ Channel: ", c.name, "]", "Closing channel...")
	if c.IsOpened() {
		c.channel.Close()
	}
	for _, ch := range c.onReady {
		close(ch)
	}
	for _, ch := range c.onErrorClose {
		close(ch)
	}
	c.mutex.Lock()
	c.isClosed = true
	c.onReady = []chan bool{}
	c.onErrorClose = []chan bool{}
	c.mutex.Unlock()
	logger.Log().Errorln("[MQ Channel: ", c.name, "]", "Channel is manually closed")
}

/**
 * Notify when channel is ready/opened.
 * Note: returned go channel will be closed when this channel is gracefully closed.
 */
func (c *MQChannel) NotifyReady(ch chan bool) <-chan bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.onReady = append(c.onReady, ch)
	return ch
}

/**
 * Notify when channel is closed because of error.
 * Note: returned go channel will be closed when this channel is gracefully closed.
 */
func (c *MQChannel) NotifyErrorClose(ch chan bool) <-chan bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.onErrorClose = append(c.onErrorClose, ch)
	return ch
}
