package rms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_LISTING          endpoints.Endpoint = "rms_addListing"
	UPDATE_LISTING       endpoints.Endpoint = "rms_updateListing"
	DELETE_LISTING       endpoints.Endpoint = "rms_deleteListing"
	CLONE_LISTING        endpoints.Endpoint = "rms_cloneListing"
	CHANGE_LISTING_OWNER endpoints.Endpoint = "rms_changeListingOwner"
	SET_LISTING_FAVORITE endpoints.Endpoint = "crm_setListingFavorite"

	// READ Operations.
	LIST_LISTINGS     endpoints.Endpoint = "rms_listListings"
	COUNT_LISTINGS    endpoints.Endpoint = "rms_countListings"
	GET_LISTING_BY_ID endpoints.Endpoint = "rms_getListingById"
)
