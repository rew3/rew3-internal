package sale

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Customer
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_CUSTOMER               api.APIOperation = "financial_addCustomer"
	UPDATE_CUSTOMER            api.APIOperation = "financial_updateCustomer"
	DELETE_CUSTOMER            api.APIOperation = "financial_deleteCustomer"
	CLONE_CUSTOMER             api.APIOperation = "financial_cloneCustomer"
	CHANGE_CUSTOMER_OWNER      api.APIOperation = "financial_changeCustomerOwner"
	BULK_ADD_CUSTOMER          api.APIOperation = "financial_bulkAddCustomer"
	BULK_UPDATE_CUSTOMER       api.APIOperation = "financial_bulkUpdateCustomer"
	BULK_DELETE_CUSTOMER       api.APIOperation = "financial_bulkDeleteCustomer"
	BULK_CHANGE_CUSTOMER_OWNER api.APIOperation = "financial_bulkChangeCustomerOwner"

	// READ APIs.
	LIST_CUSTOMERS     api.APIOperation = "financial_listCustomers"
	COUNT_CUSTOMERS    api.APIOperation = "financial_countCustomers"
	GET_CUSTOMER_BY_ID api.APIOperation = "cfinancial_getCustomerById"
)
