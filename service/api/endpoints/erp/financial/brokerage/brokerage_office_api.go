package brokerage

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Office
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_OFFICE               endpoints.Endpoint = "financial_addOffice"
	UPDATE_OFFICE            endpoints.Endpoint = "financial_updateOffice"
	DELETE_OFFICE            endpoints.Endpoint = "financial_deleteOffice"
	CLONE_OFFICE             endpoints.Endpoint = "financial_cloneOffice"
	CHANGE_OFFICE_OWNER      endpoints.Endpoint = "financial_changeOfficeOwner"
	BULK_ADD_OFFICE          endpoints.Endpoint = "financial_bulkAddOffice"
	BULK_UPDATE_OFFICE       endpoints.Endpoint = "financial_bulkUpdateOffice"
	BULK_DELETE_OFFICE       endpoints.Endpoint = "financial_bulkDeleteOffice"
	BULK_CHANGE_OFFICE_OWNER endpoints.Endpoint = "financial_bulkChangeOfficeOwner"

	// READ APIs.
	LIST_OFFICES     endpoints.Endpoint = "financial_listOffices"
	COUNT_OFFICES    endpoints.Endpoint = "financial_countOffices"
	GET_OFFICE_BY_ID endpoints.Endpoint = "cfinancial_getOfficeById"
)
