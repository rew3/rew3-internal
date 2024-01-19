package rabbitmq

import (
	"sync"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rew3/rew3-internal/pkg/utils/logger"
)

/**
 * Message Queue Connection.
 * This connection has automatic connection recovery with every connection retry in every 15 seconds
 * in case of network faliure or disgraceful disconnect.
 */
type MQConnection struct {
	mutex      sync.Mutex
	RetryTime  int //retry time in secondss
	URL        string
	connection *amqp091.Connection
	onReady    []chan bool // list of callback, to call when connection is ready/connected.
}

func NewConnection(url string) *MQConnection {
	return &MQConnection{
		mutex:      sync.Mutex{},
		RetryTime:  15,
		URL:        url,
		connection: nil,
		onReady:    []chan bool{},
	}
}

/**
 * Connect to the message queue server.
 * If there is error in connecting, it will be retried after every 10 seconds in
 * non-blocking async way. When connected, status will be notified via NotifyConnected
 * go channel.
 */
func (c *MQConnection) Connect() {
	logger.Log().Info("[MQ Connection] ", "Connecting...")
	if c.IsConnected() {
		logger.Log().Info("[MQ Connection] ", "Already Connected.")
		return // connection is already connected.
	}
	c.mutex.Lock()
	connection, err := amqp091.Dial(c.URL)
	c.mutex.Unlock()
	if err != nil {
		logger.Log().Error("[MQ Connection] ", "Unable to establish MQ Connection: ", c.URL)
		logger.Log().Info("[MQ Connection] ", "Retrying to establish connection in ", c.RetryTime, " seconds...")
		time.AfterFunc(time.Duration(c.RetryTime)*time.Second, func() {
			c.Connect()
		})
		return
	}
	c.mutex.Lock()
	c.connection = connection
	c.mutex.Unlock()
	logger.Log().Info("[MQ Connection] ", "Connection established successful")
	go func() {
		logger.Log().Info("[MQ Connection] ", "Notifying to registered callbacks...")
		for _, cb := range c.onReady {
			go func(callback chan<- bool) {
				callback <- true
			}(cb)
		}
	}()
	errorChannel := c.connection.NotifyClose(make(chan *amqp091.Error))
	go func() {
		err := <-errorChannel
		logger.Log().Error("[MQ Connection] ", "Connection closed unexpectedly: ", err)
		logger.Log().Info("[MQ Connection] ", "Retrying to establish connection in ", c.RetryTime, " seconds...")
		time.AfterFunc(time.Duration(c.RetryTime)*time.Second, func() {
			c.Connect()
		})
	}()
}

/**
 * Get connection.
 * If not connected, return nil. Note, listen to NotifyReady channel, to get
 * notified when connection is established, and call Connection() again to get connection instance.
 */
func (c *MQConnection) Connection() *amqp091.Connection {
	return c.connection
}

/**
 * Use this channel to get notified when connection is connected.
 */
func (c *MQConnection) NotifyReady(ch chan bool) <-chan bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.onReady = append(c.onReady, ch)
	return ch
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
	logger.Log().Info("[MQ Connection] ", "Stopping connection...")
	if c.connection != nil {
		c.connection.Close()
	}
	c.mutex.Lock()
	for _, ch := range c.onReady {
		close(ch)
	}
	c.onReady = []chan bool{}
	c.mutex.Unlock()
	logger.Log().Info("[MQ Connection] ", "Connection Stopped.")
}
