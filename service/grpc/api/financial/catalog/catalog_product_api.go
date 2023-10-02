package catalog

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Product
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_PRODUCT                   api.APIOperation = "financial_addProduct"
	UPDATE_PRODUCT                 api.APIOperation = "financial_updateProduct"
	DELETE_PRODUCT                 api.APIOperation = "financial_deleteProduct"
	CLONE_PRODUCT                  api.APIOperation = "financial_cloneProduct"
	CHANGE_PRODUCT_OWNER           api.APIOperation = "financial_changeProductOwner"
	BULK_ADD_PRODUCT          api.APIOperation = "financial_bulkAddProduct"
	BULK_UPDATE_PRODUCT           api.APIOperation = "financial_bulkUpdateProduct"
	BULK_DELETE_PRODUCT           api.APIOperation = "financial_bulkDeleteProduct"
	BULK_CHANGE_PRODUCT_OWNER           api.APIOperation = "financial_bulkChangeProductOwner"

	// READ APIs.
	LIST_PRODUCTS     api.APIOperation = "financial_listProducts"
	COUNT_PRODUCTS    api.APIOperation = "financial_countProducts"
	GET_PRODUCT_BY_ID api.APIOperation = "cfinancial_getProductById"
    
)