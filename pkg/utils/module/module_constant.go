package module

import (
	cs "github.com/rew3/rew3-internal/app/common/constants"
	baseResponse "github.com/rew3/rew3-internal/service/common/response/constants"
)

// A method for conversing constant (enum) to string value
var modules = []cs.Module{
	cs.CRM,
	cs.RMS,
}

var entities = []cs.Entity{
	cs.CRM_CONTACT,
	cs.CRM_ACCOUNT,
	cs.CRM_LEAD,
	cs.CRM_OPPORTUNITY,
	cs.CRM_CASE,
	cs.RMS_LISTING,
	cs.RMS_REQUEST,
}

var statuses = []baseResponse.StatusType{
	baseResponse.OK,
	baseResponse.CREATED,
	baseResponse.DELETED,
	baseResponse.ACCEPTED,
	baseResponse.BAD_REQUEST,
	baseResponse.UNAUTHORIZED,
	baseResponse.FORBIDDEN,
	baseResponse.INTERNAL_SERVER_ERROR,
	baseResponse.BAD_GATEWAY,
	baseResponse.SERVICE_UNAVAILABLE,
	baseResponse.GATEWAY_TIMEOUT,
	baseResponse.NOT_FOUND,
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
	case cs.Entity:
		for _, entity := range entities {
			if entity == c {
				return string(c), nil //try to return the value of constant.
			}
		}
	case cs.Module:
		for _, module := range modules {
			if module == c {
				return string(c), nil
			}
		}
	case baseResponse.StatusType:
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
