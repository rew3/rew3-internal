package brokerage/agent

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Agent
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_AGENT                   api.APIOperation = "financial_addAgent"
	UPDATE_AGENT                 api.APIOperation = "financial_updateAgent"
	DELETE_AGENT                 api.APIOperation = "financial_deleteAgent"
	CLONE_AGENT                  api.APIOperation = "financial_cloneAgent"
	CHANGE_AGENT_OWNER           api.APIOperation = "financial_changeAgentOwner"
	BULK_ADD_AGENT          api.APIOperation = "financial_bulkAddAgent"
	BULK_UPDATE_AGENT           api.APIOperation = "financial_bulkUpdateAgent"
	BULK_DELETE_AGENT           api.APIOperation = "financial_bulkDeleteAgent"
	BULK_CHANGE_AGENT_OWNER           api.APIOperation = "financial_bulkChangeAgentOwner"

	// READ APIs.
	LIST_AGENTS     api.APIOperation = "financial_listAgents"
	COUNT_AGENTS    api.APIOperation = "financial_countAgents"
	GET_AGENT_BY_ID api.APIOperation = "cfinancial_getAgentById"
    
)