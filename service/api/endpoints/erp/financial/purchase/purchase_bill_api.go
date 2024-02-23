package purchase

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Bill
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_BILL               endpoints.Endpoint = "financial_addBill"
	UPDATE_BILL            endpoints.Endpoint = "financial_updateBill"
	DELETE_BILL            endpoints.Endpoint = "financial_deleteBill"
	CLONE_BILL             endpoints.Endpoint = "financial_cloneBill"
	CHANGE_BILL_OWNER      endpoints.Endpoint = "financial_changeBillOwner"
	BULK_ADD_BILL          endpoints.Endpoint = "financial_bulkAddBill"
	BULK_UPDATE_BILL       endpoints.Endpoint = "financial_bulkUpdateBill"
	BULK_DELETE_BILL       endpoints.Endpoint = "financial_bulkDeleteBill"
	BULK_CHANGE_BILL_OWNER endpoints.Endpoint = "financial_bulkChangeBillOwner"

	// READ APIs.
	LIST_BILLS     endpoints.Endpoint = "financial_listBills"
	COUNT_BILLS    endpoints.Endpoint = "financial_countBills"
	GET_BILL_BY_ID endpoints.Endpoint = "cfinancial_getBillById"
)
