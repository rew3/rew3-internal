package socialinterest

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for SocialInterest
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_SOCIALINTEREST               endpoints.Endpoint = "collaboration_addSocialInterest"
	UPDATE_SOCIALINTEREST            endpoints.Endpoint = "collaboration_updateSocialInterest"
	DELETE_SOCIALINTEREST            endpoints.Endpoint = "collaboration_deleteSocialInterest"
	CLONE_SOCIALINTEREST             endpoints.Endpoint = "collaboration_cloneSocialInterest"
	CHANGE_SOCIALINTEREST_OWNER      endpoints.Endpoint = "collaboration_changeSocialInterestOwner"
	BULK_ADD_SOCIALINTEREST          endpoints.Endpoint = "collaboration_bulkAddSocialInterest"
	BULK_UPDATE_SOCIALINTEREST       endpoints.Endpoint = "collaboration_bulkUpdateSocialInterest"
	BULK_DELETE_SOCIALINTEREST       endpoints.Endpoint = "collaboration_bulkDeleteSocialInterest"
	BULK_CHANGE_SOCIALINTEREST_OWNER endpoints.Endpoint = "collaboration_bulkChangeSocialInterestOwner"

	// READ APIs.
	LIST_SOCIALINTERESTS     endpoints.Endpoint = "collaboration_listSocialInterests"
	COUNT_SOCIALINTERESTS    endpoints.Endpoint = "collaboration_countSocialInterests"
	GET_SOCIALINTEREST_BY_ID endpoints.Endpoint = "ccollaboration_getSocialInterestById"
)
