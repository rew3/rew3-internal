package chat

import "github.com/rew3/rew3-internal/service/api/endpoints"

const (
	// WRITE APIs.
	ADD_MESSAGE endpoints.Endpoint = "message_addMessage"

	// READ APIs.
	LIST_MESSAGE      endpoints.Endpoint = "message_listMessages"
	GET_MESSAGE_BY_ID endpoints.Endpoint = "message_GetMessageById"
)
