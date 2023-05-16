package constants

type AccessType string

const (
	PRIVATEOPEN  AccessType = "privateOpen"
	PUBLICCLOSE  AccessType = "publicClose"
	PRIVATECLOSE AccessType = "privateClose"
	PUBLICOPEN   AccessType = "publicOpen"
)
