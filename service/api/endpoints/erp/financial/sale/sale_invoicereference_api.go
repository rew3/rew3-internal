package sale

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for InvoiceReference
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_INVOICEREFERENCE               endpoints.Endpoint = "financial_addInvoiceReference"
	UPDATE_INVOICEREFERENCE            endpoints.Endpoint = "financial_updateInvoiceReference"
	DELETE_INVOICEREFERENCE            endpoints.Endpoint = "financial_deleteInvoiceReference"
	CLONE_INVOICEREFERENCE             endpoints.Endpoint = "financial_cloneInvoiceReference"
	CHANGE_INVOICEREFERENCE_OWNER      endpoints.Endpoint = "financial_changeInvoiceReferenceOwner"
	BULK_ADD_INVOICEREFERENCE          endpoints.Endpoint = "financial_bulkAddInvoiceReference"
	BULK_UPDATE_INVOICEREFERENCE       endpoints.Endpoint = "financial_bulkUpdateInvoiceReference"
	BULK_DELETE_INVOICEREFERENCE       endpoints.Endpoint = "financial_bulkDeleteInvoiceReference"
	BULK_CHANGE_INVOICEREFERENCE_OWNER endpoints.Endpoint = "financial_bulkChangeInvoiceReferenceOwner"

	// READ APIs.
	LIST_INVOICEREFERENCES     endpoints.Endpoint = "financial_listInvoiceReferences"
	COUNT_INVOICEREFERENCES    endpoints.Endpoint = "financial_countInvoiceReferences"
	GET_INVOICEREFERENCE_BY_ID endpoints.Endpoint = "cfinancial_getInvoiceReferenceById"
)
