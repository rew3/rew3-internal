package payment

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for PaymentTerm
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_PAYMENTTERM               endpoints.Endpoint = "financial_addPaymentTerm"
	UPDATE_PAYMENTTERM            endpoints.Endpoint = "financial_updatePaymentTerm"
	DELETE_PAYMENTTERM            endpoints.Endpoint = "financial_deletePaymentTerm"
	CLONE_PAYMENTTERM             endpoints.Endpoint = "financial_clonePaymentTerm"
	CHANGE_PAYMENTTERM_OWNER      endpoints.Endpoint = "financial_changePaymentTermOwner"
	BULK_ADD_PAYMENTTERM          endpoints.Endpoint = "financial_bulkAddPaymentTerm"
	BULK_UPDATE_PAYMENTTERM       endpoints.Endpoint = "financial_bulkUpdatePaymentTerm"
	BULK_DELETE_PAYMENTTERM       endpoints.Endpoint = "financial_bulkDeletePaymentTerm"
	BULK_CHANGE_PAYMENTTERM_OWNER endpoints.Endpoint = "financial_bulkChangePaymentTermOwner"

	// READ APIs.
	LIST_PAYMENTTERMS     endpoints.Endpoint = "financial_listPaymentTerms"
	COUNT_PAYMENTTERMS    endpoints.Endpoint = "financial_countPaymentTerms"
	GET_PAYMENTTERM_BY_ID endpoints.Endpoint = "cfinancial_getPaymentTermById"
)
