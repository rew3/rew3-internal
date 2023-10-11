package document

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for DocumentAccess
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DOCUMENTACCESS               api.APIOperation = "dms_addDocumentAccess"
	UPDATE_DOCUMENTACCESS            api.APIOperation = "dms_updateDocumentAccess"
	DELETE_DOCUMENTACCESS            api.APIOperation = "dms_deleteDocumentAccess"
	CLONE_DOCUMENTACCESS             api.APIOperation = "dms_cloneDocumentAccess"
	CHANGE_DOCUMENTACCESS_OWNER      api.APIOperation = "dms_changeDocumentAccessOwner"
	BULK_ADD_DOCUMENTACCESS          api.APIOperation = "dms_bulkAddDocumentAccess"
	BULK_UPDATE_DOCUMENTACCESS       api.APIOperation = "dms_bulkUpdateDocumentAccess"
	BULK_DELETE_DOCUMENTACCESS       api.APIOperation = "dms_bulkDeleteDocumentAccess"
	BULK_CHANGE_DOCUMENTACCESS_OWNER api.APIOperation = "dms_bulkChangeDocumentAccessOwner"

	// READ APIs.
	LIST_DOCUMENTACCESSES    api.APIOperation = "dms_listDocumentAccesses"
	COUNT_DOCUMENTACCESSES   api.APIOperation = "dms_countDocumentAccesses"
	GET_DOCUMENTACCESS_BY_ID api.APIOperation = "cdms_getDocumentAccessById"
)
