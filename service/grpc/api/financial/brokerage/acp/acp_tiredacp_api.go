package acp

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for TiredACP
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_TIREDACP                   api.APIOperation = "financial_addTiredACP"
	UPDATE_TIREDACP                 api.APIOperation = "financial_updateTiredACP"
	DELETE_TIREDACP                 api.APIOperation = "financial_deleteTiredACP"
	CLONE_TIREDACP                  api.APIOperation = "financial_cloneTiredACP"
	CHANGE_TIREDACP_OWNER           api.APIOperation = "financial_changeTiredACPOwner"
	BULK_ADD_TIREDACP          api.APIOperation = "financial_bulkAddTiredACP"
	BULK_UPDATE_TIREDACP           api.APIOperation = "financial_bulkUpdateTiredACP"
	BULK_DELETE_TIREDACP           api.APIOperation = "financial_bulkDeleteTiredACP"
	BULK_CHANGE_TIREDACP_OWNER           api.APIOperation = "financial_bulkChangeTiredACPOwner"

	// READ APIs.
	LIST_TIREDACPS     api.APIOperation = "financial_listTiredACPs"
	COUNT_TIREDACPS    api.APIOperation = "financial_countTiredACPs"
	GET_TIREDACP_BY_ID api.APIOperation = "cfinancial_getTiredACPById"
    
)