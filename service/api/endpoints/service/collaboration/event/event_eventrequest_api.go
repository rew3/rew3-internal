package event

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for EventRequest
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_EVENTREQUEST               endpoints.Endpoint = "collaboration_addEventRequest"
	UPDATE_EVENTREQUEST            endpoints.Endpoint = "collaboration_updateEventRequest"
	DELETE_EVENTREQUEST            endpoints.Endpoint = "collaboration_deleteEventRequest"
	CLONE_EVENTREQUEST             endpoints.Endpoint = "collaboration_cloneEventRequest"
	CHANGE_EVENTREQUEST_OWNER      endpoints.Endpoint = "collaboration_changeEventRequestOwner"
	BULK_ADD_EVENTREQUEST          endpoints.Endpoint = "collaboration_bulkAddEventRequest"
	BULK_UPDATE_EVENTREQUEST       endpoints.Endpoint = "collaboration_bulkUpdateEventRequest"
	BULK_DELETE_EVENTREQUEST       endpoints.Endpoint = "collaboration_bulkDeleteEventRequest"
	BULK_CHANGE_EVENTREQUEST_OWNER endpoints.Endpoint = "collaboration_bulkChangeEventRequestOwner"

	// READ APIs.
	LIST_EVENTREQUESTS     endpoints.Endpoint = "collaboration_listEventRequests"
	COUNT_EVENTREQUESTS    endpoints.Endpoint = "collaboration_countEventRequests"
	GET_EVENTREQUEST_BY_ID endpoints.Endpoint = "ccollaboration_getEventRequestById"
)
