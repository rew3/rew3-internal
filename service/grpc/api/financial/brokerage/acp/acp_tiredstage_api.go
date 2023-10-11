package acp

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for TiredStage
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_TIREDSTAGE               api.APIOperation = "financial_addTiredStage"
	UPDATE_TIREDSTAGE            api.APIOperation = "financial_updateTiredStage"
	DELETE_TIREDSTAGE            api.APIOperation = "financial_deleteTiredStage"
	CLONE_TIREDSTAGE             api.APIOperation = "financial_cloneTiredStage"
	CHANGE_TIREDSTAGE_OWNER      api.APIOperation = "financial_changeTiredStageOwner"
	BULK_ADD_TIREDSTAGE          api.APIOperation = "financial_bulkAddTiredStage"
	BULK_UPDATE_TIREDSTAGE       api.APIOperation = "financial_bulkUpdateTiredStage"
	BULK_DELETE_TIREDSTAGE       api.APIOperation = "financial_bulkDeleteTiredStage"
	BULK_CHANGE_TIREDSTAGE_OWNER api.APIOperation = "financial_bulkChangeTiredStageOwner"

	// READ APIs.
	LIST_TIREDSTAGES     api.APIOperation = "financial_listTiredStages"
	COUNT_TIREDSTAGES    api.APIOperation = "financial_countTiredStages"
	GET_TIREDSTAGE_BY_ID api.APIOperation = "cfinancial_getTiredStageById"
)
