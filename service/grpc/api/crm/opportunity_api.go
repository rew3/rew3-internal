package crm

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_OPPORTUNITY          api.APIOperation = "crm_addOpportunity"
	UPDATE_OPPORTUNITY       api.APIOperation = "crm_updateOpportunity"
	DELETE_OPPORTUNITY       api.APIOperation = "crm_deleteOpportunity"
	CLONE_OPPORTUNITY        api.APIOperation = "crm_cloneOpportunity"
	CHANGE_OPPORTUNITY_OWNER api.APIOperation = "crm_changeOpportunityOwner"
	SET_OPPORTUNITY_FAVORITE api.APIOperation = "crm_setOpportunityFavorite"

	// READ Operations.
	LIST_OPPORTUNITIES    api.APIOperation = "crm_listOpportunities"
	COUNT_OPPORTUNITIES   api.APIOperation = "crm_countOpportunities"
	GET_OPPORTUNITY_BY_ID api.APIOperation = "crm_getOpportunityById"
)
