package crm

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_LEAD          endpoints.Endpoint = "crm_addLead"
	UPDATE_LEAD       endpoints.Endpoint = "crm_pdateLead"
	DELETE_LEAD       endpoints.Endpoint = "crm_deleteLead"
	CLONE_LEAD        endpoints.Endpoint = "crm_cloneLead"
	CHANGE_LEAD_OWNER endpoints.Endpoint = "crm_changeLeadOwner"
	SET_LEAD_FAVORITE endpoints.Endpoint = "crm_setLeadFavorite"

	// READ Operations.
	LIST_LEADS     endpoints.Endpoint = "crm_listLeads"
	COUNT_LEADS    endpoints.Endpoint = "crm_countLeads"
	GET_LEAD_BY_ID endpoints.Endpoint = "crm_getLeadsById"
)
