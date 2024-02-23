package rms

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_OFFER          endpoints.Endpoint = "rms_addOffer"
	UPDATE_OFFER       endpoints.Endpoint = "rms_updateOffer"
	DELETE_OFFER       endpoints.Endpoint = "rms_deleteOffer"
	CLONE_OFFER        endpoints.Endpoint = "rms_cloneOffer"
	CHANGE_OFFER_OWNER endpoints.Endpoint = "rms_changeOfferOwner"
	SET_OFFER_FAVORITE endpoints.Endpoint = "crm_setOfferFavorite"

	// READ Operations.
	LIST_OFFERS     endpoints.Endpoint = "rms_listOffers"
	COUNT_OFFERS    endpoints.Endpoint = "rms_countOffers"
	GET_OFFER_BY_ID endpoints.Endpoint = "rms_getOfferById"
)
