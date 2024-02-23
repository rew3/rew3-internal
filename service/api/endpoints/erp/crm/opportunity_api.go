package crm

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_OPPORTUNITY          endpoints.Endpoint = "crm_addOpportunity"
	UPDATE_OPPORTUNITY       endpoints.Endpoint = "crm_updateOpportunity"
	DELETE_OPPORTUNITY       endpoints.Endpoint = "crm_deleteOpportunity"
	CLONE_OPPORTUNITY        endpoints.Endpoint = "crm_cloneOpportunity"
	CHANGE_OPPORTUNITY_OWNER endpoints.Endpoint = "crm_changeOpportunityOwner"
	SET_OPPORTUNITY_FAVORITE endpoints.Endpoint = "crm_setOpportunityFavorite"

	// READ Operations.
	LIST_OPPORTUNITIES    endpoints.Endpoint = "crm_listOpportunities"
	COUNT_OPPORTUNITIES   endpoints.Endpoint = "crm_countOpportunities"
	GET_OPPORTUNITY_BY_ID endpoints.Endpoint = "crm_getOpportunityById"
)
