package types

import "github.com/rew3/rew3-internal/pubsub/rt/message"

/**
 * Realtime Publisher interface.
 */
type RTPublisher interface {
	Publish(message.Message) error
	PublishAsync(message.Message, func(error)) error // Publish asynchronously.
}

/**
 * Realtime Consumer interface.
 */
type RTConsumer interface {
	Subscribe(message.MessageCategory, func(message.Message))
	SubscribeAll(func(message.Message))
}
