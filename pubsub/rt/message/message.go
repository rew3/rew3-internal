package message

import (
	"time"

	"github.com/google/uuid"
)

/**
 * Realtime messaging Message structure.
 */
type Message struct {
	MessageId string           `json:"message_id"`
	HistoryId int64            `json:"history_id"` // history id, in order or all received message.
	Payload   RTMessagePayload `json:"payload"`
	TimeStamp time.Time        `json:"timestamp"`
}

/**
 * Create New Message.
 */
func NewMessage(mType MessageType, data map[string]interface{}) Message {
	return Message{
		MessageId: uuid.NewString(),
		HistoryId: time.Now().UnixMilli(),
		Payload: RTMessagePayload{
			Type: mType,
			Data: data,
		},
		TimeStamp: time.Now(),
	}
}

/**
 * Message payload.
 */
type RTMessagePayload struct {
	Type MessageType            `json:"message_type"`
	Data map[string]interface{} `json:"data"`
}

/**
 * Type of Message.
 */
type MessageType string

/**
 * Event Name for categorizing messages.
 */
type EventName string
