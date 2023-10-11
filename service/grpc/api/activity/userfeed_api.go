package activity

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for UserFeed
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_USERFEED               api.APIOperation = "activity_addUserFeed"
	UPDATE_USERFEED            api.APIOperation = "activity_updateUserFeed"
	DELETE_USERFEED            api.APIOperation = "activity_deleteUserFeed"

	// READ APIs.
	LIST_USERFEEDS     api.APIOperation = "activity_listUserFeeds"
	COUNT_USERFEEDS    api.APIOperation = "activity_countUserFeeds"
	GET_USERFEED_BY_ID api.APIOperation = "cactivity_getUserFeedById"
)
