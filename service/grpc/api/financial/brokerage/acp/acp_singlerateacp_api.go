package acp

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for SingleRateACP
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_SINGLERATEACP                   api.APIOperation = "financial_addSingleRateACP"
	UPDATE_SINGLERATEACP                 api.APIOperation = "financial_updateSingleRateACP"
	DELETE_SINGLERATEACP                 api.APIOperation = "financial_deleteSingleRateACP"
	CLONE_SINGLERATEACP                  api.APIOperation = "financial_cloneSingleRateACP"
	CHANGE_SINGLERATEACP_OWNER           api.APIOperation = "financial_changeSingleRateACPOwner"
	BULK_ADD_SINGLERATEACP          api.APIOperation = "financial_bulkAddSingleRateACP"
	BULK_UPDATE_SINGLERATEACP           api.APIOperation = "financial_bulkUpdateSingleRateACP"
	BULK_DELETE_SINGLERATEACP           api.APIOperation = "financial_bulkDeleteSingleRateACP"
	BULK_CHANGE_SINGLERATEACP_OWNER           api.APIOperation = "financial_bulkChangeSingleRateACPOwner"

	// READ APIs.
	LIST_SINGLERATEACPS     api.APIOperation = "financial_listSingleRateACPs"
	COUNT_SINGLERATEACPS    api.APIOperation = "financial_countSingleRateACPs"
	GET_SINGLERATEACP_BY_ID api.APIOperation = "cfinancial_getSingleRateACPById"
    
)