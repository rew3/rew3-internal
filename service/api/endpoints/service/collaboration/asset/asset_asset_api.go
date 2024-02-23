package asset

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Asset
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ASSET               endpoints.Endpoint = "collaboration_addAsset"
	UPDATE_ASSET            endpoints.Endpoint = "collaboration_updateAsset"
	DELETE_ASSET            endpoints.Endpoint = "collaboration_deleteAsset"
	CLONE_ASSET             endpoints.Endpoint = "collaboration_cloneAsset"
	CHANGE_ASSET_OWNER      endpoints.Endpoint = "collaboration_changeAssetOwner"
	BULK_ADD_ASSET          endpoints.Endpoint = "collaboration_bulkAddAsset"
	BULK_UPDATE_ASSET       endpoints.Endpoint = "collaboration_bulkUpdateAsset"
	BULK_DELETE_ASSET       endpoints.Endpoint = "collaboration_bulkDeleteAsset"
	BULK_CHANGE_ASSET_OWNER endpoints.Endpoint = "collaboration_bulkChangeAssetOwner"

	// READ APIs.
	LIST_ASSETS     endpoints.Endpoint = "collaboration_listAssets"
	COUNT_ASSETS    endpoints.Endpoint = "collaboration_countAssets"
	GET_ASSET_BY_ID endpoints.Endpoint = "ccollaboration_getAssetById"
)
