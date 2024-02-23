package catalog

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Product
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_PRODUCT               endpoints.Endpoint = "financial_addProduct"
	UPDATE_PRODUCT            endpoints.Endpoint = "financial_updateProduct"
	DELETE_PRODUCT            endpoints.Endpoint = "financial_deleteProduct"
	CLONE_PRODUCT             endpoints.Endpoint = "financial_cloneProduct"
	CHANGE_PRODUCT_OWNER      endpoints.Endpoint = "financial_changeProductOwner"
	BULK_ADD_PRODUCT          endpoints.Endpoint = "financial_bulkAddProduct"
	BULK_UPDATE_PRODUCT       endpoints.Endpoint = "financial_bulkUpdateProduct"
	BULK_DELETE_PRODUCT       endpoints.Endpoint = "financial_bulkDeleteProduct"
	BULK_CHANGE_PRODUCT_OWNER endpoints.Endpoint = "financial_bulkChangeProductOwner"

	// READ APIs.
	LIST_PRODUCTS     endpoints.Endpoint = "financial_listProducts"
	COUNT_PRODUCTS    endpoints.Endpoint = "financial_countProducts"
	GET_PRODUCT_BY_ID endpoints.Endpoint = "cfinancial_getProductById"
)
