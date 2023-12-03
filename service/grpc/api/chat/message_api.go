package chat

import "github.com/rew3/rew3-internal/service/grpc/api"

const (
	// WRITE APIs.
	ADD_MESSAGE api.APIOperation = "message_addMessage"

	// READ APIs.
	LIST_MESSAGE      api.APIOperation = "message_listMessages"
	GET_MESSAGE_BY_ID api.APIOperation = "message_GetMessageById"
)
