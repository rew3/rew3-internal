package payment

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for InvoicePayment
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_INVOICEPAYMENT                   api.APIOperation = "financial_addInvoicePayment"
	UPDATE_INVOICEPAYMENT                 api.APIOperation = "financial_updateInvoicePayment"
	DELETE_INVOICEPAYMENT                 api.APIOperation = "financial_deleteInvoicePayment"
	CLONE_INVOICEPAYMENT                  api.APIOperation = "financial_cloneInvoicePayment"
	CHANGE_INVOICEPAYMENT_OWNER           api.APIOperation = "financial_changeInvoicePaymentOwner"
	BULK_ADD_INVOICEPAYMENT          api.APIOperation = "financial_bulkAddInvoicePayment"
	BULK_UPDATE_INVOICEPAYMENT           api.APIOperation = "financial_bulkUpdateInvoicePayment"
	BULK_DELETE_INVOICEPAYMENT           api.APIOperation = "financial_bulkDeleteInvoicePayment"
	BULK_CHANGE_INVOICEPAYMENT_OWNER           api.APIOperation = "financial_bulkChangeInvoicePaymentOwner"

	// READ APIs.
	LIST_INVOICEPAYMENTS     api.APIOperation = "financial_listInvoicePayments"
	COUNT_INVOICEPAYMENTS    api.APIOperation = "financial_countInvoicePayments"
	GET_INVOICEPAYMENT_BY_ID api.APIOperation = "cfinancial_getInvoicePaymentById"
    
)