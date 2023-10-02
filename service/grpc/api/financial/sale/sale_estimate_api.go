package sale

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Estimate
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_ESTIMATE                   api.APIOperation = "financial_addEstimate"
	UPDATE_ESTIMATE                 api.APIOperation = "financial_updateEstimate"
	DELETE_ESTIMATE                 api.APIOperation = "financial_deleteEstimate"
	CLONE_ESTIMATE                  api.APIOperation = "financial_cloneEstimate"
	CHANGE_ESTIMATE_OWNER           api.APIOperation = "financial_changeEstimateOwner"
	BULK_ADD_ESTIMATE          api.APIOperation = "financial_bulkAddEstimate"
	BULK_UPDATE_ESTIMATE           api.APIOperation = "financial_bulkUpdateEstimate"
	BULK_DELETE_ESTIMATE           api.APIOperation = "financial_bulkDeleteEstimate"
	BULK_CHANGE_ESTIMATE_OWNER           api.APIOperation = "financial_bulkChangeEstimateOwner"

	// READ APIs.
	LIST_ESTIMATES     api.APIOperation = "financial_listEstimates"
	COUNT_ESTIMATES    api.APIOperation = "financial_countEstimates"
	GET_ESTIMATE_BY_ID api.APIOperation = "cfinancial_getEstimateById"
    
)