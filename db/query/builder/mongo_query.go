package builder

import (
	"go.mongodb.org/mongo-driver/bson"
)

/*
 * Mongo Query Builder.
 */

type ComparisonOperator string
type LogicalOperator string
type ElementOperator string
type EvaluationOperator string

const (
	EQ  ComparisonOperator = "$eq"
	NE  ComparisonOperator = "$ne"
	LT  ComparisonOperator = "$lt"
	LTE ComparisonOperator = "$lte"
	GT  ComparisonOperator = "$gt"
	GTE ComparisonOperator = "$gte"
	IN  ComparisonOperator = "$in"
	NIN ComparisonOperator = "$nin"

	AND LogicalOperator = "$and"
	OR  LogicalOperator = "$or"
	NOT LogicalOperator = "$not"

	EXISTS ElementOperator = "$exists"

	REGEX EvaluationOperator = "$regex"
)

type MongoQueryBuilder struct{}

func NewMongoQueryBuilder() *MongoQueryBuilder {
	return &MongoQueryBuilder{}
}

/**
 * Create query for Comparison operator.
 */
func (qb *MongoQueryBuilder) Comparison(op ComparisonOperator, key string, value interface{}) bson.E {
	query := bson.E{Key: string(op), Value: value}
	return query
}

/**
 * Create query for Comparison operator with Negation i.e. $not.
 */
func (qb *MongoQueryBuilder) ComparisonNot(op ComparisonOperator, key string, value interface{}) bson.E {
	query := bson.E{Key: string(op), Value: value}
	return bson.E{Key: "$not", Value: query}
}

/**
 * Create query for Logical operator.
 */
func (qb *MongoQueryBuilder) Logical(op LogicalOperator, queries ...bson.E) bson.E {
	query := bson.E{Key: string(op), Value: bson.A{queries}}
	return query
}

/**
 * Create query for Element operator $exists.
 */
func (qb *MongoQueryBuilder) ElementExists(key string, value bool) bson.E {
	query := bson.E{Key: string(EXISTS), Value: value}
	return query
}

/**
 * Create query for Evaluation operator Regex.
 */
func (qb *MongoQueryBuilder) EvaluationRegex(key string, pattern string, caseInSensitive bool, isNegation bool) bson.E {
	if caseInSensitive {
		query := bson.E{
			Key: key,
			Value: bson.M{
				string(REGEX): pattern,
				"$options":    "i",
			},
		}
		if isNegation {
			return bson.E{Key: "$not", Value: query}
		} else {
			return query
		}
	} else {
		query := bson.E{
			Key:   key,
			Value: bson.E{Key: string(REGEX), Value: pattern},
		}
		if isNegation {
			return bson.E{Key: "$not", Value: query}
		} else {
			return query
		}
	}
}
