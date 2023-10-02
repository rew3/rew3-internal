package dossier

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for DossierCollaborator
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_DOSSIERCOLLABORATOR                   api.APIOperation = "dms_addDossierCollaborator"
	UPDATE_DOSSIERCOLLABORATOR                 api.APIOperation = "dms_updateDossierCollaborator"
	DELETE_DOSSIERCOLLABORATOR                 api.APIOperation = "dms_deleteDossierCollaborator"
	CLONE_DOSSIERCOLLABORATOR                  api.APIOperation = "dms_cloneDossierCollaborator"
	CHANGE_DOSSIERCOLLABORATOR_OWNER           api.APIOperation = "dms_changeDossierCollaboratorOwner"
	BULK_ADD_DOSSIERCOLLABORATOR          api.APIOperation = "dms_bulkAddDossierCollaborator"
	BULK_UPDATE_DOSSIERCOLLABORATOR           api.APIOperation = "dms_bulkUpdateDossierCollaborator"
	BULK_DELETE_DOSSIERCOLLABORATOR           api.APIOperation = "dms_bulkDeleteDossierCollaborator"
	BULK_CHANGE_DOSSIERCOLLABORATOR_OWNER           api.APIOperation = "dms_bulkChangeDossierCollaboratorOwner"

	// READ APIs.
	LIST_DOSSIERCOLLABORATORS     api.APIOperation = "dms_listDossierCollaborators"
	COUNT_DOSSIERCOLLABORATORS    api.APIOperation = "dms_countDossierCollaborators"
	GET_DOSSIERCOLLABORATOR_BY_ID api.APIOperation = "cdms_getDossierCollaboratorById"
    
)