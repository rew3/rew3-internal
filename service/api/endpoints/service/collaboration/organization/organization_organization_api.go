package organization

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for Organization
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_ORGANIZATION               endpoints.Endpoint = "collaboration_addOrganization"
	UPDATE_ORGANIZATION            endpoints.Endpoint = "collaboration_updateOrganization"
	DELETE_ORGANIZATION            endpoints.Endpoint = "collaboration_deleteOrganization"
	CLONE_ORGANIZATION             endpoints.Endpoint = "collaboration_cloneOrganization"
	CHANGE_ORGANIZATION_OWNER      endpoints.Endpoint = "collaboration_changeOrganizationOwner"
	BULK_ADD_ORGANIZATION          endpoints.Endpoint = "collaboration_bulkAddOrganization"
	BULK_UPDATE_ORGANIZATION       endpoints.Endpoint = "collaboration_bulkUpdateOrganization"
	BULK_DELETE_ORGANIZATION       endpoints.Endpoint = "collaboration_bulkDeleteOrganization"
	BULK_CHANGE_ORGANIZATION_OWNER endpoints.Endpoint = "collaboration_bulkChangeOrganizationOwner"

	// READ APIs.
	LIST_ORGANIZATIONS     endpoints.Endpoint = "collaboration_listOrganizations"
	COUNT_ORGANIZATIONS    endpoints.Endpoint = "collaboration_countOrganizations"
	GET_ORGANIZATION_BY_ID endpoints.Endpoint = "ccollaboration_getOrganizationById"
)
