package crm

import (
	"github.com/rew3/rew3-internal/service/grpc/api"
)

const (
	// WRITE Operations.
	ADD_CONTACT                    api.APIOperation = "crm_addContact"
	UPDATE_CONTACT                 api.APIOperation = "crm_updateContact"
	DELETE_CONTACT                 api.APIOperation = "crm_deleteContact"
	CLONE_CONTACT                  api.APIOperation = "crm_cloneContact"
	CHANGE_CONTACT_OWNER           api.APIOperation = "crm_changeContactOwner"
	CHANGE_CONTACT_FAVORITE_STATUS api.APIOperation = "crm_changeContactFavoriteStatus"

	// READ Operations.
	LIST_CONTACTS     api.APIOperation = "crm_listContacts"
	COUNT_CONTACTS    api.APIOperation = "crm_countContacts"
	GET_CONTACT_BY_ID api.APIOperation = "crm_getContactById"
)
