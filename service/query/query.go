package query

/**
 * A super type for Query message.
 * All query message must use this root command using composition.
 */
type Query interface {
	GetName() string
}
