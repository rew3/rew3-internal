package rabbitmq

import (
	"sync"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rew3/rew3-internal/pkg/utils/logger"
)

/**
 * Message Queue Connection.
 * This connection has automatic connection recovery with every connection retry in every 10 seconds
 * in case of network faliure or disgraceful disconnect.
 */
type MQConnection struct {
	mutex           sync.Mutex
	URL             string
	connection      *amqp091.Connection
	notifyConnected chan bool // Use this channel to get notified when connection is connected.
}

func NewConnection(url string) *MQConnection {
	return &MQConnection{
		mutex:           sync.Mutex{},
		URL:             url,
		connection:      nil,
		notifyConnected: make(chan bool, 1),
	}
}

/**
 * Connect to the message queue server.
 * If there is error in connecting, it will be retried after every 10 seconds in
 * non-blocking async way. When connected, status will be notified via NotifyConnected
 * go channel.
 */
func (c *MQConnection) Connect() {
	if c.IsConnected() {
		return // connection is already connected.
	}
	c.mutex.Lock()
	connection, err := amqp091.Dial(c.URL)
	c.mutex.Unlock()
	if err != nil {
		logger.Log().Error("Unable to connect to Message Queue Server: ", c.URL)
		logger.Log().Info("Retrying in 10 seconds...")
		time.AfterFunc(time.Duration(10)*time.Second, func() {
			c.Connect()
		})
		return
	}
	c.mutex.Lock()
	c.connection = connection
	c.mutex.Unlock()
	c.notifyConnected <- true
	logger.Log().Info("Connection establish successful")

	errorChannel := c.connection.NotifyClose(make(chan *amqp091.Error))
	go func() {
		err, ok := <-errorChannel
		logger.Log().Error("Connection closed unexpectedly: ", err)
		logger.Log().Info("Reconnecting connection in 10 seconds...")
		if ok { // close the channel if not closed.
			close(errorChannel)
		}
		time.AfterFunc(time.Duration(10)*time.Second, func() {
			c.Connect()
		})
	}()
}

/**
 * Get connection.
 * If not connected, return nil. Note, listen to NotifyConnected channel, to get
 * notified when connection is established, and calling this will have valid connection
 * instance.If there is connection error, connection is retried to re-established and
 * when connected, is notified with NotifyConnected.
 */
func (c *MQConnection) Connection() *amqp091.Connection {
	return c.connection
}

/**
 * Use this channel to get notified when connection is connected.
 */
func (c *MQConnection) NotifyConnected() <-chan bool {
	return c.notifyConnected
}

/**
 * Check if connection is connected or not.
 */
func (c *MQConnection) IsConnected() bool {
	return c.connection != nil && !c.connection.IsClosed()
}

/**
 * Stop the connection.
 * Note: must be accessed from same goroutine where connection is created. 
 */
func (c *MQConnection) Stop() {
	if c.connection != nil {
		c.connection.Close()
	}
	close(c.notifyConnected)
}
