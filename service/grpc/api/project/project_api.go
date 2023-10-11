package project

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Project
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_PROJECT               api.APIOperation = "project_addProject"
	UPDATE_PROJECT            api.APIOperation = "project_updateProject"
	DELETE_PROJECT            api.APIOperation = "project_deleteProject"
	CLONE_PROJECT             api.APIOperation = "project_cloneProject"
	CHANGE_PROJECT_OWNER      api.APIOperation = "project_changeProjectOwner"
	BULK_ADD_PROJECT          api.APIOperation = "project_bulkAddProject"
	BULK_UPDATE_PROJECT       api.APIOperation = "project_bulkUpdateProject"
	BULK_DELETE_PROJECT       api.APIOperation = "project_bulkDeleteProject"
	BULK_CHANGE_PROJECT_OWNER api.APIOperation = "project_bulkChangeProjectOwner"

	// READ APIs.
	LIST_PROJECT      api.APIOperation = "project_listProject"
	COUNT_PROJECT     api.APIOperation = "project_countProject"
	GET_PROJECT_BY_ID api.APIOperation = "cproject_getProjectById"
)
