package payment

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for RecurringInvoicePayment
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_RECURRINGINVOICEPAYMENT               api.APIOperation = "financial_addRecurringInvoicePayment"
	UPDATE_RECURRINGINVOICEPAYMENT            api.APIOperation = "financial_updateRecurringInvoicePayment"
	DELETE_RECURRINGINVOICEPAYMENT            api.APIOperation = "financial_deleteRecurringInvoicePayment"
	CLONE_RECURRINGINVOICEPAYMENT             api.APIOperation = "financial_cloneRecurringInvoicePayment"
	CHANGE_RECURRINGINVOICEPAYMENT_OWNER      api.APIOperation = "financial_changeRecurringInvoicePaymentOwner"
	BULK_ADD_RECURRINGINVOICEPAYMENT          api.APIOperation = "financial_bulkAddRecurringInvoicePayment"
	BULK_UPDATE_RECURRINGINVOICEPAYMENT       api.APIOperation = "financial_bulkUpdateRecurringInvoicePayment"
	BULK_DELETE_RECURRINGINVOICEPAYMENT       api.APIOperation = "financial_bulkDeleteRecurringInvoicePayment"
	BULK_CHANGE_RECURRINGINVOICEPAYMENT_OWNER api.APIOperation = "financial_bulkChangeRecurringInvoicePaymentOwner"

	// READ APIs.
	LIST_RECURRINGINVOICEPAYMENTS     api.APIOperation = "financial_listRecurringInvoicePayments"
	COUNT_RECURRINGINVOICEPAYMENTS    api.APIOperation = "financial_countRecurringInvoicePayments"
	GET_RECURRINGINVOICEPAYMENT_BY_ID api.APIOperation = "cfinancial_getRecurringInvoicePaymentById"
)
