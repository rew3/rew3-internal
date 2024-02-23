package event

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for EventInvitation
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_EVENTINVITATION               endpoints.Endpoint = "collaboration_addEventInvitation"
	UPDATE_EVENTINVITATION            endpoints.Endpoint = "collaboration_updateEventInvitation"
	DELETE_EVENTINVITATION            endpoints.Endpoint = "collaboration_deleteEventInvitation"
	CLONE_EVENTINVITATION             endpoints.Endpoint = "collaboration_cloneEventInvitation"
	CHANGE_EVENTINVITATION_OWNER      endpoints.Endpoint = "collaboration_changeEventInvitationOwner"
	BULK_ADD_EVENTINVITATION          endpoints.Endpoint = "collaboration_bulkAddEventInvitation"
	BULK_UPDATE_EVENTINVITATION       endpoints.Endpoint = "collaboration_bulkUpdateEventInvitation"
	BULK_DELETE_EVENTINVITATION       endpoints.Endpoint = "collaboration_bulkDeleteEventInvitation"
	BULK_CHANGE_EVENTINVITATION_OWNER endpoints.Endpoint = "collaboration_bulkChangeEventInvitationOwner"

	// READ APIs.
	LIST_EVENTINVITATIONS     endpoints.Endpoint = "collaboration_listEventInvitations"
	COUNT_EVENTINVITATIONS    endpoints.Endpoint = "collaboration_countEventInvitations"
	GET_EVENTINVITATION_BY_ID endpoints.Endpoint = "ccollaboration_getEventInvitationById"
)
