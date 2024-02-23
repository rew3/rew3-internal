package document

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Document
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DOCUMENT               endpoints.Endpoint = "dms_addDocument"
	UPDATE_DOCUMENT            endpoints.Endpoint = "dms_updateDocument"
	DELETE_DOCUMENT            endpoints.Endpoint = "dms_deleteDocument"
	CLONE_DOCUMENT             endpoints.Endpoint = "dms_cloneDocument"
	CHANGE_DOCUMENT_OWNER      endpoints.Endpoint = "dms_changeDocumentOwner"
	BULK_ADD_DOCUMENT          endpoints.Endpoint = "dms_bulkAddDocument"
	BULK_UPDATE_DOCUMENT       endpoints.Endpoint = "dms_bulkUpdateDocument"
	BULK_DELETE_DOCUMENT       endpoints.Endpoint = "dms_bulkDeleteDocument"
	BULK_CHANGE_DOCUMENT_OWNER endpoints.Endpoint = "dms_bulkChangeDocumentOwner"

	// READ APIs.
	LIST_DOCUMENTS     endpoints.Endpoint = "dms_listDocuments"
	COUNT_DOCUMENTS    endpoints.Endpoint = "dms_countDocuments"
	GET_DOCUMENT_BY_ID endpoints.Endpoint = "cdms_getDocumentById"
)
