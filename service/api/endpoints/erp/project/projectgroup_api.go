package project

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for ProjectGroup
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_PROJECTGROUP               endpoints.Endpoint = "project_addProjectGroup"
	UPDATE_PROJECTGROUP            endpoints.Endpoint = "project_updateProjectGroup"
	DELETE_PROJECTGROUP            endpoints.Endpoint = "project_deleteProjectGroup"
	CLONE_PROJECTGROUP             endpoints.Endpoint = "project_cloneProjectGroup"
	CHANGE_PROJECTGROUP_OWNER      endpoints.Endpoint = "project_changeProjectGroupOwner"
	BULK_ADD_PROJECTGROUP          endpoints.Endpoint = "project_bulkAddProjectGroup"
	BULK_UPDATE_PROJECTGROUP       endpoints.Endpoint = "project_bulkUpdateProjectGroup"
	BULK_DELETE_PROJECTGROUP       endpoints.Endpoint = "project_bulkDeleteProjectGroup"
	BULK_CHANGE_PROJECTGROUP_OWNER endpoints.Endpoint = "project_bulkChangeProjectGroupOwner"

	// READ APIs.
	LIST_PROJECTGROUPS     endpoints.Endpoint = "project_listProjectGroups"
	COUNT_PROJECTGROUPS    endpoints.Endpoint = "project_countProjectGroups"
	GET_PROJECTGROUP_BY_ID endpoints.Endpoint = "cproject_getProjectGroupById"
)
