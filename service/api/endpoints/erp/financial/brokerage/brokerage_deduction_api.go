package brokerage

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Deduction
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DEDUCTION               endpoints.Endpoint = "financial_addDeduction"
	UPDATE_DEDUCTION            endpoints.Endpoint = "financial_updateDeduction"
	DELETE_DEDUCTION            endpoints.Endpoint = "financial_deleteDeduction"
	CLONE_DEDUCTION             endpoints.Endpoint = "financial_cloneDeduction"
	CHANGE_DEDUCTION_OWNER      endpoints.Endpoint = "financial_changeDeductionOwner"
	BULK_ADD_DEDUCTION          endpoints.Endpoint = "financial_bulkAddDeduction"
	BULK_UPDATE_DEDUCTION       endpoints.Endpoint = "financial_bulkUpdateDeduction"
	BULK_DELETE_DEDUCTION       endpoints.Endpoint = "financial_bulkDeleteDeduction"
	BULK_CHANGE_DEDUCTION_OWNER endpoints.Endpoint = "financial_bulkChangeDeductionOwner"

	// READ APIs.
	LIST_DEDUCTIONS     endpoints.Endpoint = "financial_listDeductions"
	COUNT_DEDUCTIONS    endpoints.Endpoint = "financial_countDeductions"
	GET_DEDUCTION_BY_ID endpoints.Endpoint = "cfinancial_getDeductionById"
)
