package file

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for File
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_FILE               api.APIOperation = "dms_addFile"
	UPDATE_FILE            api.APIOperation = "dms_updateFile"
	DELETE_FILE            api.APIOperation = "dms_deleteFile"
	CLONE_FILE             api.APIOperation = "dms_cloneFile"
	CHANGE_FILE_OWNER      api.APIOperation = "dms_changeFileOwner"
	BULK_ADD_FILE          api.APIOperation = "dms_bulkAddFile"
	BULK_UPDATE_FILE       api.APIOperation = "dms_bulkUpdateFile"
	BULK_DELETE_FILE       api.APIOperation = "dms_bulkDeleteFile"
	BULK_CHANGE_FILE_OWNER api.APIOperation = "dms_bulkChangeFileOwner"

	// READ APIs.
	LIST_FILES     api.APIOperation = "dms_listFiles"
	COUNT_FILES    api.APIOperation = "dms_countFiles"
	GET_FILE_BY_ID api.APIOperation = "cdms_getFileById"
)
