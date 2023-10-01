package payment

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for PaymentTerm
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_PAYMENTTERM               api.APIOperation = "financial_addPaymentTerm"
	UPDATE_PAYMENTTERM            api.APIOperation = "financial_updatePaymentTerm"
	DELETE_PAYMENTTERM            api.APIOperation = "financial_deletePaymentTerm"
	CLONE_PAYMENTTERM             api.APIOperation = "financial_clonePaymentTerm"
	CHANGE_PAYMENTTERM_OWNER      api.APIOperation = "financial_changePaymentTermOwner"
	BULK_ADD_PAYMENTTERM          api.APIOperation = "financial_bulkAddPaymentTerm"
	BULK_UPDATE_PAYMENTTERM       api.APIOperation = "financial_bulkUpdatePaymentTerm"
	BULK_DELETE_PAYMENTTERM       api.APIOperation = "financial_bulkDeletePaymentTerm"
	BULK_CHANGE_PAYMENTTERM_OWNER api.APIOperation = "financial_bulkChangePaymentTermOwner"

	// READ APIs.
	LIST_PAYMENTTERMS     api.APIOperation = "financial_listPaymentTerms"
	COUNT_PAYMENTTERMS    api.APIOperation = "financial_countPaymentTerms"
	GET_PAYMENTTERM_BY_ID api.APIOperation = "cfinancial_getPaymentTermById"
)
