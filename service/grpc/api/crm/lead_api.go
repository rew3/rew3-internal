package crm

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_LEAD                    api.APIOperation = "crm_addLead"
	UPDATE_LEAD                 api.APIOperation = "crm_pdateLead"
	DELETE_LEAD                 api.APIOperation = "crm_deleteLead"
	CLONE_LEAD                  api.APIOperation = "crm_cloneLead"
	CHANGE_LEAD_OWNER           api.APIOperation = "crm_changeLeadOwner"
	CHANGE_LEAD_FAVORITE_STATUS api.APIOperation = "crm_changeLeadFavoriteStatus"

	// READ Operations.
	LIST_LEADS     api.APIOperation = "crm_listLeads"
	COUNT_LEADS    api.APIOperation = "crm_countLeads"
	GET_LEAD_BY_ID api.APIOperation = "crm_getLeadsById"
)
