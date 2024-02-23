package acp

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for SingleRateACP
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_SINGLERATEACP               endpoints.Endpoint = "financial_addSingleRateACP"
	UPDATE_SINGLERATEACP            endpoints.Endpoint = "financial_updateSingleRateACP"
	DELETE_SINGLERATEACP            endpoints.Endpoint = "financial_deleteSingleRateACP"
	CLONE_SINGLERATEACP             endpoints.Endpoint = "financial_cloneSingleRateACP"
	CHANGE_SINGLERATEACP_OWNER      endpoints.Endpoint = "financial_changeSingleRateACPOwner"
	BULK_ADD_SINGLERATEACP          endpoints.Endpoint = "financial_bulkAddSingleRateACP"
	BULK_UPDATE_SINGLERATEACP       endpoints.Endpoint = "financial_bulkUpdateSingleRateACP"
	BULK_DELETE_SINGLERATEACP       endpoints.Endpoint = "financial_bulkDeleteSingleRateACP"
	BULK_CHANGE_SINGLERATEACP_OWNER endpoints.Endpoint = "financial_bulkChangeSingleRateACPOwner"

	// READ APIs.
	LIST_SINGLERATEACPS     endpoints.Endpoint = "financial_listSingleRateACPs"
	COUNT_SINGLERATEACPS    endpoints.Endpoint = "financial_countSingleRateACPs"
	GET_SINGLERATEACP_BY_ID endpoints.Endpoint = "cfinancial_getSingleRateACPById"
)
