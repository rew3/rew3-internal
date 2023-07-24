package crm

import (
	"github.com/rew3/rew3-internal/service/grpc/api"
)

const (
	// WRITE Operations.
	ADD_CONTACT                    api.APIOperation = "addContact"
	UPDATE_CONTACT                 api.APIOperation = "updateContact"
	DELETE_CONTACT                 api.APIOperation = "deleteContact"
	CLONE_CONTACT                  api.APIOperation = "cloneContact"
	CHANGE_CONTACT_OWNER           api.APIOperation = "changeOwner"
	CHANGE_CONTACT_FAVORITE_STATUS api.APIOperation = "changeFavoriteStatus"

	// READ Operations.
	LIST_CONTACTS     api.APIOperation = "listContacts"
	GET_CONTACT_BY_ID api.APIOperation = "getContactById"
)
