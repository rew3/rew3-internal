package acp

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for TiredStage
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_TIREDSTAGE               endpoints.Endpoint = "financial_addTiredStage"
	UPDATE_TIREDSTAGE            endpoints.Endpoint = "financial_updateTiredStage"
	DELETE_TIREDSTAGE            endpoints.Endpoint = "financial_deleteTiredStage"
	CLONE_TIREDSTAGE             endpoints.Endpoint = "financial_cloneTiredStage"
	CHANGE_TIREDSTAGE_OWNER      endpoints.Endpoint = "financial_changeTiredStageOwner"
	BULK_ADD_TIREDSTAGE          endpoints.Endpoint = "financial_bulkAddTiredStage"
	BULK_UPDATE_TIREDSTAGE       endpoints.Endpoint = "financial_bulkUpdateTiredStage"
	BULK_DELETE_TIREDSTAGE       endpoints.Endpoint = "financial_bulkDeleteTiredStage"
	BULK_CHANGE_TIREDSTAGE_OWNER endpoints.Endpoint = "financial_bulkChangeTiredStageOwner"

	// READ APIs.
	LIST_TIREDSTAGES     endpoints.Endpoint = "financial_listTiredStages"
	COUNT_TIREDSTAGES    endpoints.Endpoint = "financial_countTiredStages"
	GET_TIREDSTAGE_BY_ID endpoints.Endpoint = "cfinancial_getTiredStageById"
)
