package document

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Document
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DOCUMENT               api.APIOperation = "dms_addDocument"
	UPDATE_DOCUMENT            api.APIOperation = "dms_updateDocument"
	DELETE_DOCUMENT            api.APIOperation = "dms_deleteDocument"
	CLONE_DOCUMENT             api.APIOperation = "dms_cloneDocument"
	CHANGE_DOCUMENT_OWNER      api.APIOperation = "dms_changeDocumentOwner"
	BULK_ADD_DOCUMENT          api.APIOperation = "dms_bulkAddDocument"
	BULK_UPDATE_DOCUMENT       api.APIOperation = "dms_bulkUpdateDocument"
	BULK_DELETE_DOCUMENT       api.APIOperation = "dms_bulkDeleteDocument"
	BULK_CHANGE_DOCUMENT_OWNER api.APIOperation = "dms_bulkChangeDocumentOwner"

	// READ APIs.
	LIST_DOCUMENTS     api.APIOperation = "dms_listDocuments"
	COUNT_DOCUMENTS    api.APIOperation = "dms_countDocuments"
	GET_DOCUMENT_BY_ID api.APIOperation = "cdms_getDocumentById"
)
