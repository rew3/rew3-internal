package asset

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Asset
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ASSET               api.APIOperation = "collaboration_addAsset"
	UPDATE_ASSET            api.APIOperation = "collaboration_updateAsset"
	DELETE_ASSET            api.APIOperation = "collaboration_deleteAsset"
	CLONE_ASSET             api.APIOperation = "collaboration_cloneAsset"
	CHANGE_ASSET_OWNER      api.APIOperation = "collaboration_changeAssetOwner"
	BULK_ADD_ASSET          api.APIOperation = "collaboration_bulkAddAsset"
	BULK_UPDATE_ASSET       api.APIOperation = "collaboration_bulkUpdateAsset"
	BULK_DELETE_ASSET       api.APIOperation = "collaboration_bulkDeleteAsset"
	BULK_CHANGE_ASSET_OWNER api.APIOperation = "collaboration_bulkChangeAssetOwner"

	// READ APIs.
	LIST_ASSETS     api.APIOperation = "collaboration_listAssets"
	COUNT_ASSETS    api.APIOperation = "collaboration_countAssets"
	GET_ASSET_BY_ID api.APIOperation = "ccollaboration_getAssetById"
)
