package event

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for EventInvitation
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_EVENTINVITATION                   api.APIOperation = "collaboration_addEventInvitation"
	UPDATE_EVENTINVITATION                 api.APIOperation = "collaboration_updateEventInvitation"
	DELETE_EVENTINVITATION                 api.APIOperation = "collaboration_deleteEventInvitation"
	CLONE_EVENTINVITATION                  api.APIOperation = "collaboration_cloneEventInvitation"
	CHANGE_EVENTINVITATION_OWNER           api.APIOperation = "collaboration_changeEventInvitationOwner"
	BULK_ADD_EVENTINVITATION          api.APIOperation = "collaboration_bulkAddEventInvitation"
	BULK_UPDATE_EVENTINVITATION           api.APIOperation = "collaboration_bulkUpdateEventInvitation"
	BULK_DELETE_EVENTINVITATION           api.APIOperation = "collaboration_bulkDeleteEventInvitation"
	BULK_CHANGE_EVENTINVITATION_OWNER           api.APIOperation = "collaboration_bulkChangeEventInvitationOwner"

	// READ APIs.
	LIST_EVENTINVITATIONS     api.APIOperation = "collaboration_listEventInvitations"
	COUNT_EVENTINVITATIONS    api.APIOperation = "collaboration_countEventInvitations"
	GET_EVENTINVITATION_BY_ID api.APIOperation = "ccollaboration_getEventInvitationById"
    
)