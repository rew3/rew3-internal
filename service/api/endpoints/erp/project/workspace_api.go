package project

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Workspace
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_WORKSPACE               endpoints.Endpoint = "project_addWorkspace"
	UPDATE_WORKSPACE            endpoints.Endpoint = "project_updateWorkspace"
	DELETE_WORKSPACE            endpoints.Endpoint = "project_deleteWorkspace"
	CLONE_WORKSPACE             endpoints.Endpoint = "project_cloneWorkspace"
	CHANGE_WORKSPACE_OWNER      endpoints.Endpoint = "project_changeWorkspaceOwner"
	BULK_ADD_WORKSPACE          endpoints.Endpoint = "project_bulkAddWorkspace"
	BULK_UPDATE_WORKSPACE       endpoints.Endpoint = "project_bulkUpdateWorkspace"
	BULK_DELETE_WORKSPACE       endpoints.Endpoint = "project_bulkDeleteWorkspace"
	BULK_CHANGE_WORKSPACE_OWNER endpoints.Endpoint = "project_bulkChangeWorkspaceOwner"

	// READ APIs.
	LIST_WORKSPACES     endpoints.Endpoint = "project_listWorkspaces"
	COUNT_WORKSPACES    endpoints.Endpoint = "project_countWorkspaces"
	GET_WORKSPACE_BY_ID endpoints.Endpoint = "cproject_getWorkspaceById"
)
