package mq

import (
	"github.com/rew3/rew3-pkg/pubsub/mq/config"
	"github.com/rew3/rew3-pkg/pubsub/mq/rabbitmq"
	"github.com/rew3/rew3-pkg/pubsub/mq/types"
)

/**
 * Message Queue Client.
 */
type MQClient struct {
	URL        string
	connection *rabbitmq.MQConnection
}

/**
 * Create new message queue server.
 */
func NewMQClient(url string) *MQClient {
	connection := rabbitmq.NewConnection(url)
	return &MQClient{
		connection: connection,
	}
}

/**
 * Connect to message queue server.
 */
func (s *MQClient) ConnectServer() {
	s.connection.Connect()
}

/**
 * Create new publisher.
 */
func (s *MQClient) Publisher(props config.PublisherProps) types.Publisher {
	return rabbitmq.NewPublisher(s.connection, props)
}

/**
 * Create new consumer.
 */
func (s *MQClient) Consumer(props config.ConsumerProps) types.Consumer {
	return rabbitmq.NewConsumer(s.connection, props)
}

/**
 * Create new consumer group.
 */
func (s *MQClient) ConsumerGroup(name string) types.ConsumerGroup {
	return rabbitmq.NewConsumerGroup(name, s.connection)
}
