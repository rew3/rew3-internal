package checklist

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Checklist
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_CHECKLIST               endpoints.Endpoint = "collaboration_addChecklist"
	UPDATE_CHECKLIST            endpoints.Endpoint = "collaboration_updateChecklist"
	DELETE_CHECKLIST            endpoints.Endpoint = "collaboration_deleteChecklist"
	CLONE_CHECKLIST             endpoints.Endpoint = "collaboration_cloneChecklist"
	CHANGE_CHECKLIST_OWNER      endpoints.Endpoint = "collaboration_changeChecklistOwner"
	BULK_ADD_CHECKLIST          endpoints.Endpoint = "collaboration_bulkAddChecklist"
	BULK_UPDATE_CHECKLIST       endpoints.Endpoint = "collaboration_bulkUpdateChecklist"
	BULK_DELETE_CHECKLIST       endpoints.Endpoint = "collaboration_bulkDeleteChecklist"
	BULK_CHANGE_CHECKLIST_OWNER endpoints.Endpoint = "collaboration_bulkChangeChecklistOwner"

	// READ APIs.
	LIST_CHECKLISTS     endpoints.Endpoint = "collaboration_listChecklists"
	COUNT_CHECKLISTS    endpoints.Endpoint = "collaboration_countChecklists"
	GET_CHECKLIST_BY_ID endpoints.Endpoint = "ccollaboration_getChecklistById"
)
