package consts

type ClosingStatusType string

const (
	CS_PRE_CLOSE         ClosingStatusType = "PRE_CLOSE"
	CS_PRE_CLOSED        ClosingStatusType = "CLOSED"
	CS_PRE_PARTIAL_CLOSE ClosingStatusType = "PARTIAL_CLOSE"
	CS_PRE_UN_CLOSE      ClosingStatusType = "UN_CLOSE"
	CS_PRE_RE_CLOSE      ClosingStatusType = "RE_CLOSE"
)
