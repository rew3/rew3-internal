package brokerage

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for CommissionPlan
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_COMMISSIONPLAN               api.APIOperation = "financial_addCommissionPlan"
	UPDATE_COMMISSIONPLAN            api.APIOperation = "financial_updateCommissionPlan"
	DELETE_COMMISSIONPLAN            api.APIOperation = "financial_deleteCommissionPlan"
	CLONE_COMMISSIONPLAN             api.APIOperation = "financial_cloneCommissionPlan"
	CHANGE_COMMISSIONPLAN_OWNER      api.APIOperation = "financial_changeCommissionPlanOwner"
	BULK_ADD_COMMISSIONPLAN          api.APIOperation = "financial_bulkAddCommissionPlan"
	BULK_UPDATE_COMMISSIONPLAN       api.APIOperation = "financial_bulkUpdateCommissionPlan"
	BULK_DELETE_COMMISSIONPLAN       api.APIOperation = "financial_bulkDeleteCommissionPlan"
	BULK_CHANGE_COMMISSIONPLAN_OWNER api.APIOperation = "financial_bulkChangeCommissionPlanOwner"

	// READ APIs.
	LIST_COMMISSIONPLANS     api.APIOperation = "financial_listCommissionPlans"
	COUNT_COMMISSIONPLANS    api.APIOperation = "financial_countCommissionPlans"
	GET_COMMISSIONPLAN_BY_ID api.APIOperation = "cfinancial_getCommissionPlanById"
)
