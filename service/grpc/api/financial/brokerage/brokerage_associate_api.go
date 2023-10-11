package brokerage

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Associate
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ASSOCIATE               api.APIOperation = "financial_addAssociate"
	UPDATE_ASSOCIATE            api.APIOperation = "financial_updateAssociate"
	DELETE_ASSOCIATE            api.APIOperation = "financial_deleteAssociate"
	CLONE_ASSOCIATE             api.APIOperation = "financial_cloneAssociate"
	CHANGE_ASSOCIATE_OWNER      api.APIOperation = "financial_changeAssociateOwner"
	BULK_ADD_ASSOCIATE          api.APIOperation = "financial_bulkAddAssociate"
	BULK_UPDATE_ASSOCIATE       api.APIOperation = "financial_bulkUpdateAssociate"
	BULK_DELETE_ASSOCIATE       api.APIOperation = "financial_bulkDeleteAssociate"
	BULK_CHANGE_ASSOCIATE_OWNER api.APIOperation = "financial_bulkChangeAssociateOwner"

	// READ APIs.
	LIST_ASSOCIATES     api.APIOperation = "financial_listAssociates"
	COUNT_ASSOCIATES    api.APIOperation = "financial_countAssociates"
	GET_ASSOCIATE_BY_ID api.APIOperation = "cfinancial_getAssociateById"
)
