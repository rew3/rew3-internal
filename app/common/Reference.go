package common

import (
	"encoding/json"

	c "github.com/rew3/rew3-base/app/common/constants"
)

type Reference struct {
	Data     json.RawMessage `bson:"data,omitempty"`
	Module   c.Module          `bson:"module"`
	Entity   c.Entity          `bson:"entity,omitempty"`
	EntityID string          `bson:"entity_id,omitempty"`
	Title    string          `bson:"title,omitempty"`
}
