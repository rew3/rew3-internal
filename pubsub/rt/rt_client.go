package rt

import (
	"github.com/rew3/rew3-internal/pkg/utils/logger"
	rtc "github.com/rew3/rew3-internal/pubsub/rt/ably"
	"github.com/rew3/rew3-internal/pubsub/rt/config"
	"github.com/rew3/rew3-internal/pubsub/rt/types"
)

/**
 * Realtime sevice client.
 */
type RTClient struct {
	Connection *rtc.AblyConnection
}

/**
 * Create and initialize Realtime client.
 * Note: must call Connection.Connect() to connection to realtime service.
 */
func NewRTClient(config config.ClientConfig) (*RTClient, error) {
	connection, err := rtc.NewAblyConnection(config)
	if err != nil {
		logger.Log().Errorln("Unable create Realtime Client Connection: ", err)
		return nil, err
	}
	return &RTClient{
		Connection: connection,
	}, nil
}

/**
 * Create new Publisher.
 * Note: every call will initialize and create new publisher.
 */
func (c *RTClient) Publisher(name config.ChannelName) *types.RTPublisher {
	publisher := rtc.NewAblyPublisher(name, c.Connection)
	logger.Log().Infoln("Publisher created for channel: ", name)
	return &publisher
}

/**
 * Create new Publisher.
 * Note: every call will initialize and create new subscriber.
 */
func (c *RTClient) Consumer(name config.ChannelName) *types.RTConsumer {
	consumer := rtc.NewAblyConsumer(name, c.Connection)
	logger.Log().Infoln("Consumer created for channel: ", name)
	return &consumer
}
