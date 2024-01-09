package config

/**
 * Client configuration.
 * Contains configuration parameters to setup client for realtime server.
 * Note: Client ID can be also used while publishing message, so the message
 * will be directly received by that particular client service.
 */
type ClientConfig struct {
	APIKey   string // API key for realtime server.
	ClientId string // Identifier of the client (current service using rt)
}

/**
 * Channel name.
 */
type ChannelName string

/**
 * Route name for message/events being published/consumed. 
 * Client can use this in order to publish or consume message of particular type.   
 */
type MessageRoute string
