package brokerage/agent

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for AgentPreCommissions
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_AGENTPRECOMMISSIONS                   api.APIOperation = "financial_addAgentPreCommissions"
	UPDATE_AGENTPRECOMMISSIONS                 api.APIOperation = "financial_updateAgentPreCommissions"
	DELETE_AGENTPRECOMMISSIONS                 api.APIOperation = "financial_deleteAgentPreCommissions"
	CLONE_AGENTPRECOMMISSIONS                  api.APIOperation = "financial_cloneAgentPreCommissions"
	CHANGE_AGENTPRECOMMISSIONS_OWNER           api.APIOperation = "financial_changeAgentPreCommissionsOwner"
	BULK_ADD_AGENTPRECOMMISSIONS          api.APIOperation = "financial_bulkAddAgentPreCommissions"
	BULK_UPDATE_AGENTPRECOMMISSIONS           api.APIOperation = "financial_bulkUpdateAgentPreCommissions"
	BULK_DELETE_AGENTPRECOMMISSIONS           api.APIOperation = "financial_bulkDeleteAgentPreCommissions"
	BULK_CHANGE_AGENTPRECOMMISSIONS_OWNER           api.APIOperation = "financial_bulkChangeAgentPreCommissionsOwner"

	// READ APIs.
	LIST_AGENTPRECOMMISSIONS     api.APIOperation = "financial_listAgentPreCommissions"
	COUNT_AGENTPRECOMMISSIONS    api.APIOperation = "financial_countAgentPreCommissions"
	GET_AGENTPRECOMMISSIONS_BY_ID api.APIOperation = "cfinancial_getAgentPreCommissionsById"
    
)