package payment

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for BillPayment
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_BILLPAYMENT               endpoints.Endpoint = "financial_addBillPayment"
	UPDATE_BILLPAYMENT            endpoints.Endpoint = "financial_updateBillPayment"
	DELETE_BILLPAYMENT            endpoints.Endpoint = "financial_deleteBillPayment"
	CLONE_BILLPAYMENT             endpoints.Endpoint = "financial_cloneBillPayment"
	CHANGE_BILLPAYMENT_OWNER      endpoints.Endpoint = "financial_changeBillPaymentOwner"
	BULK_ADD_BILLPAYMENT          endpoints.Endpoint = "financial_bulkAddBillPayment"
	BULK_UPDATE_BILLPAYMENT       endpoints.Endpoint = "financial_bulkUpdateBillPayment"
	BULK_DELETE_BILLPAYMENT       endpoints.Endpoint = "financial_bulkDeleteBillPayment"
	BULK_CHANGE_BILLPAYMENT_OWNER endpoints.Endpoint = "financial_bulkChangeBillPaymentOwner"

	// READ APIs.
	LIST_BILLPAYMENTS     endpoints.Endpoint = "financial_listBillPayments"
	COUNT_BILLPAYMENTS    endpoints.Endpoint = "financial_countBillPayments"
	GET_BILLPAYMENT_BY_ID endpoints.Endpoint = "cfinancial_getBillPaymentById"
)
