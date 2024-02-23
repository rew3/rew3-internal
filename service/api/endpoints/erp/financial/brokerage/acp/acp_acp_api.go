package acp

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for ACP
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ACP               endpoints.Endpoint = "financial_addACP"
	UPDATE_ACP            endpoints.Endpoint = "financial_updateACP"
	DELETE_ACP            endpoints.Endpoint = "financial_deleteACP"
	CLONE_ACP             endpoints.Endpoint = "financial_cloneACP"
	CHANGE_ACP_OWNER      endpoints.Endpoint = "financial_changeACPOwner"
	BULK_ADD_ACP          endpoints.Endpoint = "financial_bulkAddACP"
	BULK_UPDATE_ACP       endpoints.Endpoint = "financial_bulkUpdateACP"
	BULK_DELETE_ACP       endpoints.Endpoint = "financial_bulkDeleteACP"
	BULK_CHANGE_ACP_OWNER endpoints.Endpoint = "financial_bulkChangeACPOwner"

	// READ APIs.
	LIST_ACPS     endpoints.Endpoint = "financial_listACPs"
	COUNT_ACPS    endpoints.Endpoint = "financial_countACPs"
	GET_ACP_BY_ID endpoints.Endpoint = "cfinancial_getACPById"
)
