package purchase

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Bill
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_BILL               api.APIOperation = "financial_addBill"
	UPDATE_BILL            api.APIOperation = "financial_updateBill"
	DELETE_BILL            api.APIOperation = "financial_deleteBill"
	CLONE_BILL             api.APIOperation = "financial_cloneBill"
	CHANGE_BILL_OWNER      api.APIOperation = "financial_changeBillOwner"
	BULK_ADD_BILL          api.APIOperation = "financial_bulkAddBill"
	BULK_UPDATE_BILL       api.APIOperation = "financial_bulkUpdateBill"
	BULK_DELETE_BILL       api.APIOperation = "financial_bulkDeleteBill"
	BULK_CHANGE_BILL_OWNER api.APIOperation = "financial_bulkChangeBillOwner"

	// READ APIs.
	LIST_BILLS     api.APIOperation = "financial_listBills"
	COUNT_BILLS    api.APIOperation = "financial_countBills"
	GET_BILL_BY_ID api.APIOperation = "cfinancial_getBillById"
)
