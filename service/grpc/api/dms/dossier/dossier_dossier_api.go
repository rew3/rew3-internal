package dossier

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Dossier
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DOSSIER               api.APIOperation = "dms_addDossier"
	UPDATE_DOSSIER            api.APIOperation = "dms_updateDossier"
	DELETE_DOSSIER            api.APIOperation = "dms_deleteDossier"
	CLONE_DOSSIER             api.APIOperation = "dms_cloneDossier"
	CHANGE_DOSSIER_OWNER      api.APIOperation = "dms_changeDossierOwner"
	BULK_ADD_DOSSIER          api.APIOperation = "dms_bulkAddDossier"
	BULK_UPDATE_DOSSIER       api.APIOperation = "dms_bulkUpdateDossier"
	BULK_DELETE_DOSSIER       api.APIOperation = "dms_bulkDeleteDossier"
	BULK_CHANGE_DOSSIER_OWNER api.APIOperation = "dms_bulkChangeDossierOwner"

	// READ APIs.
	LIST_DOSSIERS     api.APIOperation = "dms_listDossiers"
	COUNT_DOSSIERS    api.APIOperation = "dms_countDossiers"
	GET_DOSSIER_BY_ID api.APIOperation = "cdms_getDossierById"
)
