package brokerage

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for AgentPlan
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_AGENTPLAN               endpoints.Endpoint = "financial_addAgentPlan"
	UPDATE_AGENTPLAN            endpoints.Endpoint = "financial_updateAgentPlan"
	DELETE_AGENTPLAN            endpoints.Endpoint = "financial_deleteAgentPlan"
	CLONE_AGENTPLAN             endpoints.Endpoint = "financial_cloneAgentPlan"
	CHANGE_AGENTPLAN_OWNER      endpoints.Endpoint = "financial_changeAgentPlanOwner"
	BULK_ADD_AGENTPLAN          endpoints.Endpoint = "financial_bulkAddAgentPlan"
	BULK_UPDATE_AGENTPLAN       endpoints.Endpoint = "financial_bulkUpdateAgentPlan"
	BULK_DELETE_AGENTPLAN       endpoints.Endpoint = "financial_bulkDeleteAgentPlan"
	BULK_CHANGE_AGENTPLAN_OWNER endpoints.Endpoint = "financial_bulkChangeAgentPlanOwner"

	// READ APIs.
	LIST_AGENTPLANS     endpoints.Endpoint = "financial_listAgentPlans"
	COUNT_AGENTPLANS    endpoints.Endpoint = "financial_countAgentPlans"
	GET_AGENTPLAN_BY_ID endpoints.Endpoint = "cfinancial_getAgentPlanById"
)
