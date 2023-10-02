package purchase

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Vendor
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_VENDOR                   api.APIOperation = "financial_addVendor"
	UPDATE_VENDOR                 api.APIOperation = "financial_updateVendor"
	DELETE_VENDOR                 api.APIOperation = "financial_deleteVendor"
	CLONE_VENDOR                  api.APIOperation = "financial_cloneVendor"
	CHANGE_VENDOR_OWNER           api.APIOperation = "financial_changeVendorOwner"
	BULK_ADD_VENDOR          api.APIOperation = "financial_bulkAddVendor"
	BULK_UPDATE_VENDOR           api.APIOperation = "financial_bulkUpdateVendor"
	BULK_DELETE_VENDOR           api.APIOperation = "financial_bulkDeleteVendor"
	BULK_CHANGE_VENDOR_OWNER           api.APIOperation = "financial_bulkChangeVendorOwner"

	// READ APIs.
	LIST_VENDORS     api.APIOperation = "financial_listVendors"
	COUNT_VENDORS    api.APIOperation = "financial_countVendors"
	GET_VENDOR_BY_ID api.APIOperation = "cfinancial_getVendorById"
    
)