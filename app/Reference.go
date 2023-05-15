package app

import (
	"encoding/json"

	. "github.com/rew3/core-base/app/_constants"
)

type Reference struct {
	Data     *json.RawMessage `bson:"data,omitempty"`
	Module   Module           `bson:"module"`
	Entity   *Entity          `bson:"entity,omitempty"`
	EntityID *string          `bson:"entity_id,omitempty"`
	Title    *string          `bson:"title,omitempty"`
}
