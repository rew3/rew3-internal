package customization

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (

	// API for Entity Relationship.
	// ----------------------------------------------------
	// WRITE APIs.
	LINK_ENTITY   endpoints.Endpoint = "erm_linkEntity"
	UNLINK_ENTITY endpoints.Endpoint = "erm_unlinkEntity"

	// READ APIs.
	LIST_LINKED_ENTITIES endpoints.Endpoint = "erm_listLinkedEntities"
)
