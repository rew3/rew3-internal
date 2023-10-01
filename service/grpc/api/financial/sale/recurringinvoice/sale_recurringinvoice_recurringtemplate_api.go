package sale/recurringinvoice

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for RecurringTemplate
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_RECURRINGTEMPLATE                   api.APIOperation = "financial_addRecurringTemplate"
	UPDATE_RECURRINGTEMPLATE                 api.APIOperation = "financial_updateRecurringTemplate"
	DELETE_RECURRINGTEMPLATE                 api.APIOperation = "financial_deleteRecurringTemplate"
	CLONE_RECURRINGTEMPLATE                  api.APIOperation = "financial_cloneRecurringTemplate"
	CHANGE_RECURRINGTEMPLATE_OWNER           api.APIOperation = "financial_changeRecurringTemplateOwner"
	BULK_ADD_RECURRINGTEMPLATE          api.APIOperation = "financial_bulkAddRecurringTemplate"
	BULK_UPDATE_RECURRINGTEMPLATE           api.APIOperation = "financial_bulkUpdateRecurringTemplate"
	BULK_DELETE_RECURRINGTEMPLATE           api.APIOperation = "financial_bulkDeleteRecurringTemplate"
	BULK_CHANGE_RECURRINGTEMPLATE_OWNER           api.APIOperation = "financial_bulkChangeRecurringTemplateOwner"

	// READ APIs.
	LIST_RECURRINGTEMPLATES     api.APIOperation = "financial_listRecurringTemplates"
	COUNT_RECURRINGTEMPLATES    api.APIOperation = "financial_countRecurringTemplates"
	GET_RECURRINGTEMPLATE_BY_ID api.APIOperation = "cfinancial_getRecurringTemplateById"
    
)