package task

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for TaskTemplate
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_TASKTEMPLATE               api.APIOperation = "collaboration_addTaskTemplate"
	UPDATE_TASKTEMPLATE            api.APIOperation = "collaboration_updateTaskTemplate"
	DELETE_TASKTEMPLATE            api.APIOperation = "collaboration_deleteTaskTemplate"
	CLONE_TASKTEMPLATE             api.APIOperation = "collaboration_cloneTaskTemplate"
	CHANGE_TASKTEMPLATE_OWNER      api.APIOperation = "collaboration_changeTaskTemplateOwner"
	BULK_ADD_TASKTEMPLATE          api.APIOperation = "collaboration_bulkAddTaskTemplate"
	BULK_UPDATE_TASKTEMPLATE       api.APIOperation = "collaboration_bulkUpdateTaskTemplate"
	BULK_DELETE_TASKTEMPLATE       api.APIOperation = "collaboration_bulkDeleteTaskTemplate"
	BULK_CHANGE_TASKTEMPLATE_OWNER api.APIOperation = "collaboration_bulkChangeTaskTemplateOwner"

	// READ APIs.
	LIST_TASKTEMPLATES     api.APIOperation = "collaboration_listTaskTemplates"
	COUNT_TASKTEMPLATES    api.APIOperation = "collaboration_countTaskTemplates"
	GET_TASKTEMPLATE_BY_ID api.APIOperation = "ccollaboration_getTaskTemplateById"
)
