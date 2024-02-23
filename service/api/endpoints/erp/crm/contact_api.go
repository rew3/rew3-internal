package crm

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE Operations.
	ADD_CONTACT          endpoints.Endpoint = "crm_addContact"
	UPDATE_CONTACT       endpoints.Endpoint = "crm_updateContact"
	DELETE_CONTACT       endpoints.Endpoint = "crm_deleteContact"
	CLONE_CONTACT        endpoints.Endpoint = "crm_cloneContact"
	CHANGE_CONTACT_OWNER endpoints.Endpoint = "crm_changeContactOwner"
	SET_CONTACT_FAVORITE endpoints.Endpoint = "crm_setContactFavorite"

	// READ Operations.
	LIST_CONTACTS     endpoints.Endpoint = "crm_listContacts"
	COUNT_CONTACTS    endpoints.Endpoint = "crm_countContacts"
	GET_CONTACT_BY_ID endpoints.Endpoint = "crm_getContactById"
)
