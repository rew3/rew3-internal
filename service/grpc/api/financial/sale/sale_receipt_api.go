package sale

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Receipt
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_RECEIPT               api.APIOperation = "financial_addReceipt"
	UPDATE_RECEIPT            api.APIOperation = "financial_updateReceipt"
	DELETE_RECEIPT            api.APIOperation = "financial_deleteReceipt"
	CLONE_RECEIPT             api.APIOperation = "financial_cloneReceipt"
	CHANGE_RECEIPT_OWNER      api.APIOperation = "financial_changeReceiptOwner"
	BULK_ADD_RECEIPT          api.APIOperation = "financial_bulkAddReceipt"
	BULK_UPDATE_RECEIPT       api.APIOperation = "financial_bulkUpdateReceipt"
	BULK_DELETE_RECEIPT       api.APIOperation = "financial_bulkDeleteReceipt"
	BULK_CHANGE_RECEIPT_OWNER api.APIOperation = "financial_bulkChangeReceiptOwner"

	// READ APIs.
	LIST_RECEIPTS     api.APIOperation = "financial_listReceipts"
	COUNT_RECEIPTS    api.APIOperation = "financial_countReceipts"
	GET_RECEIPT_BY_ID api.APIOperation = "cfinancial_getReceiptById"
)
