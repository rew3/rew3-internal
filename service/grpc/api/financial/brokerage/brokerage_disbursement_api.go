package brokerage

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Disbursement
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_DISBURSEMENT                   api.APIOperation = "financial_addDisbursement"
	UPDATE_DISBURSEMENT                 api.APIOperation = "financial_updateDisbursement"
	DELETE_DISBURSEMENT                 api.APIOperation = "financial_deleteDisbursement"
	CLONE_DISBURSEMENT                  api.APIOperation = "financial_cloneDisbursement"
	CHANGE_DISBURSEMENT_OWNER           api.APIOperation = "financial_changeDisbursementOwner"
	BULK_ADD_DISBURSEMENT          api.APIOperation = "financial_bulkAddDisbursement"
	BULK_UPDATE_DISBURSEMENT           api.APIOperation = "financial_bulkUpdateDisbursement"
	BULK_DELETE_DISBURSEMENT           api.APIOperation = "financial_bulkDeleteDisbursement"
	BULK_CHANGE_DISBURSEMENT_OWNER           api.APIOperation = "financial_bulkChangeDisbursementOwner"

	// READ APIs.
	LIST_DISBURSEMENTS     api.APIOperation = "financial_listDisbursements"
	COUNT_DISBURSEMENTS    api.APIOperation = "financial_countDisbursements"
	GET_DISBURSEMENT_BY_ID api.APIOperation = "cfinancial_getDisbursementById"
    
)