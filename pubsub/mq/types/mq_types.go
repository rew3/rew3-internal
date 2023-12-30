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
}

/**
 * Consumer for Message Queue.
 */
type Consumer interface {
	Subscribe(int, bool) (chan<- message.Message, <-chan bool) // message channel, ack channel.
}

/**
 * DeliveryTag identifier for particular message.
 * This tag can be used to check if message has been delivered or not.
 */
type DeliveryTag uint64
