package types

import (
	"github.com/rew3/rew3-internal/pubsub/mq/config"
	"github.com/rew3/rew3-internal/pubsub/mq/message"
)

/**
 * Publisher for Message Queue.
 */
type Publisher interface {
	Publish(config.RoutingKey, message.Message) (DeliveryTag, error)
	IsAck(DeliveryTag) bool
	OnAck(DeliveryTag, func())
	OnError(string, func())
	Stop()
}

/**
 * Consumer for Message Queue.
 */
type Consumer interface {
	Subscribe(int, bool) (<-chan message.Message, chan<- bool) // message channel, ack channel.
	Stop()
}

/**
 * ConsumerGroup for Message Queue.
 * This provides efficient way to define list of multiple consumer with shared channels.
 */
type ConsumerGroup interface {
	NewConsumer(props config.ConsumerProps) Consumer
	Stop()
}

/**
 * DeliveryTag identifier for particular message.
 * This tag can be used to check if message has been delivered or not.
 */
type DeliveryTag uint64
