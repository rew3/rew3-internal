package payment

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for InvoicePayment
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_INVOICEPAYMENT               endpoints.Endpoint = "financial_addInvoicePayment"
	UPDATE_INVOICEPAYMENT            endpoints.Endpoint = "financial_updateInvoicePayment"
	DELETE_INVOICEPAYMENT            endpoints.Endpoint = "financial_deleteInvoicePayment"
	CLONE_INVOICEPAYMENT             endpoints.Endpoint = "financial_cloneInvoicePayment"
	CHANGE_INVOICEPAYMENT_OWNER      endpoints.Endpoint = "financial_changeInvoicePaymentOwner"
	BULK_ADD_INVOICEPAYMENT          endpoints.Endpoint = "financial_bulkAddInvoicePayment"
	BULK_UPDATE_INVOICEPAYMENT       endpoints.Endpoint = "financial_bulkUpdateInvoicePayment"
	BULK_DELETE_INVOICEPAYMENT       endpoints.Endpoint = "financial_bulkDeleteInvoicePayment"
	BULK_CHANGE_INVOICEPAYMENT_OWNER endpoints.Endpoint = "financial_bulkChangeInvoicePaymentOwner"

	// READ APIs.
	LIST_INVOICEPAYMENTS     endpoints.Endpoint = "financial_listInvoicePayments"
	COUNT_INVOICEPAYMENTS    endpoints.Endpoint = "financial_countInvoicePayments"
	GET_INVOICEPAYMENT_BY_ID endpoints.Endpoint = "cfinancial_getInvoicePaymentById"
)
