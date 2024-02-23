package agent

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for AgentPreCommissions
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_AGENTPRECOMMISSIONS               endpoints.Endpoint = "financial_addAgentPreCommissions"
	UPDATE_AGENTPRECOMMISSIONS            endpoints.Endpoint = "financial_updateAgentPreCommissions"
	DELETE_AGENTPRECOMMISSIONS            endpoints.Endpoint = "financial_deleteAgentPreCommissions"
	CLONE_AGENTPRECOMMISSIONS             endpoints.Endpoint = "financial_cloneAgentPreCommissions"
	CHANGE_AGENTPRECOMMISSIONS_OWNER      endpoints.Endpoint = "financial_changeAgentPreCommissionsOwner"
	BULK_ADD_AGENTPRECOMMISSIONS          endpoints.Endpoint = "financial_bulkAddAgentPreCommissions"
	BULK_UPDATE_AGENTPRECOMMISSIONS       endpoints.Endpoint = "financial_bulkUpdateAgentPreCommissions"
	BULK_DELETE_AGENTPRECOMMISSIONS       endpoints.Endpoint = "financial_bulkDeleteAgentPreCommissions"
	BULK_CHANGE_AGENTPRECOMMISSIONS_OWNER endpoints.Endpoint = "financial_bulkChangeAgentPreCommissionsOwner"

	// READ APIs.
	LIST_AGENTPRECOMMISSIONS      endpoints.Endpoint = "financial_listAgentPreCommissions"
	COUNT_AGENTPRECOMMISSIONS     endpoints.Endpoint = "financial_countAgentPreCommissions"
	GET_AGENTPRECOMMISSIONS_BY_ID endpoints.Endpoint = "cfinancial_getAgentPreCommissionsById"
)
