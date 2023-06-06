package module

import cs "github.com/rew3/rew3-internal/app/common/constants"

/* Variables */
var modulesToEntitiesMap = map[cs.Module][]cs.Entity{
	cs.CRM: {cs.Entity(cs.CRM_CONTACT), cs.Entity(cs.CRM_ACCOUNT), cs.Entity(cs.CRM_LEAD), cs.Entity(cs.CRM_OPPORTUNITY), cs.Entity(cs.CRM_CASE)},
	cs.RMS: {cs.Entity(cs.RMS_LISTING), cs.Entity(cs.RMS_REQUEST)},
}

type InvalidInputException struct {
	message string
}

func (e *InvalidInputException) Error() string {
	return e.message
}

func GetModule(entity cs.Entity) (cs.Module, error) {
	for module, entities := range modulesToEntitiesMap {
		for _, e := range entities {
			if e == entity {
				return module, nil
			}
		}
	}
	return cs.Module(0), &InvalidInputException{"Invalid entity."}
}
