package brokerage

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for CommissionPlan
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_COMMISSIONPLAN               endpoints.Endpoint = "financial_addCommissionPlan"
	UPDATE_COMMISSIONPLAN            endpoints.Endpoint = "financial_updateCommissionPlan"
	DELETE_COMMISSIONPLAN            endpoints.Endpoint = "financial_deleteCommissionPlan"
	CLONE_COMMISSIONPLAN             endpoints.Endpoint = "financial_cloneCommissionPlan"
	CHANGE_COMMISSIONPLAN_OWNER      endpoints.Endpoint = "financial_changeCommissionPlanOwner"
	BULK_ADD_COMMISSIONPLAN          endpoints.Endpoint = "financial_bulkAddCommissionPlan"
	BULK_UPDATE_COMMISSIONPLAN       endpoints.Endpoint = "financial_bulkUpdateCommissionPlan"
	BULK_DELETE_COMMISSIONPLAN       endpoints.Endpoint = "financial_bulkDeleteCommissionPlan"
	BULK_CHANGE_COMMISSIONPLAN_OWNER endpoints.Endpoint = "financial_bulkChangeCommissionPlanOwner"

	// READ APIs.
	LIST_COMMISSIONPLANS     endpoints.Endpoint = "financial_listCommissionPlans"
	COUNT_COMMISSIONPLANS    endpoints.Endpoint = "financial_countCommissionPlans"
	GET_COMMISSIONPLAN_BY_ID endpoints.Endpoint = "cfinancial_getCommissionPlanById"
)
