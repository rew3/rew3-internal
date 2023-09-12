package crm

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_OPPORTUNITY                    api.APIOperation = "addOpportunity"
	UPDATE_OPPORTUNITY                 api.APIOperation = "updateOpportunity"
	DELETE_OPPORTUNITY                 api.APIOperation = "deleteOpportunity"
	CLONE_OPPORTUNITY                  api.APIOperation = "cloneOpportunity"
	CHANGE_OPPORTUNITY_OWNER           api.APIOperation = "changeOwner"
	CHANGE_OPPORTUNITY_FAVORITE_STATUS api.APIOperation = "changeFavoriteStatus"

	ADD_PIPELINE                    api.APIOperation = "addPipeline"
	UPDATE_PIPELINE                 api.APIOperation = "updatePipeline"
	DELETE_PIPELINE                 api.APIOperation = "deletePipeline"
	CLONE_PIPELINE                  api.APIOperation = "clonePipeline"
	CHANGE_PIPELINE_OWNER           api.APIOperation = "changeOwner"
	CHANGE_PIPELINE_FAVORITE_STATUS api.APIOperation = "changeFavoriteStatus"

	// READ Operations.
	LIST_OPPORTUNITIES    api.APIOperation = "listOpportunities"
	COUNT_OPPORTUNITIES   api.APIOperation = "countOpportunities"
	GET_OPPORTUNITY_BY_ID api.APIOperation = "getOpportunityById"

	LIST_PIPELINES     api.APIOperation = "listPipelines"
	COUNT_PIPELINES    api.APIOperation = "countPipelines"
	GET_PIPELINE_BY_ID api.APIOperation = "getPipelineById"
)
