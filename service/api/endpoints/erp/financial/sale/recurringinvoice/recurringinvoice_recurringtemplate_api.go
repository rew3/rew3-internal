package recurringinvoice

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for RecurringTemplate
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_RECURRINGTEMPLATE               endpoints.Endpoint = "financial_addRecurringTemplate"
	UPDATE_RECURRINGTEMPLATE            endpoints.Endpoint = "financial_updateRecurringTemplate"
	DELETE_RECURRINGTEMPLATE            endpoints.Endpoint = "financial_deleteRecurringTemplate"
	CLONE_RECURRINGTEMPLATE             endpoints.Endpoint = "financial_cloneRecurringTemplate"
	CHANGE_RECURRINGTEMPLATE_OWNER      endpoints.Endpoint = "financial_changeRecurringTemplateOwner"
	BULK_ADD_RECURRINGTEMPLATE          endpoints.Endpoint = "financial_bulkAddRecurringTemplate"
	BULK_UPDATE_RECURRINGTEMPLATE       endpoints.Endpoint = "financial_bulkUpdateRecurringTemplate"
	BULK_DELETE_RECURRINGTEMPLATE       endpoints.Endpoint = "financial_bulkDeleteRecurringTemplate"
	BULK_CHANGE_RECURRINGTEMPLATE_OWNER endpoints.Endpoint = "financial_bulkChangeRecurringTemplateOwner"

	// READ APIs.
	LIST_RECURRINGTEMPLATES     endpoints.Endpoint = "financial_listRecurringTemplates"
	COUNT_RECURRINGTEMPLATES    endpoints.Endpoint = "financial_countRecurringTemplates"
	GET_RECURRINGTEMPLATE_BY_ID endpoints.Endpoint = "cfinancial_getRecurringTemplateById"
)
