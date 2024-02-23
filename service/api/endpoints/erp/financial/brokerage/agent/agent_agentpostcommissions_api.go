package agent

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for AgentPostCommissions
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_AGENTPOSTCOMMISSIONS               endpoints.Endpoint = "financial_addAgentPostCommissions"
	UPDATE_AGENTPOSTCOMMISSIONS            endpoints.Endpoint = "financial_updateAgentPostCommissions"
	DELETE_AGENTPOSTCOMMISSIONS            endpoints.Endpoint = "financial_deleteAgentPostCommissions"
	CLONE_AGENTPOSTCOMMISSIONS             endpoints.Endpoint = "financial_cloneAgentPostCommissions"
	CHANGE_AGENTPOSTCOMMISSIONS_OWNER      endpoints.Endpoint = "financial_changeAgentPostCommissionsOwner"
	BULK_ADD_AGENTPOSTCOMMISSIONS          endpoints.Endpoint = "financial_bulkAddAgentPostCommissions"
	BULK_UPDATE_AGENTPOSTCOMMISSIONS       endpoints.Endpoint = "financial_bulkUpdateAgentPostCommissions"
	BULK_DELETE_AGENTPOSTCOMMISSIONS       endpoints.Endpoint = "financial_bulkDeleteAgentPostCommissions"
	BULK_CHANGE_AGENTPOSTCOMMISSIONS_OWNER endpoints.Endpoint = "financial_bulkChangeAgentPostCommissionsOwner"

	// READ APIs.
	LIST_AGENTPOSTCOMMISSIONS      endpoints.Endpoint = "financial_listAgentPostCommissions"
	COUNT_AGENTPOSTCOMMISSIONS     endpoints.Endpoint = "financial_countAgentPostCommissions"
	GET_AGENTPOSTCOMMISSIONS_BY_ID endpoints.Endpoint = "cfinancial_getAgentPostCommissionsById"
)
