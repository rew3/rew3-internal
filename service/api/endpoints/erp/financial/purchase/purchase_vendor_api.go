package purchase

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Vendor
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_VENDOR               endpoints.Endpoint = "financial_addVendor"
	UPDATE_VENDOR            endpoints.Endpoint = "financial_updateVendor"
	DELETE_VENDOR            endpoints.Endpoint = "financial_deleteVendor"
	CLONE_VENDOR             endpoints.Endpoint = "financial_cloneVendor"
	CHANGE_VENDOR_OWNER      endpoints.Endpoint = "financial_changeVendorOwner"
	BULK_ADD_VENDOR          endpoints.Endpoint = "financial_bulkAddVendor"
	BULK_UPDATE_VENDOR       endpoints.Endpoint = "financial_bulkUpdateVendor"
	BULK_DELETE_VENDOR       endpoints.Endpoint = "financial_bulkDeleteVendor"
	BULK_CHANGE_VENDOR_OWNER endpoints.Endpoint = "financial_bulkChangeVendorOwner"

	// READ APIs.
	LIST_VENDORS     endpoints.Endpoint = "financial_listVendors"
	COUNT_VENDORS    endpoints.Endpoint = "financial_countVendors"
	GET_VENDOR_BY_ID endpoints.Endpoint = "cfinancial_getVendorById"
)
