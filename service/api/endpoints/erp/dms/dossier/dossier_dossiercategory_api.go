package dossier

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for DossierCategory
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DOSSIERCATEGORY               endpoints.Endpoint = "dms_addDossierCategory"
	UPDATE_DOSSIERCATEGORY            endpoints.Endpoint = "dms_updateDossierCategory"
	DELETE_DOSSIERCATEGORY            endpoints.Endpoint = "dms_deleteDossierCategory"
	CLONE_DOSSIERCATEGORY             endpoints.Endpoint = "dms_cloneDossierCategory"
	CHANGE_DOSSIERCATEGORY_OWNER      endpoints.Endpoint = "dms_changeDossierCategoryOwner"
	BULK_ADD_DOSSIERCATEGORY          endpoints.Endpoint = "dms_bulkAddDossierCategory"
	BULK_UPDATE_DOSSIERCATEGORY       endpoints.Endpoint = "dms_bulkUpdateDossierCategory"
	BULK_DELETE_DOSSIERCATEGORY       endpoints.Endpoint = "dms_bulkDeleteDossierCategory"
	BULK_CHANGE_DOSSIERCATEGORY_OWNER endpoints.Endpoint = "dms_bulkChangeDossierCategoryOwner"

	// READ APIs.
	LIST_DOSSIERCATEGORIES    endpoints.Endpoint = "dms_listDossierCategories"
	COUNT_DOSSIERCATEGORIES   endpoints.Endpoint = "dms_countDossierCategories"
	GET_DOSSIERCATEGORY_BY_ID endpoints.Endpoint = "cdms_getDossierCategoryById"
)
