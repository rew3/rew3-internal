package brokerage

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Deduction
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_DEDUCTION                   api.APIOperation = "financial_addDeduction"
	UPDATE_DEDUCTION                 api.APIOperation = "financial_updateDeduction"
	DELETE_DEDUCTION                 api.APIOperation = "financial_deleteDeduction"
	CLONE_DEDUCTION                  api.APIOperation = "financial_cloneDeduction"
	CHANGE_DEDUCTION_OWNER           api.APIOperation = "financial_changeDeductionOwner"
	BULK_ADD_DEDUCTION          api.APIOperation = "financial_bulkAddDeduction"
	BULK_UPDATE_DEDUCTION           api.APIOperation = "financial_bulkUpdateDeduction"
	BULK_DELETE_DEDUCTION           api.APIOperation = "financial_bulkDeleteDeduction"
	BULK_CHANGE_DEDUCTION_OWNER           api.APIOperation = "financial_bulkChangeDeductionOwner"

	// READ APIs.
	LIST_DEDUCTIONS     api.APIOperation = "financial_listDeductions"
	COUNT_DEDUCTIONS    api.APIOperation = "financial_countDeductions"
	GET_DEDUCTION_BY_ID api.APIOperation = "cfinancial_getDeductionById"
    
)