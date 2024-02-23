package activity

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for UserActivity
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_USERACTIVITY         endpoints.Endpoint = "activity_addUserActivity"
	UPDATE_USERACTIVITY      endpoints.Endpoint = "activity_updateUserActivity"
	DELETE_USERACTIVITY      endpoints.Endpoint = "activity_deleteUserActivity"
	BULK_ADD_USERACTIVITY    endpoints.Endpoint = "activity_bulkAddUserActivity"
	BULK_UPDATE_USERACTIVITY endpoints.Endpoint = "activity_bulkUpdateUserActivity"
	BULK_DELETE_USERACTIVITY endpoints.Endpoint = "activity_bulkDeleteUserActivity"

	// READ APIs.
	LIST_USERACTIVITIES    endpoints.Endpoint = "activity_listUserActivities"
	COUNT_USERACTIVITIES   endpoints.Endpoint = "activity_countUserActivities"
	GET_USERACTIVITY_BY_ID endpoints.Endpoint = "cactivity_getUserActivityById"
)
