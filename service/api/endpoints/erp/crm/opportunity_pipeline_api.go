package crm

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_OPPORTUNITY_PIPELINE          endpoints.Endpoint = "crm_addOpportunityPipeline"
	UPDATE_OPPORTUNITY_PIPELINE       endpoints.Endpoint = "crm_updateOpportunityPipeline"
	DELETE_OPPORTUNITY_PIPELINE       endpoints.Endpoint = "crm_deleteOpportunityPipeline"
	CLONE_OPPORTUNITY_PIPELINE        endpoints.Endpoint = "crm_cloneOpportunityPipeline"
	CHANGE_OPPORTUNITY_PIPELINE_OWNER endpoints.Endpoint = "crm_changeOpportunityPipelineOwner"
	SET_OPPORTUNITY_PIPELINE_FAVORITE endpoints.Endpoint = "crm_setOpportunityPipelineFavorite"

	// READ Operations.
	LIST_OPPORTUNITY_PIPELINES     endpoints.Endpoint = "crm_listOpportunityPipelines"
	COUNT_OPPORTUNITY_PIPELINES    endpoints.Endpoint = "crm_countOpportunityPipelines"
	GET_OPPORTUNITY_PIPELINE_BY_ID endpoints.Endpoint = "crm_getOpportunityPipelineById"
)
