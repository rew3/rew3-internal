package constants

type StatusType string

const (
	OK                    StatusType = "OK"
	CREATED               StatusType = "CREATED"
	DELETED               StatusType = "DELETED"
	ACCEPTED              StatusType = "ACCEPTED"
	BAD_REQUEST           StatusType = "BAD_REQUEST"
	UNAUTHORIZED          StatusType = "UNAUTHORIZED"
	FORBIDDEN             StatusType = "FORBIDDEN"
	INTERNAL_SERVER_ERROR StatusType = "INTERNAL_SERVER_ERROR"
	BAD_GATEWAY           StatusType = "BAD_GATEWAY"
	SERVICE_UNAVAILABLE   StatusType = "SERVICE_UNAVAILABLE"
	GATEWAY_TIMEOUT       StatusType = "GATEWAY_TIMEOUT"
	NOT_FOUND             StatusType = "NOT_FOUND"
)
