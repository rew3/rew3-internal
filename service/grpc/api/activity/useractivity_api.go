package activity

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for UserActivity
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_USERACTIVITY         api.APIOperation = "activity_addUserActivity"
	UPDATE_USERACTIVITY      api.APIOperation = "activity_updateUserActivity"
	DELETE_USERACTIVITY      api.APIOperation = "activity_deleteUserActivity"
	BULK_ADD_USERACTIVITY    api.APIOperation = "activity_bulkAddUserActivity"
	BULK_UPDATE_USERACTIVITY api.APIOperation = "activity_bulkUpdateUserActivity"
	BULK_DELETE_USERACTIVITY api.APIOperation = "activity_bulkDeleteUserActivity"

	// READ APIs.
	LIST_USERACTIVITIES    api.APIOperation = "activity_listUserActivities"
	COUNT_USERACTIVITIES   api.APIOperation = "activity_countUserActivities"
	GET_USERACTIVITY_BY_ID api.APIOperation = "cactivity_getUserActivityById"
)
