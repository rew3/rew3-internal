package rms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_LISTING          api.APIOperation = "rms_addListing"
	UPDATE_LISTING       api.APIOperation = "rms_updateListing"
	DELETE_LISTING       api.APIOperation = "rms_deleteListing"
	CLONE_LISTING        api.APIOperation = "rms_cloneListing"
	CHANGE_LISTING_OWNER api.APIOperation = "rms_changeListingOwner"
	SET_LISTING_FAVORITE api.APIOperation = "crm_setListingFavorite"

	// READ Operations.
	LIST_LISTINGS     api.APIOperation = "rms_listListings"
	COUNT_LISTINGS    api.APIOperation = "rms_countListings"
	GET_LISTING_BY_ID api.APIOperation = "rms_getListingById"
)
