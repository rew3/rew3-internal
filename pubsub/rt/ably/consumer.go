package rtc

import (
	"context"
	"encoding/json"

	"github.com/ably/ably-go/ably"
	"github.com/rew3/rew3-internal/pkg/utils/logger"
	"github.com/rew3/rew3-internal/pubsub/rt/config"
	"github.com/rew3/rew3-internal/pubsub/rt/message"
	"github.com/rew3/rew3-internal/pubsub/rt/types"
)

/**
 * Ably Realtime Consumer.
 * Use this consumer to subscribe to new message from realtime service.
 */
type AblyConsumer struct {
	types.RTConsumer
	channel *AblyChannel
}

/**
 * Create new consumer with given channel name.
 */
func NewAblyConsumer(channelName config.ChannelName, connection *AblyConnection) types.RTConsumer {
	channel := connection.GetChannel(channelName)
	return &AblyConsumer{
		channel: channel,
	}
}

/**
 * Subscribe to message for given category messages in this consumer channel.
 */
func (c *AblyConsumer) Subscribe(category message.MessageCategory, onMessage func(received message.Message)) {
	channel := c.channel.Channel
	channel.Subscribe(context.Background(), string(category), func(msg *ably.Message) {
		serialized, err := c.serialize(msg.Data)
		if err != nil {
			logger.Log().Errorln("Unable to serialize received message.")
			return
		}
		onMessage(serialized)
	})
}

/**
 * Subscribe to all messages in this consumer channel.
 */
func (c *AblyConsumer) SubscribeAll(onMessage func(received message.Message)) {
	channel := c.channel.Channel
	channel.SubscribeAll(context.Background(), func(msg *ably.Message) {
		serialized, err := c.serialize(msg.Data)
		if err != nil {
			logger.Log().Errorln("Unable to serialize received message.")
			return
		}
		onMessage(serialized)
	})
}

/**
 * Serialize given data from ably server to RT Message type.
 */
func (c *AblyConsumer) serialize(data interface{}) (message.Message, error) {
	var message message.Message
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Log().Errorln("Error converting to JSON: %v\n", err)
		return message, err
	}
	err = json.Unmarshal(jsonData, &message)
	if err != nil {
		logger.Log().Errorln("Error converting to Message type: %v\n", err)
		return message, err
	}
	return message, nil
}
