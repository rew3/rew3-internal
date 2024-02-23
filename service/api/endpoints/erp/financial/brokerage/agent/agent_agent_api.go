package agent

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Agent
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_AGENT               endpoints.Endpoint = "financial_addAgent"
	UPDATE_AGENT            endpoints.Endpoint = "financial_updateAgent"
	DELETE_AGENT            endpoints.Endpoint = "financial_deleteAgent"
	CLONE_AGENT             endpoints.Endpoint = "financial_cloneAgent"
	CHANGE_AGENT_OWNER      endpoints.Endpoint = "financial_changeAgentOwner"
	BULK_ADD_AGENT          endpoints.Endpoint = "financial_bulkAddAgent"
	BULK_UPDATE_AGENT       endpoints.Endpoint = "financial_bulkUpdateAgent"
	BULK_DELETE_AGENT       endpoints.Endpoint = "financial_bulkDeleteAgent"
	BULK_CHANGE_AGENT_OWNER endpoints.Endpoint = "financial_bulkChangeAgentOwner"

	// READ APIs.
	LIST_AGENTS     endpoints.Endpoint = "financial_listAgents"
	COUNT_AGENTS    endpoints.Endpoint = "financial_countAgents"
	GET_AGENT_BY_ID endpoints.Endpoint = "cfinancial_getAgentById"
)
