package dossier

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Dossier
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DOSSIER               endpoints.Endpoint = "dms_addDossier"
	UPDATE_DOSSIER            endpoints.Endpoint = "dms_updateDossier"
	DELETE_DOSSIER            endpoints.Endpoint = "dms_deleteDossier"
	CLONE_DOSSIER             endpoints.Endpoint = "dms_cloneDossier"
	CHANGE_DOSSIER_OWNER      endpoints.Endpoint = "dms_changeDossierOwner"
	BULK_ADD_DOSSIER          endpoints.Endpoint = "dms_bulkAddDossier"
	BULK_UPDATE_DOSSIER       endpoints.Endpoint = "dms_bulkUpdateDossier"
	BULK_DELETE_DOSSIER       endpoints.Endpoint = "dms_bulkDeleteDossier"
	BULK_CHANGE_DOSSIER_OWNER endpoints.Endpoint = "dms_bulkChangeDossierOwner"

	// READ APIs.
	LIST_DOSSIERS     endpoints.Endpoint = "dms_listDossiers"
	COUNT_DOSSIERS    endpoints.Endpoint = "dms_countDossiers"
	GET_DOSSIER_BY_ID endpoints.Endpoint = "cdms_getDossierById"
)
