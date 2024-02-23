package document

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for DocumentAccess
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DOCUMENTACCESS               endpoints.Endpoint = "dms_addDocumentAccess"
	UPDATE_DOCUMENTACCESS            endpoints.Endpoint = "dms_updateDocumentAccess"
	DELETE_DOCUMENTACCESS            endpoints.Endpoint = "dms_deleteDocumentAccess"
	CLONE_DOCUMENTACCESS             endpoints.Endpoint = "dms_cloneDocumentAccess"
	CHANGE_DOCUMENTACCESS_OWNER      endpoints.Endpoint = "dms_changeDocumentAccessOwner"
	BULK_ADD_DOCUMENTACCESS          endpoints.Endpoint = "dms_bulkAddDocumentAccess"
	BULK_UPDATE_DOCUMENTACCESS       endpoints.Endpoint = "dms_bulkUpdateDocumentAccess"
	BULK_DELETE_DOCUMENTACCESS       endpoints.Endpoint = "dms_bulkDeleteDocumentAccess"
	BULK_CHANGE_DOCUMENTACCESS_OWNER endpoints.Endpoint = "dms_bulkChangeDocumentAccessOwner"

	// READ APIs.
	LIST_DOCUMENTACCESSES    endpoints.Endpoint = "dms_listDocumentAccesses"
	COUNT_DOCUMENTACCESSES   endpoints.Endpoint = "dms_countDocumentAccesses"
	GET_DOCUMENTACCESS_BY_ID endpoints.Endpoint = "cdms_getDocumentAccessById"
)
