package customization

import "github.com/rew3/rew3-internal/service/grpc/api"

const (

	// API for Entity Relationship.
	// ----------------------------------------------------
	// WRITE APIs.
	LINK_ENTITY   api.APIOperation = "erm_linkEntity"
	UNLINK_ENTITY api.APIOperation = "erm_unlinkEntity"

	// READ APIs.
	LIST_LINKED_ENTITIES  api.APIOperation = "erm_listLinkedEntities"
)
