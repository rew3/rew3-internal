package activity

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for ActivityLog
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ACTIVITYLOG    api.APIOperation = "activity_addActivityLog"
	UPDATE_ACTIVITYLOG api.APIOperation = "activity_updateActivityLog"
	DELETE_ACTIVITYLOG api.APIOperation = "activity_deleteActivityLog"

	// READ APIs.
	LIST_ACTIVITYLOGS     api.APIOperation = "activity_listActivityLogs"
	COUNT_ACTIVITYLOGS    api.APIOperation = "activity_countActivityLogs"
	GET_ACTIVITYLOG_BY_ID api.APIOperation = "cactivity_getActivityLogById"
)
