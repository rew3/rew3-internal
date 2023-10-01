package brokerage/acp

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for ACP
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_ACP                   api.APIOperation = "financial_addACP"
	UPDATE_ACP                 api.APIOperation = "financial_updateACP"
	DELETE_ACP                 api.APIOperation = "financial_deleteACP"
	CLONE_ACP                  api.APIOperation = "financial_cloneACP"
	CHANGE_ACP_OWNER           api.APIOperation = "financial_changeACPOwner"
	BULK_ADD_ACP          api.APIOperation = "financial_bulkAddACP"
	BULK_UPDATE_ACP           api.APIOperation = "financial_bulkUpdateACP"
	BULK_DELETE_ACP           api.APIOperation = "financial_bulkDeleteACP"
	BULK_CHANGE_ACP_OWNER           api.APIOperation = "financial_bulkChangeACPOwner"

	// READ APIs.
	LIST_ACPS     api.APIOperation = "financial_listACPs"
	COUNT_ACPS    api.APIOperation = "financial_countACPs"
	GET_ACP_BY_ID api.APIOperation = "cfinancial_getACPById"
    
)