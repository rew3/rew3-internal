package rtc

import (
	"github.com/ably/ably-go/ably"
	"github.com/rew3/rew3-internal/pkg/utils/logger"
	"github.com/rew3/rew3-internal/pubsub/rt/config"
)

/**
 * Ably Realtime client connection.
 *
 * AblyConnection and Recovery:
 * https://ably.com/docs/connect/states
 */
type AblyConnection struct {
	client *ably.Realtime
}

/**
 * Realtime channel.
 * This channel is used to publish and subscribe to messages.
 */
type AblyChannel struct {
	Name    config.ChannelName
	Channel *ably.RealtimeChannel
}

/**
 * Create new ably client connection.
 * Return error, if there is error while initializing connection.
 */
func NewAblyConnection(config config.ClientConfig) (*AblyConnection, error) {
	client, err := ably.NewRealtime(
		ably.WithKey(config.APIKey),
		ably.WithClientID(config.ClientId),
		ably.WithAutoConnect(false), // Set this option to avoid missing state changes.
	)
	if err != nil {
		return nil, err
	}
	// Set up connection events handler.
	client.Connection.OnAll(func(change ably.ConnectionStateChange) {
		logger.Log().Infoln("Connection event:", change.Event, ", state=", change.Current, ", reason=", change.Reason)
	})
	return &AblyConnection{client: client}, nil
}

/**
 * Connect to the realtime server and start the connection.
 * Note: you will be only able to listen/publish messages after call this Connect().
 */
func (client *AblyConnection) Connect() {
	client.client.Connect()
}

/**
 * Close the connection explicitly.
 * All the channels created with this connection will be automatically closed.
 */
func (client *AblyConnection) Close() {
	client.client.Close()
}

/**
 * Get the channel for given channel name.
 * A channel is used to publish and subscribe to the realtime messages. This returns existing channel
 * if its already created with given channel name otherwise create new channel and return.
 *
 * Note: Channel connection is automatically handled.
 */
func (client *AblyConnection) GetChannel(name config.ChannelName) *AblyChannel {
	channel := client.client.Channels.Get(string(name))
	channel.OnAll(func(change ably.ChannelStateChange) {
		logger.Log().Infoln("Connection event:", change.Event, "channel=", channel.Name, ", state=", change.Current, ", reason=", change.Reason)
	})
	return &AblyChannel{Name: name, Channel: channel}
}
