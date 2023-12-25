package message

import (
	"time"

	"github.com/google/uuid"
)

/**
 * Message to be transmitted by message service.
 */
type Message struct {
	MessageId string           `json:"message_id"`
	TimeStamp time.Time        `json:"timestamp"`
	Payload   MessagePayload   `json:"payload"`
	Sender    MessageSender    `json:"sender"`
	Broadcast MessageBroadcast `json:"broadcast"`
}

/**
 * Create new message.
 */
func NewMessage(payload MessagePayload, sender MessageSender, broadcast MessageBroadcast) Message {
	return Message{
		MessageId: uuid.NewString(),
		TimeStamp: time.Now(),
		Payload:   payload,
		Sender:    sender,
		Broadcast: broadcast,
	}
}

/**
 * Message payload.
 * Contain message type and main message data.
 */
type MessagePayload struct {
	Type    MessageType            `json:"message_type"`
	Message map[string]interface{} `json:"message"`
}

/**
 * Message type. We can define multiple message type using this type that need to be exchanged over queue.
 */
type MessageType string

/**
 * Message sender.
 * Contains sender information, if sent by system, it will be Internal.
 */
type MessageSender struct {
	Internal bool   `json:"internal"`
	SenderId string `json:"sender_id"`
	MemberId string `json:"member_id"`
}

/**
 * Message broadcast. Message will be intended for receipent as provided in message broadcast.
 * If system, then message is used for system and internal use.
 */
type MessageBroadcast struct {
	Internal  bool     `json:"internal"`
	Users     []string `json:"users"`      // broadcast to provided users if present.
	OmitUsers []string `json:"omit_users"` // Dont broadcast to these users if present.
	MemberId  string   `json:"member_id"`  // broadcast to all user of this member id if users is empty.
}
