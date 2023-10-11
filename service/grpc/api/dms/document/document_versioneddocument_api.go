package document

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for VersionedDocument
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_VERSIONEDDOCUMENT               api.APIOperation = "dms_addVersionedDocument"
	UPDATE_VERSIONEDDOCUMENT            api.APIOperation = "dms_updateVersionedDocument"
	DELETE_VERSIONEDDOCUMENT            api.APIOperation = "dms_deleteVersionedDocument"
	CLONE_VERSIONEDDOCUMENT             api.APIOperation = "dms_cloneVersionedDocument"
	CHANGE_VERSIONEDDOCUMENT_OWNER      api.APIOperation = "dms_changeVersionedDocumentOwner"
	BULK_ADD_VERSIONEDDOCUMENT          api.APIOperation = "dms_bulkAddVersionedDocument"
	BULK_UPDATE_VERSIONEDDOCUMENT       api.APIOperation = "dms_bulkUpdateVersionedDocument"
	BULK_DELETE_VERSIONEDDOCUMENT       api.APIOperation = "dms_bulkDeleteVersionedDocument"
	BULK_CHANGE_VERSIONEDDOCUMENT_OWNER api.APIOperation = "dms_bulkChangeVersionedDocumentOwner"

	// READ APIs.
	LIST_VERSIONEDDOCUMENTS     api.APIOperation = "dms_listVersionedDocuments"
	COUNT_VERSIONEDDOCUMENTS    api.APIOperation = "dms_countVersionedDocuments"
	GET_VERSIONEDDOCUMENT_BY_ID api.APIOperation = "cdms_getVersionedDocumentById"
)
