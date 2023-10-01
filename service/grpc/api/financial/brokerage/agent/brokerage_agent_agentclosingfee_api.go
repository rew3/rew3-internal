package brokerage/agent

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for AgentClosingFee
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_AGENTCLOSINGFEE                   api.APIOperation = "financial_addAgentClosingFee"
	UPDATE_AGENTCLOSINGFEE                 api.APIOperation = "financial_updateAgentClosingFee"
	DELETE_AGENTCLOSINGFEE                 api.APIOperation = "financial_deleteAgentClosingFee"
	CLONE_AGENTCLOSINGFEE                  api.APIOperation = "financial_cloneAgentClosingFee"
	CHANGE_AGENTCLOSINGFEE_OWNER           api.APIOperation = "financial_changeAgentClosingFeeOwner"
	BULK_ADD_AGENTCLOSINGFEE          api.APIOperation = "financial_bulkAddAgentClosingFee"
	BULK_UPDATE_AGENTCLOSINGFEE           api.APIOperation = "financial_bulkUpdateAgentClosingFee"
	BULK_DELETE_AGENTCLOSINGFEE           api.APIOperation = "financial_bulkDeleteAgentClosingFee"
	BULK_CHANGE_AGENTCLOSINGFEE_OWNER           api.APIOperation = "financial_bulkChangeAgentClosingFeeOwner"

	// READ APIs.
	LIST_AGENTCLOSINGFEES     api.APIOperation = "financial_listAgentClosingFees"
	COUNT_AGENTCLOSINGFEES    api.APIOperation = "financial_countAgentClosingFees"
	GET_AGENTCLOSINGFEE_BY_ID api.APIOperation = "cfinancial_getAgentClosingFeeById"
    
)