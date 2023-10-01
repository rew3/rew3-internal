package accounting

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for AccountPeriod
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ACCOUNTPERIOD               api.APIOperation = "financial_addAccountPeriod"
	UPDATE_ACCOUNTPERIOD            api.APIOperation = "financial_updateAccountPeriod"
	DELETE_ACCOUNTPERIOD            api.APIOperation = "financial_deleteAccountPeriod"
	CLONE_ACCOUNTPERIOD             api.APIOperation = "financial_cloneAccountPeriod"
	CHANGE_ACCOUNTPERIOD_OWNER      api.APIOperation = "financial_changeAccountPeriodOwner"
	BULK_ADD_ACCOUNTPERIOD          api.APIOperation = "financial_bulkAddAccountPeriod"
	BULK_UPDATE_ACCOUNTPERIOD       api.APIOperation = "financial_bulkUpdateAccountPeriod"
	BULK_DELETE_ACCOUNTPERIOD       api.APIOperation = "financial_bulkDeleteAccountPeriod"
	BULK_CHANGE_ACCOUNTPERIOD_OWNER api.APIOperation = "financial_bulkChangeAccountPeriodOwner"

	// READ APIs.
	LIST_ACCOUNTPERIODS     api.APIOperation = "financial_listAccountPeriods"
	COUNT_ACCOUNTPERIODS    api.APIOperation = "financial_countAccountPeriods"
	GET_ACCOUNTPERIOD_BY_ID api.APIOperation = "cfinancial_getAccountPeriodById"
)
