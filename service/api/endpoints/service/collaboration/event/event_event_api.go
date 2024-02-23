package event

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Event
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_EVENT               endpoints.Endpoint = "collaboration_addEvent"
	UPDATE_EVENT            endpoints.Endpoint = "collaboration_updateEvent"
	DELETE_EVENT            endpoints.Endpoint = "collaboration_deleteEvent"
	CLONE_EVENT             endpoints.Endpoint = "collaboration_cloneEvent"
	CHANGE_EVENT_OWNER      endpoints.Endpoint = "collaboration_changeEventOwner"
	BULK_ADD_EVENT          endpoints.Endpoint = "collaboration_bulkAddEvent"
	BULK_UPDATE_EVENT       endpoints.Endpoint = "collaboration_bulkUpdateEvent"
	BULK_DELETE_EVENT       endpoints.Endpoint = "collaboration_bulkDeleteEvent"
	BULK_CHANGE_EVENT_OWNER endpoints.Endpoint = "collaboration_bulkChangeEventOwner"

	// READ APIs.
	LIST_EVENTS     endpoints.Endpoint = "collaboration_listEvents"
	COUNT_EVENTS    endpoints.Endpoint = "collaboration_countEvents"
	GET_EVENT_BY_ID endpoints.Endpoint = "ccollaboration_getEventById"
)
