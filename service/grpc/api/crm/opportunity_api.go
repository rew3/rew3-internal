package crm

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_OPPORTUNITY                    api.APIOperation = "addOpportunity"
	UPDATE_OPPORTUNITY                 api.APIOperation = "updateOpportunity"
	DELETE_OPPORTUNITY                 api.APIOperation = "deleteOpportunity"
	CLONE_OPPORTUNITY                  api.APIOperation = "cloneOpportunity"
	CHANGE_OPPORTUNITY_OWNER           api.APIOperation = "changeOpportunityOwner"
	CHANGE_OPPORTUNITY_FAVORITE_STATUS api.APIOperation = "changeOpportunityFavoriteStatus"

	// READ Operations.
	LIST_OPPORTUNITIES    api.APIOperation = "listOpportunities"
	COUNT_OPPORTUNITIES   api.APIOperation = "countOpportunities"
	GET_OPPORTUNITY_BY_ID api.APIOperation = "getOpportunityById"
)
