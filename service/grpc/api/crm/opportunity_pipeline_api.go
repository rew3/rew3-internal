package crm

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_OPPORTUNITY_PIPELINE                    api.APIOperation = "addOpportunityPipeline"
	UPDATE_OPPORTUNITY_PIPELINE                 api.APIOperation = "updateOpportunityPipeline"
	DELETE_OPPORTUNITY_PIPELINE                 api.APIOperation = "deleteOpportunityPipeline"
	CLONE_OPPORTUNITY_PIPELINE                  api.APIOperation = "cloneOpportunityPipeline"
	CHANGE_OPPORTUNITY_PIPELINE_OWNER           api.APIOperation = "changeOpportunityPipelineOwner"
	CHANGE_OPPORTUNITY_PIPELINE_FAVORITE_STATUS api.APIOperation = "changeOpportunityPipelineFavoriteStatus"

	// READ Operations.
	LIST_OPPORTUNITY_PIPELINES     api.APIOperation = "listOpportunityPipelines"
	COUNT_OPPORTUNITY_PIPELINES    api.APIOperation = "countOpportunityPipelines"
	GET_OPPORTUNITY_PIPELINE_BY_ID api.APIOperation = "getOpportunityPipelineById"
)
