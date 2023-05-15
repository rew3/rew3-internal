package misc

import (
	. "github.com/rew3/core-base/app/_constants"
)

var modules = []Module{
	CRM,
	RMS,
}

var entities = []Entity{
	CRM_CONTACT,
	CRM_ACCOUNT,
	CRM_LEAD,
	CRM_OPPORTUNITY,
	CRM_CASE,
	RMS_LISTING,
	RMS_REQUEST,
}

var statuses = []StatusType{
	OK,
	CREATED,
	DELETED,
	ACCEPTED,
	BAD_REQUEST,
	UNAUTHORIZED,
	FORBIDDEN,
	INTERNAL_SERVER_ERROR,
	BAD_GATEWAY,
	SERVICE_UNAVAILABLE,
	GATEWAY_TIMEOUT,
	NOT_FOUND,
}

type InvalidException struct {
	message string
}

func (e *InvalidException) Error() string {
	return e.message
}

// GetConstantValue retrieves the value of any constant provided as a parameter. Example it returns response like CRM_CONTACT
// If the constant is not found, an empty string is returned.
func GetConstantValue(constant interface{}) (string, error) {
	switch c := constant.(type) {
	case Entity:
		for _, entity := range entities {
			if entity == c {
				return string(c), nil //try to return the value of constant.
			}
		}
	case Module:
		for _, module := range modules {
			if module == c {
				return string(c), nil
			}
		}
	case StatusType:
		for _, status := range statuses {
			if status == c {
				return string(c), nil
			}
		}
	default:
		return "", &InvalidException{"Invalid constant type."}
	}
	return "", &InvalidException{"Invalid constant value."}
}
