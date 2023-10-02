package calendar

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Calendar
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_CALENDAR                   api.APIOperation = "collaboration_addCalendar"
	UPDATE_CALENDAR                 api.APIOperation = "collaboration_updateCalendar"
	DELETE_CALENDAR                 api.APIOperation = "collaboration_deleteCalendar"
	CLONE_CALENDAR                  api.APIOperation = "collaboration_cloneCalendar"
	CHANGE_CALENDAR_OWNER           api.APIOperation = "collaboration_changeCalendarOwner"
	BULK_ADD_CALENDAR          api.APIOperation = "collaboration_bulkAddCalendar"
	BULK_UPDATE_CALENDAR           api.APIOperation = "collaboration_bulkUpdateCalendar"
	BULK_DELETE_CALENDAR           api.APIOperation = "collaboration_bulkDeleteCalendar"
	BULK_CHANGE_CALENDAR_OWNER           api.APIOperation = "collaboration_bulkChangeCalendarOwner"

	// READ APIs.
	LIST_CALENDARS     api.APIOperation = "collaboration_listCalendars"
	COUNT_CALENDARS    api.APIOperation = "collaboration_countCalendars"
	GET_CALENDAR_BY_ID api.APIOperation = "ccollaboration_getCalendarById"
    
)