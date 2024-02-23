package sale

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Estimate
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ESTIMATE               endpoints.Endpoint = "financial_addEstimate"
	UPDATE_ESTIMATE            endpoints.Endpoint = "financial_updateEstimate"
	DELETE_ESTIMATE            endpoints.Endpoint = "financial_deleteEstimate"
	CLONE_ESTIMATE             endpoints.Endpoint = "financial_cloneEstimate"
	CHANGE_ESTIMATE_OWNER      endpoints.Endpoint = "financial_changeEstimateOwner"
	BULK_ADD_ESTIMATE          endpoints.Endpoint = "financial_bulkAddEstimate"
	BULK_UPDATE_ESTIMATE       endpoints.Endpoint = "financial_bulkUpdateEstimate"
	BULK_DELETE_ESTIMATE       endpoints.Endpoint = "financial_bulkDeleteEstimate"
	BULK_CHANGE_ESTIMATE_OWNER endpoints.Endpoint = "financial_bulkChangeEstimateOwner"

	// READ APIs.
	LIST_ESTIMATES     endpoints.Endpoint = "financial_listEstimates"
	COUNT_ESTIMATES    endpoints.Endpoint = "financial_countEstimates"
	GET_ESTIMATE_BY_ID endpoints.Endpoint = "cfinancial_getEstimateById"
)
