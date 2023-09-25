package crm

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_OPPORTUNITY_PIPELINE                    api.APIOperation = "crm_addOpportunityPipeline"
	UPDATE_OPPORTUNITY_PIPELINE                 api.APIOperation = "crm_updateOpportunityPipeline"
	DELETE_OPPORTUNITY_PIPELINE                 api.APIOperation = "crm_deleteOpportunityPipeline"
	CLONE_OPPORTUNITY_PIPELINE                  api.APIOperation = "crm_cloneOpportunityPipeline"
	CHANGE_OPPORTUNITY_PIPELINE_OWNER           api.APIOperation = "crm_changeOpportunityPipelineOwner"
	CHANGE_OPPORTUNITY_PIPELINE_FAVORITE_STATUS api.APIOperation = "crm_changeOpportunityPipelineFavoriteStatus"

	// READ Operations.
	LIST_OPPORTUNITY_PIPELINES     api.APIOperation = "crm_listOpportunityPipelines"
	COUNT_OPPORTUNITY_PIPELINES    api.APIOperation = "crm_countOpportunityPipelines"
	GET_OPPORTUNITY_PIPELINE_BY_ID api.APIOperation = "crm_getOpportunityPipelineById"
)
