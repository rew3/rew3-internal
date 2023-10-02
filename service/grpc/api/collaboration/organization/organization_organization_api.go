package organization

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for Organization
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_ORGANIZATION                   api.APIOperation = "collaboration_addOrganization"
	UPDATE_ORGANIZATION                 api.APIOperation = "collaboration_updateOrganization"
	DELETE_ORGANIZATION                 api.APIOperation = "collaboration_deleteOrganization"
	CLONE_ORGANIZATION                  api.APIOperation = "collaboration_cloneOrganization"
	CHANGE_ORGANIZATION_OWNER           api.APIOperation = "collaboration_changeOrganizationOwner"
	BULK_ADD_ORGANIZATION          api.APIOperation = "collaboration_bulkAddOrganization"
	BULK_UPDATE_ORGANIZATION           api.APIOperation = "collaboration_bulkUpdateOrganization"
	BULK_DELETE_ORGANIZATION           api.APIOperation = "collaboration_bulkDeleteOrganization"
	BULK_CHANGE_ORGANIZATION_OWNER           api.APIOperation = "collaboration_bulkChangeOrganizationOwner"

	// READ APIs.
	LIST_ORGANIZATIONS     api.APIOperation = "collaboration_listOrganizations"
	COUNT_ORGANIZATIONS    api.APIOperation = "collaboration_countOrganizations"
	GET_ORGANIZATION_BY_ID api.APIOperation = "ccollaboration_getOrganizationById"
    
)