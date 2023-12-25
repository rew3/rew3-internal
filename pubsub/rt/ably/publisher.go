package rtc

import (
	"context"

	"github.com/rew3/rew3-internal/pkg/utils/logger"
	"github.com/rew3/rew3-internal/pubsub/rt/config"
	"github.com/rew3/rew3-internal/pubsub/rt/message"
	"github.com/rew3/rew3-internal/pubsub/rt/types"
)

/**
 * Ably Realtime Publisher.
 * Use this publisher to publish new message to realtime service.
 */
type AblyPublisher struct {
	types.RTPublisher
	channel *AblyChannel
}

/**
 * Create new publisher with given channel name.
 */
func NewAblyPublisher(channelName config.ChannelName, connection *AblyConnection) types.RTPublisher {
	channel := connection.GetChannel(channelName)
	return &AblyPublisher{
		channel: channel,
	}
}

/**
 * Publish message.
 */
func (p *AblyPublisher) Publish(message message.Message) error {
	eventType := message.Category
	err := p.channel.Channel.Publish(context.Background(), string(eventType), message)
	if err != nil {
		logger.Log().Errorln("Error while publishing: ", err)
		return err
	}
	return nil
}

/**
 * Publish message asynchronously.
 * Note: onAck() will be called when message is acknowledge successfully or failed with error.
 * onAck() will be called in goroutine to avoid ably running thread block.
 */
func (p *AblyPublisher) PublishAsync(message message.Message, onAck func(error)) error {
	eventType := message.Category
	err := p.channel.Channel.PublishAsync(string(eventType), message, func(err error) {
		go onAck(err)
	})
	if err != nil {
		logger.Log().Errorln("Error while publishing: ", err)
	}
	return nil
}
