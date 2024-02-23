package group

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// API for GroupInvitation
	// ----------------------------------------------------
	// WRITE APIs.
	ADD_GROUPINVITATION               endpoints.Endpoint = "collaboration_addGroupInvitation"
	UPDATE_GROUPINVITATION            endpoints.Endpoint = "collaboration_updateGroupInvitation"
	DELETE_GROUPINVITATION            endpoints.Endpoint = "collaboration_deleteGroupInvitation"
	CLONE_GROUPINVITATION             endpoints.Endpoint = "collaboration_cloneGroupInvitation"
	CHANGE_GROUPINVITATION_OWNER      endpoints.Endpoint = "collaboration_changeGroupInvitationOwner"
	BULK_ADD_GROUPINVITATION          endpoints.Endpoint = "collaboration_bulkAddGroupInvitation"
	BULK_UPDATE_GROUPINVITATION       endpoints.Endpoint = "collaboration_bulkUpdateGroupInvitation"
	BULK_DELETE_GROUPINVITATION       endpoints.Endpoint = "collaboration_bulkDeleteGroupInvitation"
	BULK_CHANGE_GROUPINVITATION_OWNER endpoints.Endpoint = "collaboration_bulkChangeGroupInvitationOwner"

	// READ APIs.
	LIST_GROUPINVITATIONS     endpoints.Endpoint = "collaboration_listGroupInvitations"
	COUNT_GROUPINVITATIONS    endpoints.Endpoint = "collaboration_countGroupInvitations"
	GET_GROUPINVITATION_BY_ID endpoints.Endpoint = "ccollaboration_getGroupInvitationById"
)
