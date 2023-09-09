package crm

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_CAMPAIGN                    api.APIOperation = "addCampaign"
	UPDATE_CAMPAIGN                 api.APIOperation = "updateCampaign"
	DELETE_CAMPAIGN                 api.APIOperation = "deleteCampaign"
	CLONE_CAMPAIGN                  api.APIOperation = "cloneCampaign"
	CHANGE_CAMPAIGN_OWNER           api.APIOperation = "changeOwner"
	CHANGE_CAMPAIGN_FAVORITE_STATUS api.APIOperation = "changeFavoriteStatus"

	// READ Operations.
	LIST_CAMPAIGNS     api.APIOperation = "listCampaigns"
	COUNT_CAMPAIGNS    api.APIOperation = "countCampaigns"
	GET_CAMPAIGN_BY_ID api.APIOperation = "getCampaignById"
)
