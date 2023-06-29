package filter

import (
	dsl "github.com/rew3/rew3-internal/db/query"
)

/**
 * Generate Query from given filters.
 */
func GenerateQueryDSL(filters []Filter) dsl.RootDSL {
	queries := []dsl.QueryDSL{}
	for _, filter := range filters {
		query := resolveQuery(filter)
		queries = append(queries, query)
	}
	return dsl.RootDSL{Queries: queries}
}

/**
 * Resolve Query for given filter.
 */
func resolveQuery(filter Filter) dsl.QueryDSL {
	switch filter.Operator {
	case EQ:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case NOT_EQ:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case EMPTY:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case NOT_EMPTY:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case CONTAINS:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case NOT_CONTAINS:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case STARTS_WITH:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case NOT_STARTS_WITH:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case ENDS_WITH:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case NOT_ENDS_WITH:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case LESS_THAN:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case NOT_LESS_THAN:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case LESS_THAN_EQUAL:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case NOT_LESS_THAN_EQUAL:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case GREATER_THAN:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case NOT_GREATER_THAN:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case GREATER_THAN_EQUAL:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case NOT_GREATER_THAN_EQUAL:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case IN:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case NOT_IN:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case RANGE:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case NOT_IN_RANGE:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	default:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.UNKNOWN)
	}
}
