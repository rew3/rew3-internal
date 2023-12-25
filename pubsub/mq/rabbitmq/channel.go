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
 */
type MQChannel struct {
	mutex                sync.Mutex
	connection           *MQConnection
	channel              *amqp091.Channel
	AutoReconnect        bool                  // in case of connection error, reopen channel automatically
	PublisherConfirmMode bool                  // provide confirmation so that client can ensure all publishings have successfully been received by the server.
	NotifyOpened         chan *amqp091.Channel // go channel to notify for new opened channel (e.g re-opened)
	NotifyErrorClose     chan bool             // go channel to notify when channel is closed by error.
}

/**
 * Create new channel.
 */
func NewChannel(autoReconnect bool, connection *MQConnection) *MQChannel {
	channel := &MQChannel{
		mutex:                sync.Mutex{},
		connection:           connection,
		channel:              nil,
		AutoReconnect:        autoReconnect,
		PublisherConfirmMode: true,
		NotifyOpened:         make(chan *amqp091.Channel, 1),
		NotifyErrorClose:     make(chan bool, 1),
	}
	if connection.IsConnected() {
		channel.openChannel()
	}
	go func() {
		// When connection is ready, open channel.
		for range connection.NotifyConnected() {
			channel.openChannel()
		}
	}()
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
	ch, err := c.connection.Connection().Channel()
	if err != nil {
		logger.Log().Errorln("Unable to open channel: ", err)
		logger.Log().Errorln("Retrying in 10 seconds...")
		time.AfterFunc(time.Duration(10)*time.Second, func() {
			c.openChannel()
		})
	}
	c.mutex.Lock()
	c.channel = ch
	c.mutex.Unlock()
	if c.AutoReconnect {
		errorChannel := ch.NotifyClose(make(chan *amqp091.Error, 1))
		go func() {
			for err := range errorChannel {
				logger.Log().Error("Channel closed unexpectedly: ", err)
				logger.Log().Info("Re-opening channel...")
				c.NotifyErrorClose <- true
				if !c.connection.IsConnected() {
					// seems like main connection is closed, dont do anything.
					// if auto connect enabled in connection, this channel will be initialized again.
					return
				} else {
					close(errorChannel)
					time.AfterFunc(time.Duration(10)*time.Second, func() {
						c.openChannel()
					})
				}
			}
		}()
	}
	logger.Log().Infoln("Channel opened")
	c.NotifyOpened <- c.channel
}

/**
 * Use this to enable confirmation ack by server so that, client will be notified
 * of successful message is received by sever.
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
 * Note: when closed, you cannot reopen the channel.
 */
func (c *MQChannel) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
	close(c.NotifyOpened)
}
