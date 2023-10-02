package socialinterest

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for SocialInterest
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_SOCIALINTEREST                   api.APIOperation = "collaboration_addSocialInterest"
	UPDATE_SOCIALINTEREST                 api.APIOperation = "collaboration_updateSocialInterest"
	DELETE_SOCIALINTEREST                 api.APIOperation = "collaboration_deleteSocialInterest"
	CLONE_SOCIALINTEREST                  api.APIOperation = "collaboration_cloneSocialInterest"
	CHANGE_SOCIALINTEREST_OWNER           api.APIOperation = "collaboration_changeSocialInterestOwner"
	BULK_ADD_SOCIALINTEREST          api.APIOperation = "collaboration_bulkAddSocialInterest"
	BULK_UPDATE_SOCIALINTEREST           api.APIOperation = "collaboration_bulkUpdateSocialInterest"
	BULK_DELETE_SOCIALINTEREST           api.APIOperation = "collaboration_bulkDeleteSocialInterest"
	BULK_CHANGE_SOCIALINTEREST_OWNER           api.APIOperation = "collaboration_bulkChangeSocialInterestOwner"

	// READ APIs.
	LIST_SOCIALINTERESTS     api.APIOperation = "collaboration_listSocialInterests"
	COUNT_SOCIALINTERESTS    api.APIOperation = "collaboration_countSocialInterests"
	GET_SOCIALINTEREST_BY_ID api.APIOperation = "ccollaboration_getSocialInterestById"
    
)