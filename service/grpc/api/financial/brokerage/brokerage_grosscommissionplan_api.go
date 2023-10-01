package brokerage

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for GrossCommissionPlan
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_GROSSCOMMISSIONPLAN               api.APIOperation = "financial_addGrossCommissionPlan"
	UPDATE_GROSSCOMMISSIONPLAN            api.APIOperation = "financial_updateGrossCommissionPlan"
	DELETE_GROSSCOMMISSIONPLAN            api.APIOperation = "financial_deleteGrossCommissionPlan"
	CLONE_GROSSCOMMISSIONPLAN             api.APIOperation = "financial_cloneGrossCommissionPlan"
	CHANGE_GROSSCOMMISSIONPLAN_OWNER      api.APIOperation = "financial_changeGrossCommissionPlanOwner"
	BULK_ADD_GROSSCOMMISSIONPLAN          api.APIOperation = "financial_bulkAddGrossCommissionPlan"
	BULK_UPDATE_GROSSCOMMISSIONPLAN       api.APIOperation = "financial_bulkUpdateGrossCommissionPlan"
	BULK_DELETE_GROSSCOMMISSIONPLAN       api.APIOperation = "financial_bulkDeleteGrossCommissionPlan"
	BULK_CHANGE_GROSSCOMMISSIONPLAN_OWNER api.APIOperation = "financial_bulkChangeGrossCommissionPlanOwner"

	// READ APIs.
	LIST_GROSSCOMMISSIONPLANS     api.APIOperation = "financial_listGrossCommissionPlans"
	COUNT_GROSSCOMMISSIONPLANS    api.APIOperation = "financial_countGrossCommissionPlans"
	GET_GROSSCOMMISSIONPLAN_BY_ID api.APIOperation = "cfinancial_getGrossCommissionPlanById"
)
