package statusupdate

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for StatusUpdate
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_STATUSUPDATE               api.APIOperation = "collaboration_addStatusUpdate"
	UPDATE_STATUSUPDATE            api.APIOperation = "collaboration_updateStatusUpdate"
	DELETE_STATUSUPDATE            api.APIOperation = "collaboration_deleteStatusUpdate"
	CLONE_STATUSUPDATE             api.APIOperation = "collaboration_cloneStatusUpdate"
	CHANGE_STATUSUPDATE_OWNER      api.APIOperation = "collaboration_changeStatusUpdateOwner"
	BULK_ADD_STATUSUPDATE          api.APIOperation = "collaboration_bulkAddStatusUpdate"
	BULK_UPDATE_STATUSUPDATE       api.APIOperation = "collaboration_bulkUpdateStatusUpdate"
	BULK_DELETE_STATUSUPDATE       api.APIOperation = "collaboration_bulkDeleteStatusUpdate"
	BULK_CHANGE_STATUSUPDATE_OWNER api.APIOperation = "collaboration_bulkChangeStatusUpdateOwner"

	// READ APIs.
	LIST_STATUSUPDATES     api.APIOperation = "collaboration_listStatusUpdates"
	COUNT_STATUSUPDATES    api.APIOperation = "collaboration_countStatusUpdates"
	GET_STATUSUPDATE_BY_ID api.APIOperation = "ccollaboration_getStatusUpdateById"
)
