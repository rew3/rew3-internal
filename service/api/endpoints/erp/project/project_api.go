package project

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Project
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_PROJECT               endpoints.Endpoint = "project_addProject"
	UPDATE_PROJECT            endpoints.Endpoint = "project_updateProject"
	DELETE_PROJECT            endpoints.Endpoint = "project_deleteProject"
	CLONE_PROJECT             endpoints.Endpoint = "project_cloneProject"
	CHANGE_PROJECT_OWNER      endpoints.Endpoint = "project_changeProjectOwner"
	BULK_ADD_PROJECT          endpoints.Endpoint = "project_bulkAddProject"
	BULK_UPDATE_PROJECT       endpoints.Endpoint = "project_bulkUpdateProject"
	BULK_DELETE_PROJECT       endpoints.Endpoint = "project_bulkDeleteProject"
	BULK_CHANGE_PROJECT_OWNER endpoints.Endpoint = "project_bulkChangeProjectOwner"

	// READ APIs.
	LIST_PROJECT      endpoints.Endpoint = "project_listProject"
	COUNT_PROJECT     endpoints.Endpoint = "project_countProject"
	GET_PROJECT_BY_ID endpoints.Endpoint = "cproject_getProjectById"
)
