package filter

import (
	dsl "github.com/rew3/rew3-internal/db/query"
)

/**
 * Generate Query from given filters.
 */
func GenerateQueryDSL(filters []Filter) dsl.QueryDSL {
	queries := []dsl.BaseDSL{}
	for _, filter := range filters {
		query := resolveQuery(filter)
		queries = append(queries, query)
	}
	return dsl.QueryDSL{Queries: queries}
}

/**
 * Resolve Query for given filter.
 */
func resolveQuery(filter Filter) dsl.BaseDSL {
	switch filter.Operator {
	case EQ:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case NOT_EQ:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.EQUAL)
	case EMPTY:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EMPTY)
	case NOT_EMPTY:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.EMPTY)
	case CONTAINS:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.MATCHES)
	case NOT_CONTAINS:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.MATCHES)
	case STARTS_WITH:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.STARTS_WITH)
	case NOT_STARTS_WITH:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.STARTS_WITH)
	case ENDS_WITH:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.ENDS_WITH)
	case NOT_ENDS_WITH:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.ENDS_WITH)
	case LESS_THAN:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.LESS_THAN)
	case NOT_LESS_THAN:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.LESS_THAN)
	case LESS_THAN_EQUAL:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.LESS_THAN_EQUAL)
	case NOT_LESS_THAN_EQUAL:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.LESS_THAN_EQUAL)
	case GREATER_THAN:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.GREATER_THAN)
	case NOT_GREATER_THAN:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.GREATER_THAN)
	case GREATER_THAN_EQUAL:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.GREATER_THAN_EQUAL)
	case NOT_GREATER_THAN_EQUAL:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.GREATER_THAN_EQUAL)
	case IN:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.IN)
	case NOT_IN:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.IN)
	case RANGE:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.RANGE)
	case NOT_IN_RANGE:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.RANGE)
	default:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.UNKNOWN)
	}
}
