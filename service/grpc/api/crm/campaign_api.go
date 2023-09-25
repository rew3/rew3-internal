package crm

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_CAMPAIGN                    api.APIOperation = "crm_addCampaign"
	UPDATE_CAMPAIGN                 api.APIOperation = "crm_updateCampaign"
	DELETE_CAMPAIGN                 api.APIOperation = "crm_deleteCampaign"
	CLONE_CAMPAIGN                  api.APIOperation = "crm_cloneCampaign"
	CHANGE_CAMPAIGN_OWNER           api.APIOperation = "crm_changeCampaignOwner"
	CHANGE_CAMPAIGN_FAVORITE_STATUS api.APIOperation = "crm_changeCampaignFavoriteStatus"

	// READ Operations.
	LIST_CAMPAIGNS     api.APIOperation = "crm_listCampaigns"
	COUNT_CAMPAIGNS    api.APIOperation = "crm_countCampaigns"
	GET_CAMPAIGN_BY_ID api.APIOperation = "crm_getCampaignById"
)
