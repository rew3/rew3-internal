package event

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Event
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_EVENT               api.APIOperation = "collaboration_addEvent"
	UPDATE_EVENT            api.APIOperation = "collaboration_updateEvent"
	DELETE_EVENT            api.APIOperation = "collaboration_deleteEvent"
	CLONE_EVENT             api.APIOperation = "collaboration_cloneEvent"
	CHANGE_EVENT_OWNER      api.APIOperation = "collaboration_changeEventOwner"
	BULK_ADD_EVENT          api.APIOperation = "collaboration_bulkAddEvent"
	BULK_UPDATE_EVENT       api.APIOperation = "collaboration_bulkUpdateEvent"
	BULK_DELETE_EVENT       api.APIOperation = "collaboration_bulkDeleteEvent"
	BULK_CHANGE_EVENT_OWNER api.APIOperation = "collaboration_bulkChangeEventOwner"

	// READ APIs.
	LIST_EVENTS     api.APIOperation = "collaboration_listEvents"
	COUNT_EVENTS    api.APIOperation = "collaboration_countEvents"
	GET_EVENT_BY_ID api.APIOperation = "ccollaboration_getEventById"
)
