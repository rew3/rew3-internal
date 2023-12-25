package message

import "time"

/**
 * Realtime messaging Message structure.
 */
type Message struct {
	MessageId string           `json:"message_id"`
	HistoryId string           `json:"history_id"` // history id, in order or all received message.
	Category  MessageCategory  `json:"category"`
	Payload   RTMessagePayload `json:"payload"`
	TimeStamp time.Time        `json:"timestamp"`
}

/**
 * Message payload interface.
 * Message must inherit this interface.
 */
type RTMessagePayload interface{}

/**
 * Message category type.
 * Multiple message type can be sent to each category.
 */
type MessageCategory string
