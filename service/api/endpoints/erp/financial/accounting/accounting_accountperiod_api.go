package accounting

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for AccountPeriod
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ACCOUNTPERIOD               endpoints.Endpoint = "financial_addAccountPeriod"
	UPDATE_ACCOUNTPERIOD            endpoints.Endpoint = "financial_updateAccountPeriod"
	DELETE_ACCOUNTPERIOD            endpoints.Endpoint = "financial_deleteAccountPeriod"
	CLONE_ACCOUNTPERIOD             endpoints.Endpoint = "financial_cloneAccountPeriod"
	CHANGE_ACCOUNTPERIOD_OWNER      endpoints.Endpoint = "financial_changeAccountPeriodOwner"
	BULK_ADD_ACCOUNTPERIOD          endpoints.Endpoint = "financial_bulkAddAccountPeriod"
	BULK_UPDATE_ACCOUNTPERIOD       endpoints.Endpoint = "financial_bulkUpdateAccountPeriod"
	BULK_DELETE_ACCOUNTPERIOD       endpoints.Endpoint = "financial_bulkDeleteAccountPeriod"
	BULK_CHANGE_ACCOUNTPERIOD_OWNER endpoints.Endpoint = "financial_bulkChangeAccountPeriodOwner"

	// READ APIs.
	LIST_ACCOUNTPERIODS     endpoints.Endpoint = "financial_listAccountPeriods"
	COUNT_ACCOUNTPERIODS    endpoints.Endpoint = "financial_countAccountPeriods"
	GET_ACCOUNTPERIOD_BY_ID endpoints.Endpoint = "cfinancial_getAccountPeriodById"
)
