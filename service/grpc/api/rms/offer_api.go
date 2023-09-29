package rms

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE Operations.
	ADD_OFFER          api.APIOperation = "rms_addOffer"
	UPDATE_OFFER       api.APIOperation = "rms_updateOffer"
	DELETE_OFFER       api.APIOperation = "rms_deleteOffer"
	CLONE_OFFER        api.APIOperation = "rms_cloneOffer"
	CHANGE_OFFER_OWNER api.APIOperation = "rms_changeOfferOwner"

	// READ Operations.
	LIST_OFFERS     api.APIOperation = "rms_listOffers"
	COUNT_OFFERS    api.APIOperation = "rms_countOffers"
	GET_OFFER_BY_ID api.APIOperation = "rms_getOfferById"
)
