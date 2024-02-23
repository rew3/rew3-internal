package file

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for File
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_FILE               endpoints.Endpoint = "dms_addFile"
	UPDATE_FILE            endpoints.Endpoint = "dms_updateFile"
	DELETE_FILE            endpoints.Endpoint = "dms_deleteFile"
	CLONE_FILE             endpoints.Endpoint = "dms_cloneFile"
	CHANGE_FILE_OWNER      endpoints.Endpoint = "dms_changeFileOwner"
	BULK_ADD_FILE          endpoints.Endpoint = "dms_bulkAddFile"
	BULK_UPDATE_FILE       endpoints.Endpoint = "dms_bulkUpdateFile"
	BULK_DELETE_FILE       endpoints.Endpoint = "dms_bulkDeleteFile"
	BULK_CHANGE_FILE_OWNER endpoints.Endpoint = "dms_bulkChangeFileOwner"

	// READ APIs.
	LIST_FILES     endpoints.Endpoint = "dms_listFiles"
	COUNT_FILES    endpoints.Endpoint = "dms_countFiles"
	GET_FILE_BY_ID endpoints.Endpoint = "cdms_getFileById"
)
