package recurringinvoice

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for RecurringInvoice
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_RECURRINGINVOICE                   api.APIOperation = "financial_addRecurringInvoice"
	UPDATE_RECURRINGINVOICE                 api.APIOperation = "financial_updateRecurringInvoice"
	DELETE_RECURRINGINVOICE                 api.APIOperation = "financial_deleteRecurringInvoice"
	CLONE_RECURRINGINVOICE                  api.APIOperation = "financial_cloneRecurringInvoice"
	CHANGE_RECURRINGINVOICE_OWNER           api.APIOperation = "financial_changeRecurringInvoiceOwner"
	BULK_ADD_RECURRINGINVOICE          api.APIOperation = "financial_bulkAddRecurringInvoice"
	BULK_UPDATE_RECURRINGINVOICE           api.APIOperation = "financial_bulkUpdateRecurringInvoice"
	BULK_DELETE_RECURRINGINVOICE           api.APIOperation = "financial_bulkDeleteRecurringInvoice"
	BULK_CHANGE_RECURRINGINVOICE_OWNER           api.APIOperation = "financial_bulkChangeRecurringInvoiceOwner"

	// READ APIs.
	LIST_RECURRINGINVOICES     api.APIOperation = "financial_listRecurringInvoices"
	COUNT_RECURRINGINVOICES    api.APIOperation = "financial_countRecurringInvoices"
	GET_RECURRINGINVOICE_BY_ID api.APIOperation = "cfinancial_getRecurringInvoiceById"
    
)