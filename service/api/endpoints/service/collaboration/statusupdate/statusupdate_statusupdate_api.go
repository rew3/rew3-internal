package statusupdate

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for StatusUpdate
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_STATUSUPDATE               endpoints.Endpoint = "collaboration_addStatusUpdate"
	UPDATE_STATUSUPDATE            endpoints.Endpoint = "collaboration_updateStatusUpdate"
	DELETE_STATUSUPDATE            endpoints.Endpoint = "collaboration_deleteStatusUpdate"
	CLONE_STATUSUPDATE             endpoints.Endpoint = "collaboration_cloneStatusUpdate"
	CHANGE_STATUSUPDATE_OWNER      endpoints.Endpoint = "collaboration_changeStatusUpdateOwner"
	BULK_ADD_STATUSUPDATE          endpoints.Endpoint = "collaboration_bulkAddStatusUpdate"
	BULK_UPDATE_STATUSUPDATE       endpoints.Endpoint = "collaboration_bulkUpdateStatusUpdate"
	BULK_DELETE_STATUSUPDATE       endpoints.Endpoint = "collaboration_bulkDeleteStatusUpdate"
	BULK_CHANGE_STATUSUPDATE_OWNER endpoints.Endpoint = "collaboration_bulkChangeStatusUpdateOwner"

	// READ APIs.
	LIST_STATUSUPDATES     endpoints.Endpoint = "collaboration_listStatusUpdates"
	COUNT_STATUSUPDATES    endpoints.Endpoint = "collaboration_countStatusUpdates"
	GET_STATUSUPDATE_BY_ID endpoints.Endpoint = "ccollaboration_getStatusUpdateById"
)
