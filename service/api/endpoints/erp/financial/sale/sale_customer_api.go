package sale

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Customer
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_CUSTOMER               endpoints.Endpoint = "financial_addCustomer"
	UPDATE_CUSTOMER            endpoints.Endpoint = "financial_updateCustomer"
	DELETE_CUSTOMER            endpoints.Endpoint = "financial_deleteCustomer"
	CLONE_CUSTOMER             endpoints.Endpoint = "financial_cloneCustomer"
	CHANGE_CUSTOMER_OWNER      endpoints.Endpoint = "financial_changeCustomerOwner"
	BULK_ADD_CUSTOMER          endpoints.Endpoint = "financial_bulkAddCustomer"
	BULK_UPDATE_CUSTOMER       endpoints.Endpoint = "financial_bulkUpdateCustomer"
	BULK_DELETE_CUSTOMER       endpoints.Endpoint = "financial_bulkDeleteCustomer"
	BULK_CHANGE_CUSTOMER_OWNER endpoints.Endpoint = "financial_bulkChangeCustomerOwner"

	// READ APIs.
	LIST_CUSTOMERS     endpoints.Endpoint = "financial_listCustomers"
	COUNT_CUSTOMERS    endpoints.Endpoint = "financial_countCustomers"
	GET_CUSTOMER_BY_ID endpoints.Endpoint = "cfinancial_getCustomerById"
)
