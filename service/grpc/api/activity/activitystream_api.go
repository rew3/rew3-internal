package activity

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for ActivityStream
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ACTIVITYSTREAM    api.APIOperation = "activity_addActivityStream"
	UPDATE_ACTIVITYSTREAM api.APIOperation = "activity_updateActivityStream"
	DELETE_ACTIVITYSTREAM api.APIOperation = "activity_deleteActivityStream"

	// READ APIs.
	LIST_ACTIVITYSTREAMS     api.APIOperation = "activity_listActivityStreams"
	COUNT_ACTIVITYSTREAMS    api.APIOperation = "activity_countActivityStreams"
	GET_ACTIVITYSTREAM_BY_ID api.APIOperation = "cactivity_getActivityStreamById"
)
