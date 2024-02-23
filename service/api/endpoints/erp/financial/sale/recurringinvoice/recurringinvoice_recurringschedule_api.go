package recurringinvoice

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for RecurringSchedule
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_RECURRINGSCHEDULE               endpoints.Endpoint = "financial_addRecurringSchedule"
	UPDATE_RECURRINGSCHEDULE            endpoints.Endpoint = "financial_updateRecurringSchedule"
	DELETE_RECURRINGSCHEDULE            endpoints.Endpoint = "financial_deleteRecurringSchedule"
	CLONE_RECURRINGSCHEDULE             endpoints.Endpoint = "financial_cloneRecurringSchedule"
	CHANGE_RECURRINGSCHEDULE_OWNER      endpoints.Endpoint = "financial_changeRecurringScheduleOwner"
	BULK_ADD_RECURRINGSCHEDULE          endpoints.Endpoint = "financial_bulkAddRecurringSchedule"
	BULK_UPDATE_RECURRINGSCHEDULE       endpoints.Endpoint = "financial_bulkUpdateRecurringSchedule"
	BULK_DELETE_RECURRINGSCHEDULE       endpoints.Endpoint = "financial_bulkDeleteRecurringSchedule"
	BULK_CHANGE_RECURRINGSCHEDULE_OWNER endpoints.Endpoint = "financial_bulkChangeRecurringScheduleOwner"

	// READ APIs.
	LIST_RECURRINGSCHEDULES     endpoints.Endpoint = "financial_listRecurringSchedules"
	COUNT_RECURRINGSCHEDULES    endpoints.Endpoint = "financial_countRecurringSchedules"
	GET_RECURRINGSCHEDULE_BY_ID endpoints.Endpoint = "cfinancial_getRecurringScheduleById"
)
