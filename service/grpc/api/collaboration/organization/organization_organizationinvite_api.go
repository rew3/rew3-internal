package organization

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// API for OrganizationInvite
    // ----------------------------------------------------
	// WRITE APIs. 
	ADD_ORGANIZATIONINVITE                   api.APIOperation = "collaboration_addOrganizationInvite"
	UPDATE_ORGANIZATIONINVITE                 api.APIOperation = "collaboration_updateOrganizationInvite"
	DELETE_ORGANIZATIONINVITE                 api.APIOperation = "collaboration_deleteOrganizationInvite"
	CLONE_ORGANIZATIONINVITE                  api.APIOperation = "collaboration_cloneOrganizationInvite"
	CHANGE_ORGANIZATIONINVITE_OWNER           api.APIOperation = "collaboration_changeOrganizationInviteOwner"
	BULK_ADD_ORGANIZATIONINVITE          api.APIOperation = "collaboration_bulkAddOrganizationInvite"
	BULK_UPDATE_ORGANIZATIONINVITE           api.APIOperation = "collaboration_bulkUpdateOrganizationInvite"
	BULK_DELETE_ORGANIZATIONINVITE           api.APIOperation = "collaboration_bulkDeleteOrganizationInvite"
	BULK_CHANGE_ORGANIZATIONINVITE_OWNER           api.APIOperation = "collaboration_bulkChangeOrganizationInviteOwner"

	// READ APIs.
	LIST_ORGANIZATIONINVITES     api.APIOperation = "collaboration_listOrganizationInvites"
	COUNT_ORGANIZATIONINVITES    api.APIOperation = "collaboration_countOrganizationInvites"
	GET_ORGANIZATIONINVITE_BY_ID api.APIOperation = "ccollaboration_getOrganizationInviteById"
    
)