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
 * Create regular query.
 * {first_name: value}
 */
func (qb *MongoQueryBuilder) Regular(key string, value interface{}) bson.D {
	query := bson.D{{Key: key, Value: value}}
	return query
}

/**
 * Create regular query.
 * {$not: {first_name: value}}
 */
func (qb *MongoQueryBuilder) RegularNot(key string, value interface{}) bson.D {
	//query := bson.M{key: value}
	return bson.D{
		{Key: key, Value: bson.D{
			{Key: "$not", Value: value},
		}},
	}
}

/**
 * Create query for Comparison operator.
 */
func (qb *MongoQueryBuilder) Comparison(op ComparisonOperator, key string, value interface{}) bson.D {
	query := bson.D{{Key: key, Value: bson.M{string(op): value}}}
	return query
}

/**
 * Create query for Comparison operator with Negation i.e. $not.
 */
func (qb *MongoQueryBuilder) ComparisonNot(op ComparisonOperator, key string, value interface{}) bson.D {
	//query := bson.M{key: bson.M{string(op): value}}
	//return bson.D{{Key: "$not", Value: query}}
	return bson.D{
		{Key: key, Value: bson.D{
			{Key: "$not", Value: bson.D{
				{Key: string(op), Value: value},
			}},
		}},
	}
}

/**
 * Create query for Logical operator.
 */
func (qb *MongoQueryBuilder) Logical(op LogicalOperator, queries ...bson.D) bson.D {
	query := bson.D{{Key: string(op), Value: queries}}
	return query
}

/**
 * Create query for Element operator $exists.
 */
func (qb *MongoQueryBuilder) ElementExists(key string, value bool) bson.D {
	if value {
		return bson.D{
			{Key: "$or", Value: bson.A{
				bson.D{{Key: key, Value: bson.M{string(EXISTS): false}}},
				bson.D{{Key: key, Value: bson.D{{Key: "$eq", Value: nil}}}},
				bson.D{{Key: key, Value: bson.D{{Key: "$eq", Value: ""}}}},
				bson.D{{Key: key, Value: bson.D{{Key: "$eq", Value: 0}}}},
			}},
		}
	} else {
		return bson.D{
			{Key: key, Value: bson.M{string(EXISTS): true}},
			{Key: "$or", Value: bson.A{
				bson.D{{Key: key, Value: bson.D{{Key: "$ne", Value: nil}}}},
				bson.D{{Key: key, Value: bson.D{{Key: "$ne", Value: ""}}}},
				bson.D{{Key: key, Value: bson.D{{Key: "$ne", Value: 0}}}},
			}},
		}
	}
}

/**
 * Create query for Evaluation operator Regex.
 */
func (qb *MongoQueryBuilder) EvaluationRegex(key string, pattern string, caseInSensitive bool, isNegation bool) bson.D {
	if caseInSensitive {
		query := bson.D{{
			Key: key,
			Value: bson.M{
				string(REGEX): pattern,
				"$options":    "i",
			},
		}}
		if isNegation {
			return bson.D{{Key: "$nor", Value: bson.A{query}}}
		} else {
			return query
		}
	} else {
		query := bson.D{{
			Key:   key,
			Value: bson.M{string(REGEX): pattern},
		}}
		if isNegation {
			return bson.D{{Key: "$nor", Value: bson.A{query}}}
		} else {
			return query
		}
	}
}
