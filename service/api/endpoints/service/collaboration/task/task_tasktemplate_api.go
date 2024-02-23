package task

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for TaskTemplate
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_TASKTEMPLATE               endpoints.Endpoint = "collaboration_addTaskTemplate"
	UPDATE_TASKTEMPLATE            endpoints.Endpoint = "collaboration_updateTaskTemplate"
	DELETE_TASKTEMPLATE            endpoints.Endpoint = "collaboration_deleteTaskTemplate"
	CLONE_TASKTEMPLATE             endpoints.Endpoint = "collaboration_cloneTaskTemplate"
	CHANGE_TASKTEMPLATE_OWNER      endpoints.Endpoint = "collaboration_changeTaskTemplateOwner"
	BULK_ADD_TASKTEMPLATE          endpoints.Endpoint = "collaboration_bulkAddTaskTemplate"
	BULK_UPDATE_TASKTEMPLATE       endpoints.Endpoint = "collaboration_bulkUpdateTaskTemplate"
	BULK_DELETE_TASKTEMPLATE       endpoints.Endpoint = "collaboration_bulkDeleteTaskTemplate"
	BULK_CHANGE_TASKTEMPLATE_OWNER endpoints.Endpoint = "collaboration_bulkChangeTaskTemplateOwner"

	// READ APIs.
	LIST_TASKTEMPLATES     endpoints.Endpoint = "collaboration_listTaskTemplates"
	COUNT_TASKTEMPLATES    endpoints.Endpoint = "collaboration_countTaskTemplates"
	GET_TASKTEMPLATE_BY_ID endpoints.Endpoint = "ccollaboration_getTaskTemplateById"
)
