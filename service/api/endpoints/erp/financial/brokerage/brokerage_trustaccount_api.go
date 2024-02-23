package brokerage

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for TrustAccount
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_TRUSTACCOUNT               endpoints.Endpoint = "financial_addTrustAccount"
	UPDATE_TRUSTACCOUNT            endpoints.Endpoint = "financial_updateTrustAccount"
	DELETE_TRUSTACCOUNT            endpoints.Endpoint = "financial_deleteTrustAccount"
	CLONE_TRUSTACCOUNT             endpoints.Endpoint = "financial_cloneTrustAccount"
	CHANGE_TRUSTACCOUNT_OWNER      endpoints.Endpoint = "financial_changeTrustAccountOwner"
	BULK_ADD_TRUSTACCOUNT          endpoints.Endpoint = "financial_bulkAddTrustAccount"
	BULK_UPDATE_TRUSTACCOUNT       endpoints.Endpoint = "financial_bulkUpdateTrustAccount"
	BULK_DELETE_TRUSTACCOUNT       endpoints.Endpoint = "financial_bulkDeleteTrustAccount"
	BULK_CHANGE_TRUSTACCOUNT_OWNER endpoints.Endpoint = "financial_bulkChangeTrustAccountOwner"

	// READ APIs.
	LIST_TRUSTACCOUNTS     endpoints.Endpoint = "financial_listTrustAccounts"
	COUNT_TRUSTACCOUNTS    endpoints.Endpoint = "financial_countTrustAccounts"
	GET_TRUSTACCOUNT_BY_ID endpoints.Endpoint = "cfinancial_getTrustAccountById"
)
