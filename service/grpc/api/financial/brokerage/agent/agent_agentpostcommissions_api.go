package agent

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for AgentPostCommissions
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_AGENTPOSTCOMMISSIONS                   api.APIOperation = "financial_addAgentPostCommissions"
	UPDATE_AGENTPOSTCOMMISSIONS                 api.APIOperation = "financial_updateAgentPostCommissions"
	DELETE_AGENTPOSTCOMMISSIONS                 api.APIOperation = "financial_deleteAgentPostCommissions"
	CLONE_AGENTPOSTCOMMISSIONS                  api.APIOperation = "financial_cloneAgentPostCommissions"
	CHANGE_AGENTPOSTCOMMISSIONS_OWNER           api.APIOperation = "financial_changeAgentPostCommissionsOwner"
	BULK_ADD_AGENTPOSTCOMMISSIONS          api.APIOperation = "financial_bulkAddAgentPostCommissions"
	BULK_UPDATE_AGENTPOSTCOMMISSIONS           api.APIOperation = "financial_bulkUpdateAgentPostCommissions"
	BULK_DELETE_AGENTPOSTCOMMISSIONS           api.APIOperation = "financial_bulkDeleteAgentPostCommissions"
	BULK_CHANGE_AGENTPOSTCOMMISSIONS_OWNER           api.APIOperation = "financial_bulkChangeAgentPostCommissionsOwner"

	// READ APIs.
	LIST_AGENTPOSTCOMMISSIONS     api.APIOperation = "financial_listAgentPostCommissions"
	COUNT_AGENTPOSTCOMMISSIONS    api.APIOperation = "financial_countAgentPostCommissions"
	GET_AGENTPOSTCOMMISSIONS_BY_ID api.APIOperation = "cfinancial_getAgentPostCommissionsById"
    
)