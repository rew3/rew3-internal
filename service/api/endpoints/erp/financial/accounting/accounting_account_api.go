package accounting

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Account
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ACCOUNT               endpoints.Endpoint = "financial_addAccount"
	UPDATE_ACCOUNT            endpoints.Endpoint = "financial_updateAccount"
	DELETE_ACCOUNT            endpoints.Endpoint = "financial_deleteAccount"
	CLONE_ACCOUNT             endpoints.Endpoint = "financial_cloneAccount"
	CHANGE_ACCOUNT_OWNER      endpoints.Endpoint = "financial_changeAccountOwner"
	BULK_ADD_ACCOUNT          endpoints.Endpoint = "financial_bulkAddAccount"
	BULK_UPDATE_ACCOUNT       endpoints.Endpoint = "financial_bulkUpdateAccount"
	BULK_DELETE_ACCOUNT       endpoints.Endpoint = "financial_bulkDeleteAccount"
	BULK_CHANGE_ACCOUNT_OWNER endpoints.Endpoint = "financial_bulkChangeAccountOwner"

	// READ APIs.
	LIST_ACCOUNTS     endpoints.Endpoint = "financial_listAccounts"
	COUNT_ACCOUNTS    endpoints.Endpoint = "financial_countAccounts"
	GET_ACCOUNT_BY_ID endpoints.Endpoint = "cfinancial_getAccountById"
)
