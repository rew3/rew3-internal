package crm

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_LEAD                    api.APIOperation = "addLead"
	UPDATE_LEAD                 api.APIOperation = "updateLead"
	DELETE_LEAD                 api.APIOperation = "deleteLead"
	CLONE_LEAD                  api.APIOperation = "cloneLead"
	CHANGE_LEAD_OWNER           api.APIOperation = "changeOwner"
	CHANGE_LEAD_FAVORITE_STATUS api.APIOperation = "changeFavoriteStatus"

	// READ Operations.
	LIST_LEADS     api.APIOperation = "listLeads"
	COUNT_LEADS    api.APIOperation = "countLeads"
	GET_LEAD_BY_ID api.APIOperation = "getLeadsById"
)
