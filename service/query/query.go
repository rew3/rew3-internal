package query

import "reflect"

/**
 * A super type for Query message.
 * All query message must use this root command using composition.
 */
type Query interface{}

func QueryName(query Query) string {
	return reflect.TypeOf((query)).Elem().Name()
}
