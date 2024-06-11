package types

import (
	"github.com/rew3/rew3-pkg/pubsub/rt/config"
	"github.com/rew3/rew3-pkg/pubsub/rt/message"
)

/**
 * Realtime Publisher interface.
 */
type RTPublisher interface {
	Publish(config.MessageRoute, message.Message) error
	PublishAsync(config.MessageRoute, message.Message, func(error)) error // Publish asynchronously.
}

/**
 * Realtime Consumer interface.
 */
type RTConsumer interface {
	Subscribe(config.MessageRoute, func(message.Message)) (Subscription, error)
	SubscribeAll(func(message.Message)) (Subscription, error)
}

/**
 * Subscription.
 */
type Subscription struct {
	Unsubscribe func()
}
