package constants

type PartnershipStatus string

const (
	ACTIVE              PartnershipStatus = "ACTIVE"
	REVOKED             PartnershipStatus = "REVOKED"
	REQUESTED_TO_UNLINK PartnershipStatus = "REQUESTED_TO_UNLINK"
)
