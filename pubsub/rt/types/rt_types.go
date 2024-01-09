package types

import (
	"github.com/rew3/rew3-internal/pubsub/rt/config"
	"github.com/rew3/rew3-internal/pubsub/rt/message"
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
	Subscribe(config.MessageRoute, func(message.Message))
	SubscribeAll(func(message.Message))
}
