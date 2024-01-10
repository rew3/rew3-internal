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
 * Subscribe to messages published in given route in this consumer channel.
 */
func (c *AblyConsumer) Subscribe(route config.MessageRoute, onMessage func(received message.Message)) (types.Subscription, error) {
	channel := c.channel.Channel
	unsubscribe, err := channel.Subscribe(context.Background(), string(route), func(msg *ably.Message) {
		serialized, err := c.serialize(msg.Data)
		if err != nil {
			logger.Log().Errorln("Unable to serialize received message.")
			return
		}
		onMessage(serialized)
	})
	if err != nil {
		logger.Log().Errorln("Unable to subscribe: ", err)
		return types.Subscription{}, err
	}
	return types.Subscription{Unsubscribe: unsubscribe}, nil
}

/**
 * Subscribe to all messages in this consumer channel.
 */
func (c *AblyConsumer) SubscribeAll(onMessage func(received message.Message)) (types.Subscription, error) {
	channel := c.channel.Channel
	unsubscribe, err := channel.SubscribeAll(context.Background(), func(msg *ably.Message) {
		serialized, err := c.serialize(msg.Data)
		if err != nil {
			logger.Log().Errorln("Unable to serialize received message.")
			return
		}
		onMessage(serialized)
	})
	if err != nil {
		logger.Log().Errorln("Unable to subscribe: ", err)
		return types.Subscription{}, err
	}
	return types.Subscription{Unsubscribe: unsubscribe}, nil
}

/**
 * Serialize given data from ably server to RT Message type.
 */
func (c *AblyConsumer) serialize(data interface{}) (message.Message, error) {
	var message message.Message
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Log().Errorln("Error converting to JSON: ", err)
		return message, err
	}
	err = json.Unmarshal(jsonData, &message)
	if err != nil {
		logger.Log().Errorln("Error converting to Message type: ", err)
		return message, err
	}
	return message, nil
}
