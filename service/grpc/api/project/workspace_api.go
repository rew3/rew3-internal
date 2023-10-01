package project

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Workspace
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_WORKSPACE                   api.APIOperation = "project_addWorkspace"
	UPDATE_WORKSPACE                 api.APIOperation = "project_updateWorkspace"
	DELETE_WORKSPACE                 api.APIOperation = "project_deleteWorkspace"
	CLONE_WORKSPACE                  api.APIOperation = "project_cloneWorkspace"
	CHANGE_WORKSPACE_OWNER           api.APIOperation = "project_changeWorkspaceOwner"
	BULK_ADD_WORKSPACE          api.APIOperation = "project_bulkAddWorkspace"
	BULK_UPDATE_WORKSPACE           api.APIOperation = "project_bulkUpdateWorkspace"
	BULK_DELETE_WORKSPACE           api.APIOperation = "project_bulkDeleteWorkspace"
	BULK_CHANGE_WORKSPACE_OWNER           api.APIOperation = "project_bulkChangeWorkspaceOwner"

	// READ APIs.
	LIST_WORKSPACES     api.APIOperation = "project_listWorkspaces"
	COUNT_WORKSPACES    api.APIOperation = "project_countWorkspaces"
	GET_WORKSPACE_BY_ID api.APIOperation = "cproject_getWorkspaceById"
    
)