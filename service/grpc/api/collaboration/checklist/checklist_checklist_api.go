package checklist

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Checklist
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_CHECKLIST                   api.APIOperation = "collaboration_addChecklist"
	UPDATE_CHECKLIST                 api.APIOperation = "collaboration_updateChecklist"
	DELETE_CHECKLIST                 api.APIOperation = "collaboration_deleteChecklist"
	CLONE_CHECKLIST                  api.APIOperation = "collaboration_cloneChecklist"
	CHANGE_CHECKLIST_OWNER           api.APIOperation = "collaboration_changeChecklistOwner"
	BULK_ADD_CHECKLIST          api.APIOperation = "collaboration_bulkAddChecklist"
	BULK_UPDATE_CHECKLIST           api.APIOperation = "collaboration_bulkUpdateChecklist"
	BULK_DELETE_CHECKLIST           api.APIOperation = "collaboration_bulkDeleteChecklist"
	BULK_CHANGE_CHECKLIST_OWNER           api.APIOperation = "collaboration_bulkChangeChecklistOwner"

	// READ APIs.
	LIST_CHECKLISTS     api.APIOperation = "collaboration_listChecklists"
	COUNT_CHECKLISTS    api.APIOperation = "collaboration_countChecklists"
	GET_CHECKLIST_BY_ID api.APIOperation = "ccollaboration_getChecklistById"
    
)