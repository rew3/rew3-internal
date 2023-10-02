package dossier

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for DossierTemplate
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_DOSSIERTEMPLATE                   api.APIOperation = "dms_addDossierTemplate"
	UPDATE_DOSSIERTEMPLATE                 api.APIOperation = "dms_updateDossierTemplate"
	DELETE_DOSSIERTEMPLATE                 api.APIOperation = "dms_deleteDossierTemplate"
	CLONE_DOSSIERTEMPLATE                  api.APIOperation = "dms_cloneDossierTemplate"
	CHANGE_DOSSIERTEMPLATE_OWNER           api.APIOperation = "dms_changeDossierTemplateOwner"
	BULK_ADD_DOSSIERTEMPLATE          api.APIOperation = "dms_bulkAddDossierTemplate"
	BULK_UPDATE_DOSSIERTEMPLATE           api.APIOperation = "dms_bulkUpdateDossierTemplate"
	BULK_DELETE_DOSSIERTEMPLATE           api.APIOperation = "dms_bulkDeleteDossierTemplate"
	BULK_CHANGE_DOSSIERTEMPLATE_OWNER           api.APIOperation = "dms_bulkChangeDossierTemplateOwner"

	// READ APIs.
	LIST_DOSSIERTEMPLATES     api.APIOperation = "dms_listDossierTemplates"
	COUNT_DOSSIERTEMPLATES    api.APIOperation = "dms_countDossierTemplates"
	GET_DOSSIERTEMPLATE_BY_ID api.APIOperation = "cdms_getDossierTemplateById"
    
)