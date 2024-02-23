package brokerage

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for GrossCommissionPlan
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_GROSSCOMMISSIONPLAN               endpoints.Endpoint = "financial_addGrossCommissionPlan"
	UPDATE_GROSSCOMMISSIONPLAN            endpoints.Endpoint = "financial_updateGrossCommissionPlan"
	DELETE_GROSSCOMMISSIONPLAN            endpoints.Endpoint = "financial_deleteGrossCommissionPlan"
	CLONE_GROSSCOMMISSIONPLAN             endpoints.Endpoint = "financial_cloneGrossCommissionPlan"
	CHANGE_GROSSCOMMISSIONPLAN_OWNER      endpoints.Endpoint = "financial_changeGrossCommissionPlanOwner"
	BULK_ADD_GROSSCOMMISSIONPLAN          endpoints.Endpoint = "financial_bulkAddGrossCommissionPlan"
	BULK_UPDATE_GROSSCOMMISSIONPLAN       endpoints.Endpoint = "financial_bulkUpdateGrossCommissionPlan"
	BULK_DELETE_GROSSCOMMISSIONPLAN       endpoints.Endpoint = "financial_bulkDeleteGrossCommissionPlan"
	BULK_CHANGE_GROSSCOMMISSIONPLAN_OWNER endpoints.Endpoint = "financial_bulkChangeGrossCommissionPlanOwner"

	// READ APIs.
	LIST_GROSSCOMMISSIONPLANS     endpoints.Endpoint = "financial_listGrossCommissionPlans"
	COUNT_GROSSCOMMISSIONPLANS    endpoints.Endpoint = "financial_countGrossCommissionPlans"
	GET_GROSSCOMMISSIONPLAN_BY_ID endpoints.Endpoint = "cfinancial_getGrossCommissionPlanById"
)
