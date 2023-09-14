package query

import "reflect"

/**
 * A super type for Query message.
 * All query message must use this root command using composition.
 */
type Query interface{}

/**
 * Get name of given query. 
 * Note: you must provide query instance in pointer. e.g. ContactQuery. query := ContactQuery(nil); QueryName(&query)
 */
func QueryName(query Query) string {
	queryType := reflect.TypeOf(query)
	if queryType.Kind() == reflect.Ptr {
		queryType = queryType.Elem()
	}
	return queryType.Name()
}
