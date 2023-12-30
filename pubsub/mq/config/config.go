package config

/**
 * Exchange Properties.
 */
type ExchangeProps struct {
	Name      ExchangeName
	Type      DistributionType
	IsDurable bool
}

/**
 * Properties for publisher.
 */
type PublisherProps struct {
	ExchangeProps ExchangeProps
}

/**
 * Properties for consumer.
 */
type ConsumerProps struct {
	Name           string // name for queue. dont use spaces or special characters.
	ExchangeProps  ExchangeProps
	RoutingKeys    []RoutingKey
	IsDurableQueue bool // message are not loss on server fail/restart.
}

/**
 * Define exchange properties.
 */
func DefineExchangeProps(name ExchangeName, dType DistributionType, isDurable bool) ExchangeProps {
	return ExchangeProps{
		Name:      name,
		Type:      dType,
		IsDurable: isDurable,
	}
}

/**
 * Define message queue properties for publisher.
 */
func DefinePublisherProps(exchangeProps ExchangeProps) PublisherProps {
	return PublisherProps{
		ExchangeProps: exchangeProps,
	}
}

/**
 * Define message queue properties for consumer.
 */
func DefineConsumerProps(exchangeProps ExchangeProps, keys []RoutingKey, isDurable bool) ConsumerProps {
	return ConsumerProps{
		ExchangeProps:  exchangeProps,
		RoutingKeys:    keys,
		IsDurableQueue: isDurable,
	}
}

/**
 * Message distribution type.
 */
type DistributionType string

const (
	BROADCAST DistributionType = "fanout" // Broadcast distribution. Message is distributed to all consumers.
	DIRECT    DistributionType = "direct" // Direct message distribution. Message is distributed to selected consumer with matched routing key.
	TOPIC     DistributionType = "topic"  // Topic based distribution,  Message is distributed to selected consumer with wildcard matched routing key.
)

/**
 * Routing key for selective group.
 */
type ExchangeName string

/**
 * Routing key for selective group.
 */
type RoutingKey string
