package filter

import (
	"github.com/rew3/rew3-pkg/core/service/shared/request"
	dsl "github.com/rew3/rew3-pkg/db/query"
	"github.com/rew3/rew3-pkg/utils/logger"
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
 * Resolve given request parameter and generate query dsl.
 * Note: if request param filter is invalid, empty dsl is returned.
 */
func ResolveAndGenerateQueryDSL(param request.RequestParam) dsl.QueryDSL {
	filters, err := ParseFilterParams(param)
	if err != nil {
		logger.Log().Errorln("Unable to parse request parameter while generating Query DSL: ", err.Error())
		return dsl.Query()
	}
	return GenerateQueryDSL(filters)
}

/**
 * Resolve Query for given filter.
 */
func resolveQuery(filter Filter) dsl.BaseDSL {
	switch filter.Operator {
	case IS:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case IS_NOT:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.EQUAL)
	case EMPTY:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EMPTY)
	case NOT_EMPTY:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.EMPTY)
	case IN:
		{
			if value, ok := filter.Value.([]interface{}); ok {
				return dsl.CriteriaForIn(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, value)
			} else {
				return dsl.CriteriaForIn(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, nil)
			}
		}
	case NOT_IN:
		{
			if value, ok := filter.Value.([]interface{}); ok {
				return dsl.CriteriaForIn(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, value)
			} else {
				return dsl.CriteriaForIn(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, nil)
			}
		}
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
	case EQ:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.EQUAL)
	case NOT_EQ:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, filter.Value, dsl.EQUAL)
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
	case RANGE:
		{
			if value, ok := filter.Value.([2]interface{}); ok {
				return dsl.CriteriaForRange(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, value[0], value[1])
			} else {
				return dsl.CriteriaForRange(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, nil, nil)
			}
		}
	case NOT_IN_RANGE:
		{
			if value, ok := filter.Value.([2]interface{}); ok {
				return dsl.CriteriaForRange(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, value[0], value[1])
			} else {
				return dsl.CriteriaForRange(dsl.FieldDSL{Name: filter.Field, IsNegation: true}, nil, nil)
			}
		}
	default:
		return dsl.Criteria(dsl.FieldDSL{Name: filter.Field, IsNegation: false}, filter.Value, dsl.UNKNOWN)
	}
}
