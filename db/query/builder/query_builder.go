package builder

import (
	"fmt"

	dsl "github.com/rew3/rew3-internal/db/query"
	"go.mongodb.org/mongo-driver/bson"
)

/**
 * Query Builder.
 */
type QueryBuilder struct {
	queryDSL     dsl.QueryDSL
	mongoBuilder *MongoQueryBuilder
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{mongoBuilder: NewMongoQueryBuilder()}
}

/**
 * Set the Query DSL for generating query.
 */
func (builder *QueryBuilder) ForQuery(query dsl.QueryDSL) {
	builder.queryDSL = query
}

/**
 * Generate Mongo Speficic query.
 */
func (builder *QueryBuilder) Mongo() (bson.D, error) {
	queries := builder.queryDSL.Queries
	return builder.generateMongoQuery(queries)
}

/**
 * Generate query for MongoDB.
 */
func (builder *QueryBuilder) generateMongoQuery(queries []dsl.BaseDSL) (bson.D, error) {
	// Generate a comparison query for given criteria.
	comparisonQuery := func(op ComparisonOperator, field dsl.FieldDSL, value dsl.ValueType) bson.E {
		if field.IsNegation {
			return builder.mongoBuilder.ComparisonNot(op, field.Name, value)
		} else {
			return builder.mongoBuilder.Comparison(op, field.Name, value)
		}
	}
	// Generate a regex query. throw error if value is not string type. i.e. it will only works with string values.
	regexQuery := func(field dsl.FieldDSL, value dsl.ValueType, regex func(string) string, opName string) (bson.E, error) {
		switch v := value.(type) {
		case dsl.StringValue:
			if field.IsNegation {
				return builder.mongoBuilder.EvaluationRegex(field.Name, regex(v.Value), true, true), nil
			} else {
				return builder.mongoBuilder.EvaluationRegex(field.Name, regex(v.Value), true, false), nil
			}
		default:
			err := fmt.Errorf("query error: DSL operator `%s` require string value type for field: %s", opName, field.Name)
			return bson.E{}, err
		}
	}
	doc := bson.D{}
	for _, query := range queries {
		switch q := query.(type) {
		case dsl.AssocDSL:
			for _, i := range q.Queries {
				switch i.(type) {
				case dsl.LogicalDSL:
					res, err := builder.generateMongoQuery([]dsl.BaseDSL{i})
					if err != nil {
						return nil, err
					}
					doc = append(doc, res...)
				case dsl.CriteriaDSL:
					res, err := builder.generateMongoQuery([]dsl.BaseDSL{i})
					if err != nil {
						return nil, err
					}
					doc = append(doc, res...)
				}
			}
		case dsl.ANDLogicalDSL:
			criteriaForAND := []bson.D{}
			for _, i := range q.Queries {
				res, err := builder.generateMongoQuery([]dsl.BaseDSL{i})
				if err != nil {
					return nil, err
				}
				criteriaForAND = append(criteriaForAND, res)
			}
			f := bson.D{{Key: string(AND), Value: criteriaForAND}}
			doc = append(doc, f...)
		case dsl.ORLogicalDSL:
			criteriaForOR := []bson.D{}
			for _, i := range q.Queries {
				res, err := builder.generateMongoQuery([]dsl.BaseDSL{i})
				if err != nil {
					return nil, err
				}
				criteriaForOR = append(criteriaForOR, res)
			}
			f := bson.D{{Key: string(OR), Value: criteriaForOR}}
			doc = append(doc, f...)
		case dsl.NOTLogicalDSL:
			res, err := builder.generateMongoQuery([]dsl.BaseDSL{q.Query})
			if err != nil {
				return nil, err
			}
			f := bson.D{{Key: string(OR), Value: res}}
			doc = append(doc, f...)
		case dsl.CriteriaDSL:
			// Check for scalar value end execute callback if value is scalar.
			checkScalarValueAndThen := func(callback func()) error {
				if dsl.IsScalarValue(q.Value) {
					callback()
					return nil
				} else {
					return fmt.Errorf("query error: DSL operator `%s` require scalar value (string, bool, number, datetime) type for field: %s", string(q.Operator), q.Field.Name)
				}
			}
			switch q.Operator {
			case dsl.UNKNOWN:
				err := checkScalarValueAndThen(func() {
					if q.Field.IsNegation {
						doc = append(doc, builder.mongoBuilder.RegularNot(q.Field.Name, dsl.ResolveScalarValue(q.Value)))
					} else {
						doc = append(doc, builder.mongoBuilder.Regular(q.Field.Name, dsl.ResolveScalarValue(q.Value)))
					}
				})
				if err != nil {
					return nil, err
				}
			case dsl.EQUAL:
				err := checkScalarValueAndThen(func() {
					if q.Field.IsNegation {
						doc = append(doc, builder.mongoBuilder.Comparison(NE, q.Field.Name, dsl.ResolveScalarValue(q.Value)))
					} else {
						doc = append(doc, builder.mongoBuilder.Comparison(EQ, q.Field.Name, dsl.ResolveScalarValue(q.Value)))
					}
				})
				if err != nil {
					return nil, err
				}
			case dsl.EMPTY:
				if q.Field.IsNegation {
					doc = append(doc, builder.mongoBuilder.ElementExists(q.Field.Name, false))
				} else {
					doc = append(doc, builder.mongoBuilder.ElementExists(q.Field.Name, true))
				}
			case dsl.MATCHES:
				res, err := regexQuery(q.Field, q.Value, func(str string) string { return ".*" + str + ".*" }, "MATCHES")
				if err != nil {
					return nil, err
				}
				doc = append(doc, res)
			case dsl.STARTS_WITH:
				res, err := regexQuery(q.Field, q.Value, func(str string) string { return "^" + str + ".*" }, "STARTS_WITH")
				if err != nil {
					return nil, err
				}
				doc = append(doc, res)
			case dsl.ENDS_WITH:
				res, err := regexQuery(q.Field, q.Value, func(str string) string { return ".*" + str + "$" }, "ENDS_WITH")
				if err != nil {
					return nil, err
				}
				doc = append(doc, res)
			case dsl.LESS_THAN:
				err := checkScalarValueAndThen(func() {
					doc = append(doc, comparisonQuery(LT, q.Field, dsl.ResolveScalarValue(q.Value)))
				})
				if err != nil {
					return nil, err
				}
			case dsl.LESS_THAN_EQUAL:
				err := checkScalarValueAndThen(func() {
					doc = append(doc, comparisonQuery(LTE, q.Field, dsl.ResolveScalarValue(q.Value)))
				})
				if err != nil {
					return nil, err
				}
			case dsl.GREATER_THAN:
				err := checkScalarValueAndThen(func() {
					doc = append(doc, comparisonQuery(GT, q.Field, dsl.ResolveScalarValue(q.Value)))
				})
				if err != nil {
					return nil, err
				}
			case dsl.GREATER_THAN_EQUAL:
				err := checkScalarValueAndThen(func() {
					doc = append(doc, comparisonQuery(GTE, q.Field, dsl.ResolveScalarValue(q.Value)))
				})
				if err != nil {
					return nil, err
				}
			case dsl.IN:
				switch v := q.Value.(type) {
				case dsl.ListValue:
					items := []interface{}{}
					for _, i := range v.Value {
						items = append(items, dsl.ResolveScalarValue(i))
					}
					if q.Field.IsNegation {
						doc = append(doc, builder.mongoBuilder.Comparison(NIN, q.Field.Name, items))
					} else {
						doc = append(doc, builder.mongoBuilder.Comparison(IN, q.Field.Name, items))
					}
				default:
					err := fmt.Errorf("query error: DSL operator `IN` require list value type for field: %s", q.Field.Name)
					return nil, err
				}
			case dsl.RANGE:
				switch v := q.Value.(type) {
				case dsl.RangeValue:
					gt := comparisonQuery(GTE, q.Field, dsl.ResolveScalarValue(v.Start))
					lt := comparisonQuery(LTE, q.Field, dsl.ResolveScalarValue(v.End))
					doc = append(doc, bson.E{Key: string(AND), Value: bson.A{gt, lt}})
				default:
					err := fmt.Errorf("query error: DSL operator `RANGE` require start and end value for field: %s", q.Field.Name)
					return nil, err
				}
			}
		}
	}
	return doc, nil
}
