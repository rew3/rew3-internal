package query

import "reflect"

/**
 * A super type for Query message.
 * All query message must use this root command using composition.
 */
type Query interface{}

/**
 * Get name of given query.
 * Note: you must provide query instance in pointer if it is interface. e.g. ContactQuery. query := ContactQuery(nil); QueryName(&query)
 */
func QueryName(query Query) string {
	queryType := reflect.TypeOf(query)
	if queryType.Kind() == reflect.Ptr {
		queryType = queryType.Elem()
	}
	return queryType.Name()
}

/*
 * Get Parent of given concrete query (not interface). 
 * e.g. CountContacts will return ContactQuery. 
 */
func ParentQueryName(q Query) string {
	qType := reflect.TypeOf(q)
	if qType.Kind() == reflect.Struct {
		// Loop through the fields of the struct to find the embedded type
		for i := 0; i < qType.NumField(); i++ {
			field := qType.Field(i)
			if field.Anonymous {
				// Check if the embedded field implements Query interface
				if field.Type.Implements(reflect.TypeOf((*Query)(nil)).Elem()) {
					return field.Type.Name()
				}
			}
		}
	}
	return ""
}
