package rtc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/ably/ably-go/ably"
	"github.com/rew3/rew3-pkg/pubsub/rt/config"
	"github.com/rew3/rew3-pkg/pubsub/rt/message"
	"github.com/rew3/rew3-pkg/pubsub/rt/types"
	"github.com/rew3/rew3-pkg/utils/logger"
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
	logger.Log().Infoln(c.logMsg("Subscribing to route="), route)
	channel := c.channel.Channel
	unsubscribe, err := channel.Subscribe(context.Background(), string(route), func(msg *ably.Message) {
		serialized, err := c.serialize(msg.Data)
		if err != nil {
			logger.Log().Errorln(c.logMsg("Unable to serialize received message."))
			return
		}
		onMessage(serialized)
	})
	if err != nil {
		logger.Log().Errorln(c.logMsg("Unable to subscribe: "), err)
		return types.Subscription{}, err
	}
	return types.Subscription{Unsubscribe: unsubscribe}, nil
}

/**
 * Subscribe to all messages in this consumer channel.
 */
func (c *AblyConsumer) SubscribeAll(onMessage func(received message.Message)) (types.Subscription, error) {
	logger.Log().Infoln(c.logMsg("Subscribing to all route messages..."))
	channel := c.channel.Channel
	unsubscribe, err := channel.SubscribeAll(context.Background(), func(msg *ably.Message) {
		serialized, err := c.serialize(msg.Data)
		if err != nil {
			logger.Log().Errorln(c.logMsg("Unable to serialize received message."))
			return
		}
		onMessage(serialized)
	})
	if err != nil {
		logger.Log().Errorln(c.logMsg("Unable to subscribe: "), err)
		return types.Subscription{}, err
	}
	return types.Subscription{Unsubscribe: unsubscribe}, nil
}

/**
 * Serialize given data from ably server to RT Message type.
 */
func (c *AblyConsumer) serialize(data interface{}) (message.Message, error) {
	var message message.Message
	jsonString, ok := data.(string)
	if !ok {
		logger.Log().Errorln(c.logMsg("Invalid data, Unable to serizlie message"))
		return message, errors.New("invalid data, unable to serizlie message")
	}
	err := json.Unmarshal([]byte(jsonString), &message)
	if err != nil {
		logger.Log().Errorln(c.logMsg("Error converting to Message type: "), err)
		return message, err
	}
	return message, nil
}

// create log message
func (c *AblyConsumer) logMsg(msg string) string {
	return "[RT Consumer: " + string(c.channel.Name) + "] " + msg
}
