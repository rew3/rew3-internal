package payment

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for RecurringInvoicePayment
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_RECURRINGINVOICEPAYMENT               endpoints.Endpoint = "financial_addRecurringInvoicePayment"
	UPDATE_RECURRINGINVOICEPAYMENT            endpoints.Endpoint = "financial_updateRecurringInvoicePayment"
	DELETE_RECURRINGINVOICEPAYMENT            endpoints.Endpoint = "financial_deleteRecurringInvoicePayment"
	CLONE_RECURRINGINVOICEPAYMENT             endpoints.Endpoint = "financial_cloneRecurringInvoicePayment"
	CHANGE_RECURRINGINVOICEPAYMENT_OWNER      endpoints.Endpoint = "financial_changeRecurringInvoicePaymentOwner"
	BULK_ADD_RECURRINGINVOICEPAYMENT          endpoints.Endpoint = "financial_bulkAddRecurringInvoicePayment"
	BULK_UPDATE_RECURRINGINVOICEPAYMENT       endpoints.Endpoint = "financial_bulkUpdateRecurringInvoicePayment"
	BULK_DELETE_RECURRINGINVOICEPAYMENT       endpoints.Endpoint = "financial_bulkDeleteRecurringInvoicePayment"
	BULK_CHANGE_RECURRINGINVOICEPAYMENT_OWNER endpoints.Endpoint = "financial_bulkChangeRecurringInvoicePaymentOwner"

	// READ APIs.
	LIST_RECURRINGINVOICEPAYMENTS     endpoints.Endpoint = "financial_listRecurringInvoicePayments"
	COUNT_RECURRINGINVOICEPAYMENTS    endpoints.Endpoint = "financial_countRecurringInvoicePayments"
	GET_RECURRINGINVOICEPAYMENT_BY_ID endpoints.Endpoint = "cfinancial_getRecurringInvoicePaymentById"
)
