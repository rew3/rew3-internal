package sale

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Receipt
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_RECEIPT               endpoints.Endpoint = "financial_addReceipt"
	UPDATE_RECEIPT            endpoints.Endpoint = "financial_updateReceipt"
	DELETE_RECEIPT            endpoints.Endpoint = "financial_deleteReceipt"
	CLONE_RECEIPT             endpoints.Endpoint = "financial_cloneReceipt"
	CHANGE_RECEIPT_OWNER      endpoints.Endpoint = "financial_changeReceiptOwner"
	BULK_ADD_RECEIPT          endpoints.Endpoint = "financial_bulkAddReceipt"
	BULK_UPDATE_RECEIPT       endpoints.Endpoint = "financial_bulkUpdateReceipt"
	BULK_DELETE_RECEIPT       endpoints.Endpoint = "financial_bulkDeleteReceipt"
	BULK_CHANGE_RECEIPT_OWNER endpoints.Endpoint = "financial_bulkChangeReceiptOwner"

	// READ APIs.
	LIST_RECEIPTS     endpoints.Endpoint = "financial_listReceipts"
	COUNT_RECEIPTS    endpoints.Endpoint = "financial_countReceipts"
	GET_RECEIPT_BY_ID endpoints.Endpoint = "cfinancial_getReceiptById"
)
