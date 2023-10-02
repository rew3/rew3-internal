package sale

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for InvoiceReference
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_INVOICEREFERENCE                   api.APIOperation = "financial_addInvoiceReference"
	UPDATE_INVOICEREFERENCE                 api.APIOperation = "financial_updateInvoiceReference"
	DELETE_INVOICEREFERENCE                 api.APIOperation = "financial_deleteInvoiceReference"
	CLONE_INVOICEREFERENCE                  api.APIOperation = "financial_cloneInvoiceReference"
	CHANGE_INVOICEREFERENCE_OWNER           api.APIOperation = "financial_changeInvoiceReferenceOwner"
	BULK_ADD_INVOICEREFERENCE          api.APIOperation = "financial_bulkAddInvoiceReference"
	BULK_UPDATE_INVOICEREFERENCE           api.APIOperation = "financial_bulkUpdateInvoiceReference"
	BULK_DELETE_INVOICEREFERENCE           api.APIOperation = "financial_bulkDeleteInvoiceReference"
	BULK_CHANGE_INVOICEREFERENCE_OWNER           api.APIOperation = "financial_bulkChangeInvoiceReferenceOwner"

	// READ APIs.
	LIST_INVOICEREFERENCES     api.APIOperation = "financial_listInvoiceReferences"
	COUNT_INVOICEREFERENCES    api.APIOperation = "financial_countInvoiceReferences"
	GET_INVOICEREFERENCE_BY_ID api.APIOperation = "cfinancial_getInvoiceReferenceById"
    
)