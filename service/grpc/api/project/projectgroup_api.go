package project

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for ProjectGroup
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_PROJECTGROUP                   api.APIOperation = "project_addProjectGroup"
	UPDATE_PROJECTGROUP                 api.APIOperation = "project_updateProjectGroup"
	DELETE_PROJECTGROUP                 api.APIOperation = "project_deleteProjectGroup"
	CLONE_PROJECTGROUP                  api.APIOperation = "project_cloneProjectGroup"
	CHANGE_PROJECTGROUP_OWNER           api.APIOperation = "project_changeProjectGroupOwner"
	BULK_ADD_PROJECTGROUP          api.APIOperation = "project_bulkAddProjectGroup"
	BULK_UPDATE_PROJECTGROUP           api.APIOperation = "project_bulkUpdateProjectGroup"
	BULK_DELETE_PROJECTGROUP           api.APIOperation = "project_bulkDeleteProjectGroup"
	BULK_CHANGE_PROJECTGROUP_OWNER           api.APIOperation = "project_bulkChangeProjectGroupOwner"

	// READ APIs.
	LIST_PROJECTGROUPS     api.APIOperation = "project_listProjectGroups"
	COUNT_PROJECTGROUPS    api.APIOperation = "project_countProjectGroups"
	GET_PROJECTGROUP_BY_ID api.APIOperation = "cproject_getProjectGroupById"
    
)