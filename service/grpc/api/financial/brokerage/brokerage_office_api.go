package brokerage

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Office
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_OFFICE               api.APIOperation = "financial_addOffice"
	UPDATE_OFFICE            api.APIOperation = "financial_updateOffice"
	DELETE_OFFICE            api.APIOperation = "financial_deleteOffice"
	CLONE_OFFICE             api.APIOperation = "financial_cloneOffice"
	CHANGE_OFFICE_OWNER      api.APIOperation = "financial_changeOfficeOwner"
	BULK_ADD_OFFICE          api.APIOperation = "financial_bulkAddOffice"
	BULK_UPDATE_OFFICE       api.APIOperation = "financial_bulkUpdateOffice"
	BULK_DELETE_OFFICE       api.APIOperation = "financial_bulkDeleteOffice"
	BULK_CHANGE_OFFICE_OWNER api.APIOperation = "financial_bulkChangeOfficeOwner"

	// READ APIs.
	LIST_OFFICES     api.APIOperation = "financial_listOffices"
	COUNT_OFFICES    api.APIOperation = "financial_countOffices"
	GET_OFFICE_BY_ID api.APIOperation = "cfinancial_getOfficeById"
)
