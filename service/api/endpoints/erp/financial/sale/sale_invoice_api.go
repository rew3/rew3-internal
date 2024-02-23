package sale

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Invoice
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_INVOICE               endpoints.Endpoint = "financial_addInvoice"
	UPDATE_INVOICE            endpoints.Endpoint = "financial_updateInvoice"
	DELETE_INVOICE            endpoints.Endpoint = "financial_deleteInvoice"
	CLONE_INVOICE             endpoints.Endpoint = "financial_cloneInvoice"
	CHANGE_INVOICE_OWNER      endpoints.Endpoint = "financial_changeInvoiceOwner"
	BULK_ADD_INVOICE          endpoints.Endpoint = "financial_bulkAddInvoice"
	BULK_UPDATE_INVOICE       endpoints.Endpoint = "financial_bulkUpdateInvoice"
	BULK_DELETE_INVOICE       endpoints.Endpoint = "financial_bulkDeleteInvoice"
	BULK_CHANGE_INVOICE_OWNER endpoints.Endpoint = "financial_bulkChangeInvoiceOwner"

	// READ APIs.
	LIST_INVOICES     endpoints.Endpoint = "financial_listInvoices"
	COUNT_INVOICES    endpoints.Endpoint = "financial_countInvoices"
	GET_INVOICE_BY_ID endpoints.Endpoint = "cfinancial_getInvoiceById"
)
