package brokerage

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for AgentPlan
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_AGENTPLAN               api.APIOperation = "financial_addAgentPlan"
	UPDATE_AGENTPLAN            api.APIOperation = "financial_updateAgentPlan"
	DELETE_AGENTPLAN            api.APIOperation = "financial_deleteAgentPlan"
	CLONE_AGENTPLAN             api.APIOperation = "financial_cloneAgentPlan"
	CHANGE_AGENTPLAN_OWNER      api.APIOperation = "financial_changeAgentPlanOwner"
	BULK_ADD_AGENTPLAN          api.APIOperation = "financial_bulkAddAgentPlan"
	BULK_UPDATE_AGENTPLAN       api.APIOperation = "financial_bulkUpdateAgentPlan"
	BULK_DELETE_AGENTPLAN       api.APIOperation = "financial_bulkDeleteAgentPlan"
	BULK_CHANGE_AGENTPLAN_OWNER api.APIOperation = "financial_bulkChangeAgentPlanOwner"

	// READ APIs.
	LIST_AGENTPLANS     api.APIOperation = "financial_listAgentPlans"
	COUNT_AGENTPLANS    api.APIOperation = "financial_countAgentPlans"
	GET_AGENTPLAN_BY_ID api.APIOperation = "cfinancial_getAgentPlanById"
)
