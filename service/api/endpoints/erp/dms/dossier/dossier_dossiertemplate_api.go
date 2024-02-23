package dossier

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for DossierTemplate
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DOSSIERTEMPLATE               endpoints.Endpoint = "dms_addDossierTemplate"
	UPDATE_DOSSIERTEMPLATE            endpoints.Endpoint = "dms_updateDossierTemplate"
	DELETE_DOSSIERTEMPLATE            endpoints.Endpoint = "dms_deleteDossierTemplate"
	CLONE_DOSSIERTEMPLATE             endpoints.Endpoint = "dms_cloneDossierTemplate"
	CHANGE_DOSSIERTEMPLATE_OWNER      endpoints.Endpoint = "dms_changeDossierTemplateOwner"
	BULK_ADD_DOSSIERTEMPLATE          endpoints.Endpoint = "dms_bulkAddDossierTemplate"
	BULK_UPDATE_DOSSIERTEMPLATE       endpoints.Endpoint = "dms_bulkUpdateDossierTemplate"
	BULK_DELETE_DOSSIERTEMPLATE       endpoints.Endpoint = "dms_bulkDeleteDossierTemplate"
	BULK_CHANGE_DOSSIERTEMPLATE_OWNER endpoints.Endpoint = "dms_bulkChangeDossierTemplateOwner"

	// READ APIs.
	LIST_DOSSIERTEMPLATES     endpoints.Endpoint = "dms_listDossierTemplates"
	COUNT_DOSSIERTEMPLATES    endpoints.Endpoint = "dms_countDossierTemplates"
	GET_DOSSIERTEMPLATE_BY_ID endpoints.Endpoint = "cdms_getDossierTemplateById"
)
