package dossier

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for DossierFolder
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DOSSIERFOLDER               api.APIOperation = "dms_addDossierFolder"
	UPDATE_DOSSIERFOLDER            api.APIOperation = "dms_updateDossierFolder"
	DELETE_DOSSIERFOLDER            api.APIOperation = "dms_deleteDossierFolder"
	CLONE_DOSSIERFOLDER             api.APIOperation = "dms_cloneDossierFolder"
	CHANGE_DOSSIERFOLDER_OWNER      api.APIOperation = "dms_changeDossierFolderOwner"
	BULK_ADD_DOSSIERFOLDER          api.APIOperation = "dms_bulkAddDossierFolder"
	BULK_UPDATE_DOSSIERFOLDER       api.APIOperation = "dms_bulkUpdateDossierFolder"
	BULK_DELETE_DOSSIERFOLDER       api.APIOperation = "dms_bulkDeleteDossierFolder"
	BULK_CHANGE_DOSSIERFOLDER_OWNER api.APIOperation = "dms_bulkChangeDossierFolderOwner"

	// READ APIs.
	LIST_DOSSIERFOLDERS     api.APIOperation = "dms_listDossierFolders"
	COUNT_DOSSIERFOLDERS    api.APIOperation = "dms_countDossierFolders"
	GET_DOSSIERFOLDER_BY_ID api.APIOperation = "cdms_getDossierFolderById"
)
