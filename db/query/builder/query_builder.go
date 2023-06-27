package builder

import (
	dsl "github.com/rew3/rew3-internal/db/query"
	"go.mongodb.org/mongo-driver/bson"
)

/**
* Custom Query Builder for MongoDB.
TODO - Use custom version
// Define regex for match, start with, end with, contains???

// Add more, starts with, ends with, based on our api query.
*/

// query builder.

type RegexType string

const (
	MATCHES_REGEX     RegexType = "matches"
	STARTS_WITH_REGEX RegexType = "starts_with"
	ENDS_WITH_REGEX   RegexType = "ends_with"
)

type QueryBuilder struct {
	queryDSL     dsl.RootDSL
	mongoBuilder *MongoQueryBuilder
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{mongoBuilder: NewMongoQueryBuilder()}
}

func (builder *QueryBuilder) ForQuery(query dsl.RootDSL) {
	builder.queryDSL = query
}

func (builder *QueryBuilder) Mongo() bson.D {
	generateMongoQuery := func (queries []dsl.QueryDSL) bson.D {

	}
	queries := builder.queryDSL.Queries
	for _, query := range queries {
		switch q := query.(type) {
		case dsl.ANDLogicalDSL:
			builder.mongoBuilder.Comparison()
		case dsl.ORLogicalDSL:
		case dsl.NOTLogicalDSL:

		case query.CriteriaDSL:
		default:

		}
	}
}

func (builder *QueryBuilder) generateMongoQuery(queries []dsl.QueryDSL) bson.D {
	doc := bson.D{}
	for _, query := range queries {
		switch q := query.(type) {
		case dsl.ANDLogicalDSL:
			builder.mongoBuilder.Comparison()
		case dsl.ORLogicalDSL:
		case dsl.NOTLogicalDSL:

		case dsl.CriteriaDSL:
			switch op := q.Operator {
			case dsl.EQUAL:
				//
			case dsl.EMPTY:
			case dsl.MATCHES:
			case dsl.STARTS_WITH:
			case dsl.ENDS_WITH:
			case dsl.LESS_THAN:
			case dsl.LESS_THAN_EQUAL:
			case dsl.GREATER_THAN:
			case dsl.GREATER_THAN_EQUAL:
			case dsl.IN:
			case dsl.RANGE:
			}
			doc = append(doc, )
		}
	}
}
