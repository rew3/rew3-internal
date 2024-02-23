package acp

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for TiredACP
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_TIREDACP               endpoints.Endpoint = "financial_addTiredACP"
	UPDATE_TIREDACP            endpoints.Endpoint = "financial_updateTiredACP"
	DELETE_TIREDACP            endpoints.Endpoint = "financial_deleteTiredACP"
	CLONE_TIREDACP             endpoints.Endpoint = "financial_cloneTiredACP"
	CHANGE_TIREDACP_OWNER      endpoints.Endpoint = "financial_changeTiredACPOwner"
	BULK_ADD_TIREDACP          endpoints.Endpoint = "financial_bulkAddTiredACP"
	BULK_UPDATE_TIREDACP       endpoints.Endpoint = "financial_bulkUpdateTiredACP"
	BULK_DELETE_TIREDACP       endpoints.Endpoint = "financial_bulkDeleteTiredACP"
	BULK_CHANGE_TIREDACP_OWNER endpoints.Endpoint = "financial_bulkChangeTiredACPOwner"

	// READ APIs.
	LIST_TIREDACPS     endpoints.Endpoint = "financial_listTiredACPs"
	COUNT_TIREDACPS    endpoints.Endpoint = "financial_countTiredACPs"
	GET_TIREDACP_BY_ID endpoints.Endpoint = "cfinancial_getTiredACPById"
)
