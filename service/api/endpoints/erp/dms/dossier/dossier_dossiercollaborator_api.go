package dossier

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for DossierCollaborator
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DOSSIERCOLLABORATOR               endpoints.Endpoint = "dms_addDossierCollaborator"
	UPDATE_DOSSIERCOLLABORATOR            endpoints.Endpoint = "dms_updateDossierCollaborator"
	DELETE_DOSSIERCOLLABORATOR            endpoints.Endpoint = "dms_deleteDossierCollaborator"
	CLONE_DOSSIERCOLLABORATOR             endpoints.Endpoint = "dms_cloneDossierCollaborator"
	CHANGE_DOSSIERCOLLABORATOR_OWNER      endpoints.Endpoint = "dms_changeDossierCollaboratorOwner"
	BULK_ADD_DOSSIERCOLLABORATOR          endpoints.Endpoint = "dms_bulkAddDossierCollaborator"
	BULK_UPDATE_DOSSIERCOLLABORATOR       endpoints.Endpoint = "dms_bulkUpdateDossierCollaborator"
	BULK_DELETE_DOSSIERCOLLABORATOR       endpoints.Endpoint = "dms_bulkDeleteDossierCollaborator"
	BULK_CHANGE_DOSSIERCOLLABORATOR_OWNER endpoints.Endpoint = "dms_bulkChangeDossierCollaboratorOwner"

	// READ APIs.
	LIST_DOSSIERCOLLABORATORS     endpoints.Endpoint = "dms_listDossierCollaborators"
	COUNT_DOSSIERCOLLABORATORS    endpoints.Endpoint = "dms_countDossierCollaborators"
	GET_DOSSIERCOLLABORATOR_BY_ID endpoints.Endpoint = "cdms_getDossierCollaboratorById"
)
