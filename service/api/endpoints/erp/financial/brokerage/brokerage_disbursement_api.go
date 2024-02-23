package brokerage

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Disbursement
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DISBURSEMENT               endpoints.Endpoint = "financial_addDisbursement"
	UPDATE_DISBURSEMENT            endpoints.Endpoint = "financial_updateDisbursement"
	DELETE_DISBURSEMENT            endpoints.Endpoint = "financial_deleteDisbursement"
	CLONE_DISBURSEMENT             endpoints.Endpoint = "financial_cloneDisbursement"
	CHANGE_DISBURSEMENT_OWNER      endpoints.Endpoint = "financial_changeDisbursementOwner"
	BULK_ADD_DISBURSEMENT          endpoints.Endpoint = "financial_bulkAddDisbursement"
	BULK_UPDATE_DISBURSEMENT       endpoints.Endpoint = "financial_bulkUpdateDisbursement"
	BULK_DELETE_DISBURSEMENT       endpoints.Endpoint = "financial_bulkDeleteDisbursement"
	BULK_CHANGE_DISBURSEMENT_OWNER endpoints.Endpoint = "financial_bulkChangeDisbursementOwner"

	// READ APIs.
	LIST_DISBURSEMENTS     endpoints.Endpoint = "financial_listDisbursements"
	COUNT_DISBURSEMENTS    endpoints.Endpoint = "financial_countDisbursements"
	GET_DISBURSEMENT_BY_ID endpoints.Endpoint = "cfinancial_getDisbursementById"
)
