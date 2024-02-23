package organization

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for OrganizationInvite
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ORGANIZATIONINVITE               endpoints.Endpoint = "collaboration_addOrganizationInvite"
	UPDATE_ORGANIZATIONINVITE            endpoints.Endpoint = "collaboration_updateOrganizationInvite"
	DELETE_ORGANIZATIONINVITE            endpoints.Endpoint = "collaboration_deleteOrganizationInvite"
	CLONE_ORGANIZATIONINVITE             endpoints.Endpoint = "collaboration_cloneOrganizationInvite"
	CHANGE_ORGANIZATIONINVITE_OWNER      endpoints.Endpoint = "collaboration_changeOrganizationInviteOwner"
	BULK_ADD_ORGANIZATIONINVITE          endpoints.Endpoint = "collaboration_bulkAddOrganizationInvite"
	BULK_UPDATE_ORGANIZATIONINVITE       endpoints.Endpoint = "collaboration_bulkUpdateOrganizationInvite"
	BULK_DELETE_ORGANIZATIONINVITE       endpoints.Endpoint = "collaboration_bulkDeleteOrganizationInvite"
	BULK_CHANGE_ORGANIZATIONINVITE_OWNER endpoints.Endpoint = "collaboration_bulkChangeOrganizationInviteOwner"

	// READ APIs.
	LIST_ORGANIZATIONINVITES     endpoints.Endpoint = "collaboration_listOrganizationInvites"
	COUNT_ORGANIZATIONINVITES    endpoints.Endpoint = "collaboration_countOrganizationInvites"
	GET_ORGANIZATIONINVITE_BY_ID endpoints.Endpoint = "ccollaboration_getOrganizationInviteById"
)
