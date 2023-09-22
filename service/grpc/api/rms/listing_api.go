package rms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_LISTING                    api.APIOperation = "addListing"
	UPDATE_LISTING                 api.APIOperation = "updateListing"
	DELETE_LISTING                 api.APIOperation = "deleteListing"
	CLONE_LISTING                  api.APIOperation = "cloneListing"
	CHANGE_LISTING_OWNER           api.APIOperation = "changeListingOwner"

	// READ Operations.
	LIST_LISTINGS     api.APIOperation = "listListings"
	COUNT_LISTINGS    api.APIOperation = "countListings"
	GET_LISTING_BY_ID api.APIOperation = "getListingById"
)
