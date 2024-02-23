package agent

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for AgentClosingFee
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_AGENTCLOSINGFEE               endpoints.Endpoint = "financial_addAgentClosingFee"
	UPDATE_AGENTCLOSINGFEE            endpoints.Endpoint = "financial_updateAgentClosingFee"
	DELETE_AGENTCLOSINGFEE            endpoints.Endpoint = "financial_deleteAgentClosingFee"
	CLONE_AGENTCLOSINGFEE             endpoints.Endpoint = "financial_cloneAgentClosingFee"
	CHANGE_AGENTCLOSINGFEE_OWNER      endpoints.Endpoint = "financial_changeAgentClosingFeeOwner"
	BULK_ADD_AGENTCLOSINGFEE          endpoints.Endpoint = "financial_bulkAddAgentClosingFee"
	BULK_UPDATE_AGENTCLOSINGFEE       endpoints.Endpoint = "financial_bulkUpdateAgentClosingFee"
	BULK_DELETE_AGENTCLOSINGFEE       endpoints.Endpoint = "financial_bulkDeleteAgentClosingFee"
	BULK_CHANGE_AGENTCLOSINGFEE_OWNER endpoints.Endpoint = "financial_bulkChangeAgentClosingFeeOwner"

	// READ APIs.
	LIST_AGENTCLOSINGFEES     endpoints.Endpoint = "financial_listAgentClosingFees"
	COUNT_AGENTCLOSINGFEES    endpoints.Endpoint = "financial_countAgentClosingFees"
	GET_AGENTCLOSINGFEE_BY_ID endpoints.Endpoint = "cfinancial_getAgentClosingFeeById"
)
