package task

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Task
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_TASK                   api.APIOperation = "collaboration_addTask"
	UPDATE_TASK                 api.APIOperation = "collaboration_updateTask"
	DELETE_TASK                 api.APIOperation = "collaboration_deleteTask"
	CLONE_TASK                  api.APIOperation = "collaboration_cloneTask"
	CHANGE_TASK_OWNER           api.APIOperation = "collaboration_changeTaskOwner"
	BULK_ADD_TASK          api.APIOperation = "collaboration_bulkAddTask"
	BULK_UPDATE_TASK           api.APIOperation = "collaboration_bulkUpdateTask"
	BULK_DELETE_TASK           api.APIOperation = "collaboration_bulkDeleteTask"
	BULK_CHANGE_TASK_OWNER           api.APIOperation = "collaboration_bulkChangeTaskOwner"

	// READ APIs.
	LIST_TASKS     api.APIOperation = "collaboration_listTasks"
	COUNT_TASKS    api.APIOperation = "collaboration_countTasks"
	GET_TASK_BY_ID api.APIOperation = "ccollaboration_getTaskById"
    
)