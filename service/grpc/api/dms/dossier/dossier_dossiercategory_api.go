package dossier

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for DossierCategory
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_DOSSIERCATEGORY               api.APIOperation = "dms_addDossierCategory"
	UPDATE_DOSSIERCATEGORY            api.APIOperation = "dms_updateDossierCategory"
	DELETE_DOSSIERCATEGORY            api.APIOperation = "dms_deleteDossierCategory"
	CLONE_DOSSIERCATEGORY             api.APIOperation = "dms_cloneDossierCategory"
	CHANGE_DOSSIERCATEGORY_OWNER      api.APIOperation = "dms_changeDossierCategoryOwner"
	BULK_ADD_DOSSIERCATEGORY          api.APIOperation = "dms_bulkAddDossierCategory"
	BULK_UPDATE_DOSSIERCATEGORY       api.APIOperation = "dms_bulkUpdateDossierCategory"
	BULK_DELETE_DOSSIERCATEGORY       api.APIOperation = "dms_bulkDeleteDossierCategory"
	BULK_CHANGE_DOSSIERCATEGORY_OWNER api.APIOperation = "dms_bulkChangeDossierCategoryOwner"

	// READ APIs.
	LIST_DOSSIERCATEGORIES    api.APIOperation = "dms_listDossierCategories"
	COUNT_DOSSIERCATEGORIES   api.APIOperation = "dms_countDossierCategories"
	GET_DOSSIERCATEGORY_BY_ID api.APIOperation = "cdms_getDossierCategoryById"
)
