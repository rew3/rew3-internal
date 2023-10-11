package payment

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for BillPayment
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_BILLPAYMENT               api.APIOperation = "financial_addBillPayment"
	UPDATE_BILLPAYMENT            api.APIOperation = "financial_updateBillPayment"
	DELETE_BILLPAYMENT            api.APIOperation = "financial_deleteBillPayment"
	CLONE_BILLPAYMENT             api.APIOperation = "financial_cloneBillPayment"
	CHANGE_BILLPAYMENT_OWNER      api.APIOperation = "financial_changeBillPaymentOwner"
	BULK_ADD_BILLPAYMENT          api.APIOperation = "financial_bulkAddBillPayment"
	BULK_UPDATE_BILLPAYMENT       api.APIOperation = "financial_bulkUpdateBillPayment"
	BULK_DELETE_BILLPAYMENT       api.APIOperation = "financial_bulkDeleteBillPayment"
	BULK_CHANGE_BILLPAYMENT_OWNER api.APIOperation = "financial_bulkChangeBillPaymentOwner"

	// READ APIs.
	LIST_BILLPAYMENTS     api.APIOperation = "financial_listBillPayments"
	COUNT_BILLPAYMENTS    api.APIOperation = "financial_countBillPayments"
	GET_BILLPAYMENT_BY_ID api.APIOperation = "cfinancial_getBillPaymentById"
)
