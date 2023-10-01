package brokerage

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for TrustAccount
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_TRUSTACCOUNT               api.APIOperation = "financial_addTrustAccount"
	UPDATE_TRUSTACCOUNT            api.APIOperation = "financial_updateTrustAccount"
	DELETE_TRUSTACCOUNT            api.APIOperation = "financial_deleteTrustAccount"
	CLONE_TRUSTACCOUNT             api.APIOperation = "financial_cloneTrustAccount"
	CHANGE_TRUSTACCOUNT_OWNER      api.APIOperation = "financial_changeTrustAccountOwner"
	BULK_ADD_TRUSTACCOUNT          api.APIOperation = "financial_bulkAddTrustAccount"
	BULK_UPDATE_TRUSTACCOUNT       api.APIOperation = "financial_bulkUpdateTrustAccount"
	BULK_DELETE_TRUSTACCOUNT       api.APIOperation = "financial_bulkDeleteTrustAccount"
	BULK_CHANGE_TRUSTACCOUNT_OWNER api.APIOperation = "financial_bulkChangeTrustAccountOwner"

	// READ APIs.
	LIST_TRUSTACCOUNTS     api.APIOperation = "financial_listTrustAccounts"
	COUNT_TRUSTACCOUNTS    api.APIOperation = "financial_countTrustAccounts"
	GET_TRUSTACCOUNT_BY_ID api.APIOperation = "cfinancial_getTrustAccountById"
)
