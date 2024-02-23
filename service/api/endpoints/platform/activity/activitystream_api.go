package activity

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for ActivityStream
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ACTIVITYSTREAM    endpoints.Endpoint = "activity_addActivityStream"
	UPDATE_ACTIVITYSTREAM endpoints.Endpoint = "activity_updateActivityStream"
	DELETE_ACTIVITYSTREAM endpoints.Endpoint = "activity_deleteActivityStream"

	// READ APIs.
	LIST_ACTIVITYSTREAMS     endpoints.Endpoint = "activity_listActivityStreams"
	COUNT_ACTIVITYSTREAMS    endpoints.Endpoint = "activity_countActivityStreams"
	GET_ACTIVITYSTREAM_BY_ID endpoints.Endpoint = "cactivity_getActivityStreamById"
)
