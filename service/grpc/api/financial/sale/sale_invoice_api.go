package sale

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Invoice
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_INVOICE                   api.APIOperation = "financial_addInvoice"
	UPDATE_INVOICE                 api.APIOperation = "financial_updateInvoice"
	DELETE_INVOICE                 api.APIOperation = "financial_deleteInvoice"
	CLONE_INVOICE                  api.APIOperation = "financial_cloneInvoice"
	CHANGE_INVOICE_OWNER           api.APIOperation = "financial_changeInvoiceOwner"
	BULK_ADD_INVOICE          api.APIOperation = "financial_bulkAddInvoice"
	BULK_UPDATE_INVOICE           api.APIOperation = "financial_bulkUpdateInvoice"
	BULK_DELETE_INVOICE           api.APIOperation = "financial_bulkDeleteInvoice"
	BULK_CHANGE_INVOICE_OWNER           api.APIOperation = "financial_bulkChangeInvoiceOwner"

	// READ APIs.
	LIST_INVOICES     api.APIOperation = "financial_listInvoices"
	COUNT_INVOICES    api.APIOperation = "financial_countInvoices"
	GET_INVOICE_BY_ID api.APIOperation = "cfinancial_getInvoiceById"
    
)