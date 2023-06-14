package common

import (
	"encoding/json"

	c "github.com/rew3/rew3-internal/app/common/constants"
)

type Reference struct {
	Data     json.RawMessage `json:"data,omitempty" bson:"data,omitempty"`
	Module   c.Module        `json:"module" bson:"module"`
	Entity   c.Entity        `json:"entity,omitempty" bson:"entity,omitempty"`
	EntityID string          `json:"entity_id,omitempty" bson:"entity_id,omitempty"`
	Title    string          `json:"title,omitempty" bson:"title,omitempty"`
}
