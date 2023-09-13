package types

import (
	"reflect"

	"github.com/rew3/rew3-internal/app/common/constants"
)

type TypeRegistry struct {
	registry map[constants.Entity]reflect.Type
}

func NewTypeRegistry() TypeRegistry {
	return TypeRegistry{registry: make(map[constants.Entity]reflect.Type)}
}

func (r *TypeRegistry) RegisterType(entity constants.Entity, tpe interface{}) {
	r.registry[entity] = reflect.TypeOf(tpe)
}
