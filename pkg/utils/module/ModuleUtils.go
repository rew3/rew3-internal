package service

import . "github.com/rew3/rew3-base/app/service/constants"

/* Variables */
var modules = map[Module][]Entity{
	CRM: {Entity(CRM_CONTACT), Entity(CRM_ACCOUNT), Entity(CRM_LEAD), Entity(CRM_OPPORTUNITY), Entity(CRM_CASE)},
	RMS: {Entity(RMS_LISTING), Entity(RMS_REQUEST)},
}

type InvalidInputException struct {
	message string
}

func (e *InvalidInputException) Error() string {
	return e.message
}

func getModule(entity Entity) (Module, error) {
	for module, entities := range modules {
		for _, e := range entities {
			if e == entity {
				return module, nil
			}
		}
	}
	return Module(0), &InvalidInputException{"Invalid entity."}
}
