package document

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for VersionedDocument
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_VERSIONEDDOCUMENT               endpoints.Endpoint = "dms_addVersionedDocument"
	UPDATE_VERSIONEDDOCUMENT            endpoints.Endpoint = "dms_updateVersionedDocument"
	DELETE_VERSIONEDDOCUMENT            endpoints.Endpoint = "dms_deleteVersionedDocument"
	CLONE_VERSIONEDDOCUMENT             endpoints.Endpoint = "dms_cloneVersionedDocument"
	CHANGE_VERSIONEDDOCUMENT_OWNER      endpoints.Endpoint = "dms_changeVersionedDocumentOwner"
	BULK_ADD_VERSIONEDDOCUMENT          endpoints.Endpoint = "dms_bulkAddVersionedDocument"
	BULK_UPDATE_VERSIONEDDOCUMENT       endpoints.Endpoint = "dms_bulkUpdateVersionedDocument"
	BULK_DELETE_VERSIONEDDOCUMENT       endpoints.Endpoint = "dms_bulkDeleteVersionedDocument"
	BULK_CHANGE_VERSIONEDDOCUMENT_OWNER endpoints.Endpoint = "dms_bulkChangeVersionedDocumentOwner"

	// READ APIs.
	LIST_VERSIONEDDOCUMENTS     endpoints.Endpoint = "dms_listVersionedDocuments"
	COUNT_VERSIONEDDOCUMENTS    endpoints.Endpoint = "dms_countVersionedDocuments"
	GET_VERSIONEDDOCUMENT_BY_ID endpoints.Endpoint = "cdms_getVersionedDocumentById"
)
