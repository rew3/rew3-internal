package folder

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Folder
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_FOLDER               api.APIOperation = "dms_addFolder"
	UPDATE_FOLDER            api.APIOperation = "dms_updateFolder"
	DELETE_FOLDER            api.APIOperation = "dms_deleteFolder"
	CLONE_FOLDER             api.APIOperation = "dms_cloneFolder"
	CHANGE_FOLDER_OWNER      api.APIOperation = "dms_changeFolderOwner"
	BULK_ADD_FOLDER          api.APIOperation = "dms_bulkAddFolder"
	BULK_UPDATE_FOLDER       api.APIOperation = "dms_bulkUpdateFolder"
	BULK_DELETE_FOLDER       api.APIOperation = "dms_bulkDeleteFolder"
	BULK_CHANGE_FOLDER_OWNER api.APIOperation = "dms_bulkChangeFolderOwner"

	// READ APIs.
	LIST_FOLDERS     api.APIOperation = "dms_listFolders"
	COUNT_FOLDERS    api.APIOperation = "dms_countFolders"
	GET_FOLDER_BY_ID api.APIOperation = "cdms_getFolderById"
)
