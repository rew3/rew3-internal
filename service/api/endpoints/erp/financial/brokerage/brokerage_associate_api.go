package brokerage

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Associate
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ASSOCIATE               endpoints.Endpoint = "financial_addAssociate"
	UPDATE_ASSOCIATE            endpoints.Endpoint = "financial_updateAssociate"
	DELETE_ASSOCIATE            endpoints.Endpoint = "financial_deleteAssociate"
	CLONE_ASSOCIATE             endpoints.Endpoint = "financial_cloneAssociate"
	CHANGE_ASSOCIATE_OWNER      endpoints.Endpoint = "financial_changeAssociateOwner"
	BULK_ADD_ASSOCIATE          endpoints.Endpoint = "financial_bulkAddAssociate"
	BULK_UPDATE_ASSOCIATE       endpoints.Endpoint = "financial_bulkUpdateAssociate"
	BULK_DELETE_ASSOCIATE       endpoints.Endpoint = "financial_bulkDeleteAssociate"
	BULK_CHANGE_ASSOCIATE_OWNER endpoints.Endpoint = "financial_bulkChangeAssociateOwner"

	// READ APIs.
	LIST_ASSOCIATES     endpoints.Endpoint = "financial_listAssociates"
	COUNT_ASSOCIATES    endpoints.Endpoint = "financial_countAssociates"
	GET_ASSOCIATE_BY_ID endpoints.Endpoint = "cfinancial_getAssociateById"
)
