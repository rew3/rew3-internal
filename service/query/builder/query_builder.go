package builder

import (
	"fmt"
	"strings"
)

/**
 * Custom Query Builder for MongoDB.
 */

type Query map[string]interface{}

type QueryBuilder struct {
	query Query
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		query: Query{},
	}
}

func (qb *QueryBuilder) Eq(key string, value interface{}) *QueryBuilder {
	qb.query[key] = Query{"$eq": value}
	return qb
}

func (qb *QueryBuilder) Ne(key string, value interface{}) *QueryBuilder {
	qb.query[key] = Query{"$ne": value}
	return qb
}

func (qb *QueryBuilder) Gt(key string, value interface{}) *QueryBuilder {
	qb.query[key] = Query{"$gt": value}
	return qb
}

func (qb *QueryBuilder) Gte(key string, value interface{}) *QueryBuilder {
	qb.query[key] = Query{"$gte": value}
	return qb
}

func (qb *QueryBuilder) Lt(key string, value interface{}) *QueryBuilder {
	qb.query[key] = Query{"$lt": value}
	return qb
}

func (qb *QueryBuilder) Lte(key string, value interface{}) *QueryBuilder {
	qb.query[key] = Query{"$lte": value}
	return qb
}

func (qb *QueryBuilder) In(key string, values []interface{}) *QueryBuilder {
	qb.query[key] = Query{"$in": values}
	return qb
}

func (qb *QueryBuilder) Nin(key string, values []interface{}) *QueryBuilder {
	qb.query[key] = Query{"$nin": values}
	return qb
}

func (qb *QueryBuilder) Exists(key string, exists bool) *QueryBuilder {
	qb.query[key] = Query{"$exists": exists}
	return qb
}

func (qb *QueryBuilder) Regex(key string, pattern string) *QueryBuilder {
	qb.query[key] = Query{"$regex": pattern}
	return qb
}

func (qb *QueryBuilder) And(conditions ...*QueryBuilder) *QueryBuilder {
	var andQuery []Query
	for _, cond := range conditions {
		andQuery = append(andQuery, cond.query)
	}
	qb.query["$and"] = andQuery
	return qb
}

func (qb *QueryBuilder) Or(conditions ...*QueryBuilder) *QueryBuilder {
	var orQuery []Query
	for _, cond := range conditions {
		orQuery = append(orQuery, cond.query)
	}
	qb.query["$or"] = orQuery
	return qb
}

func (qb *QueryBuilder) Not(condition *QueryBuilder) *QueryBuilder {
	qb.query["$not"] = condition.query
	return qb
}

func (qb *QueryBuilder) Build() Query {
	return qb.query
}

func (qb *QueryBuilder) String() string {
	var parts []string
	for key, value := range qb.query {
		parts = append(parts, fmt.Sprintf(`"%s": %v`, key, value))
	}
	return "{" + strings.Join(parts, ", ") + "}"
}
