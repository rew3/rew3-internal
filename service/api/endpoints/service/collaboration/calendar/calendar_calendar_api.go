package calendar

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Calendar
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_CALENDAR               endpoints.Endpoint = "collaboration_addCalendar"
	UPDATE_CALENDAR            endpoints.Endpoint = "collaboration_updateCalendar"
	DELETE_CALENDAR            endpoints.Endpoint = "collaboration_deleteCalendar"
	CLONE_CALENDAR             endpoints.Endpoint = "collaboration_cloneCalendar"
	CHANGE_CALENDAR_OWNER      endpoints.Endpoint = "collaboration_changeCalendarOwner"
	BULK_ADD_CALENDAR          endpoints.Endpoint = "collaboration_bulkAddCalendar"
	BULK_UPDATE_CALENDAR       endpoints.Endpoint = "collaboration_bulkUpdateCalendar"
	BULK_DELETE_CALENDAR       endpoints.Endpoint = "collaboration_bulkDeleteCalendar"
	BULK_CHANGE_CALENDAR_OWNER endpoints.Endpoint = "collaboration_bulkChangeCalendarOwner"

	// READ APIs.
	LIST_CALENDARS     endpoints.Endpoint = "collaboration_listCalendars"
	COUNT_CALENDARS    endpoints.Endpoint = "collaboration_countCalendars"
	GET_CALENDAR_BY_ID endpoints.Endpoint = "ccollaboration_getCalendarById"
)
