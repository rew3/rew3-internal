package recurringinvoice

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for RecurringSchedule
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_RECURRINGSCHEDULE                   api.APIOperation = "financial_addRecurringSchedule"
	UPDATE_RECURRINGSCHEDULE                 api.APIOperation = "financial_updateRecurringSchedule"
	DELETE_RECURRINGSCHEDULE                 api.APIOperation = "financial_deleteRecurringSchedule"
	CLONE_RECURRINGSCHEDULE                  api.APIOperation = "financial_cloneRecurringSchedule"
	CHANGE_RECURRINGSCHEDULE_OWNER           api.APIOperation = "financial_changeRecurringScheduleOwner"
	BULK_ADD_RECURRINGSCHEDULE          api.APIOperation = "financial_bulkAddRecurringSchedule"
	BULK_UPDATE_RECURRINGSCHEDULE           api.APIOperation = "financial_bulkUpdateRecurringSchedule"
	BULK_DELETE_RECURRINGSCHEDULE           api.APIOperation = "financial_bulkDeleteRecurringSchedule"
	BULK_CHANGE_RECURRINGSCHEDULE_OWNER           api.APIOperation = "financial_bulkChangeRecurringScheduleOwner"

	// READ APIs.
	LIST_RECURRINGSCHEDULES     api.APIOperation = "financial_listRecurringSchedules"
	COUNT_RECURRINGSCHEDULES    api.APIOperation = "financial_countRecurringSchedules"
	GET_RECURRINGSCHEDULE_BY_ID api.APIOperation = "cfinancial_getRecurringScheduleById"
    
)