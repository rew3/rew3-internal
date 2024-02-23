package recurringinvoice

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for RecurringInvoice
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_RECURRINGINVOICE               endpoints.Endpoint = "financial_addRecurringInvoice"
	UPDATE_RECURRINGINVOICE            endpoints.Endpoint = "financial_updateRecurringInvoice"
	DELETE_RECURRINGINVOICE            endpoints.Endpoint = "financial_deleteRecurringInvoice"
	CLONE_RECURRINGINVOICE             endpoints.Endpoint = "financial_cloneRecurringInvoice"
	CHANGE_RECURRINGINVOICE_OWNER      endpoints.Endpoint = "financial_changeRecurringInvoiceOwner"
	BULK_ADD_RECURRINGINVOICE          endpoints.Endpoint = "financial_bulkAddRecurringInvoice"
	BULK_UPDATE_RECURRINGINVOICE       endpoints.Endpoint = "financial_bulkUpdateRecurringInvoice"
	BULK_DELETE_RECURRINGINVOICE       endpoints.Endpoint = "financial_bulkDeleteRecurringInvoice"
	BULK_CHANGE_RECURRINGINVOICE_OWNER endpoints.Endpoint = "financial_bulkChangeRecurringInvoiceOwner"

	// READ APIs.
	LIST_RECURRINGINVOICES     endpoints.Endpoint = "financial_listRecurringInvoices"
	COUNT_RECURRINGINVOICES    endpoints.Endpoint = "financial_countRecurringInvoices"
	GET_RECURRINGINVOICE_BY_ID endpoints.Endpoint = "cfinancial_getRecurringInvoiceById"
)
