package common

import (
	"encoding/json"

	. "github.com/rew3/rew3-base/app/common/constants"
)

type Reference struct {
	Data     json.RawMessage `bson:"data,omitempty"`
	Module   Module          `bson:"module"`
	Entity   Entity          `bson:"entity,omitempty"`
	EntityID string          `bson:"entity_id,omitempty"`
	Title    string          `bson:"title,omitempty"`
}