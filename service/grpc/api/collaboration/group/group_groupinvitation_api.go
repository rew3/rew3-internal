package group

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for GroupInvitation
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_GROUPINVITATION                   api.APIOperation = "collaboration_addGroupInvitation"
	UPDATE_GROUPINVITATION                 api.APIOperation = "collaboration_updateGroupInvitation"
	DELETE_GROUPINVITATION                 api.APIOperation = "collaboration_deleteGroupInvitation"
	CLONE_GROUPINVITATION                  api.APIOperation = "collaboration_cloneGroupInvitation"
	CHANGE_GROUPINVITATION_OWNER           api.APIOperation = "collaboration_changeGroupInvitationOwner"
	BULK_ADD_GROUPINVITATION          api.APIOperation = "collaboration_bulkAddGroupInvitation"
	BULK_UPDATE_GROUPINVITATION           api.APIOperation = "collaboration_bulkUpdateGroupInvitation"
	BULK_DELETE_GROUPINVITATION           api.APIOperation = "collaboration_bulkDeleteGroupInvitation"
	BULK_CHANGE_GROUPINVITATION_OWNER           api.APIOperation = "collaboration_bulkChangeGroupInvitationOwner"

	// READ APIs.
	LIST_GROUPINVITATIONS     api.APIOperation = "collaboration_listGroupInvitations"
	COUNT_GROUPINVITATIONS    api.APIOperation = "collaboration_countGroupInvitations"
	GET_GROUPINVITATION_BY_ID api.APIOperation = "ccollaboration_getGroupInvitationById"
    
)