package accounting

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Account
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ACCOUNT               api.APIOperation = "financial_addAccount"
	UPDATE_ACCOUNT            api.APIOperation = "financial_updateAccount"
	DELETE_ACCOUNT            api.APIOperation = "financial_deleteAccount"
	CLONE_ACCOUNT             api.APIOperation = "financial_cloneAccount"
	CHANGE_ACCOUNT_OWNER      api.APIOperation = "financial_changeAccountOwner"
	BULK_ADD_ACCOUNT          api.APIOperation = "financial_bulkAddAccount"
	BULK_UPDATE_ACCOUNT       api.APIOperation = "financial_bulkUpdateAccount"
	BULK_DELETE_ACCOUNT       api.APIOperation = "financial_bulkDeleteAccount"
	BULK_CHANGE_ACCOUNT_OWNER api.APIOperation = "financial_bulkChangeAccountOwner"

	// READ APIs.
	LIST_ACCOUNTS     api.APIOperation = "financial_listAccounts"
	COUNT_ACCOUNTS    api.APIOperation = "financial_countAccounts"
	GET_ACCOUNT_BY_ID api.APIOperation = "cfinancial_getAccountById"
)
