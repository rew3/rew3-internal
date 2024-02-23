package crm

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_CAMPAIGN          endpoints.Endpoint = "crm_addCampaign"
	UPDATE_CAMPAIGN       endpoints.Endpoint = "crm_updateCampaign"
	DELETE_CAMPAIGN       endpoints.Endpoint = "crm_deleteCampaign"
	CLONE_CAMPAIGN        endpoints.Endpoint = "crm_cloneCampaign"
	CHANGE_CAMPAIGN_OWNER endpoints.Endpoint = "crm_changeCampaignOwner"
	SET_CAMPAIGN_FAVORITE endpoints.Endpoint = "crm_setCampaignFavorite"

	// READ Operations.
	LIST_CAMPAIGNS     endpoints.Endpoint = "crm_listCampaigns"
	COUNT_CAMPAIGNS    endpoints.Endpoint = "crm_countCampaigns"
	GET_CAMPAIGN_BY_ID endpoints.Endpoint = "crm_getCampaignById"
)
