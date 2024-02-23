package folder

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Folder
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_FOLDER               endpoints.Endpoint = "dms_addFolder"
	UPDATE_FOLDER            endpoints.Endpoint = "dms_updateFolder"
	DELETE_FOLDER            endpoints.Endpoint = "dms_deleteFolder"
	CLONE_FOLDER             endpoints.Endpoint = "dms_cloneFolder"
	CHANGE_FOLDER_OWNER      endpoints.Endpoint = "dms_changeFolderOwner"
	BULK_ADD_FOLDER          endpoints.Endpoint = "dms_bulkAddFolder"
	BULK_UPDATE_FOLDER       endpoints.Endpoint = "dms_bulkUpdateFolder"
	BULK_DELETE_FOLDER       endpoints.Endpoint = "dms_bulkDeleteFolder"
	BULK_CHANGE_FOLDER_OWNER endpoints.Endpoint = "dms_bulkChangeFolderOwner"

	// READ APIs.
	LIST_FOLDERS     endpoints.Endpoint = "dms_listFolders"
	COUNT_FOLDERS    endpoints.Endpoint = "dms_countFolders"
	GET_FOLDER_BY_ID endpoints.Endpoint = "cdms_getFolderById"
)
