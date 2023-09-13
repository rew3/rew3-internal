package types

import (
	"fmt"
	"reflect"

	"github.com/rew3/rew3-internal/app/common/constants"
)

func DecodeType[T any](raw interface{}, entity constants.Entity, registry TypeRegistry) (T, error) {
	t, ok := registry.registry[entity]
	if !ok {
		var data T
		return data, fmt.Errorf("Type not found in registry: %s", entity)
	}
	if reflect.TypeOf(raw) != t {
		var data T
		return data, fmt.Errorf("Type mismatch for key: %s", entity)
	}
	return raw.(T), nil
}
