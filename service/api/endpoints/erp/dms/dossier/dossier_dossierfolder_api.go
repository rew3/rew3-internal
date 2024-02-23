package dossier

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for DossierFolder
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DOSSIERFOLDER               endpoints.Endpoint = "dms_addDossierFolder"
	UPDATE_DOSSIERFOLDER            endpoints.Endpoint = "dms_updateDossierFolder"
	DELETE_DOSSIERFOLDER            endpoints.Endpoint = "dms_deleteDossierFolder"
	CLONE_DOSSIERFOLDER             endpoints.Endpoint = "dms_cloneDossierFolder"
	CHANGE_DOSSIERFOLDER_OWNER      endpoints.Endpoint = "dms_changeDossierFolderOwner"
	BULK_ADD_DOSSIERFOLDER          endpoints.Endpoint = "dms_bulkAddDossierFolder"
	BULK_UPDATE_DOSSIERFOLDER       endpoints.Endpoint = "dms_bulkUpdateDossierFolder"
	BULK_DELETE_DOSSIERFOLDER       endpoints.Endpoint = "dms_bulkDeleteDossierFolder"
	BULK_CHANGE_DOSSIERFOLDER_OWNER endpoints.Endpoint = "dms_bulkChangeDossierFolderOwner"

	// READ APIs.
	LIST_DOSSIERFOLDERS     endpoints.Endpoint = "dms_listDossierFolders"
	COUNT_DOSSIERFOLDERS    endpoints.Endpoint = "dms_countDossierFolders"
	GET_DOSSIERFOLDER_BY_ID endpoints.Endpoint = "cdms_getDossierFolderById"
)
