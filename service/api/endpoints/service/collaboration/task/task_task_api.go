package task

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Task
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_TASK               endpoints.Endpoint = "collaboration_addTask"
	UPDATE_TASK            endpoints.Endpoint = "collaboration_updateTask"
	DELETE_TASK            endpoints.Endpoint = "collaboration_deleteTask"
	CLONE_TASK             endpoints.Endpoint = "collaboration_cloneTask"
	CHANGE_TASK_OWNER      endpoints.Endpoint = "collaboration_changeTaskOwner"
	BULK_ADD_TASK          endpoints.Endpoint = "collaboration_bulkAddTask"
	BULK_UPDATE_TASK       endpoints.Endpoint = "collaboration_bulkUpdateTask"
	BULK_DELETE_TASK       endpoints.Endpoint = "collaboration_bulkDeleteTask"
	BULK_CHANGE_TASK_OWNER endpoints.Endpoint = "collaboration_bulkChangeTaskOwner"

	// READ APIs.
	LIST_TASKS     endpoints.Endpoint = "collaboration_listTasks"
	COUNT_TASKS    endpoints.Endpoint = "collaboration_countTasks"
	GET_TASK_BY_ID endpoints.Endpoint = "ccollaboration_getTaskById"
)
