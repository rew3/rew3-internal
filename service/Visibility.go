package service

import (
	"time"

	. "github.com/rew3/rew3-base/common/types"
	. "github.com/rew3/rew3-base/service/constants"
)

type TeamMini struct {
	ID   string `bson:"_id"`
	Name string `bson:"name,omitempty"`
}

type IndividualVisibility struct {
	IndividualId   string           `bson:"individual_id,omitempty"`
	PermissionType []PermissionType `bson:"permission_type,omitempty"`
}

type GroupVisibility struct {
	EntityId       string           `bson:"entity_id,omitempty"` // teamId, groupId etc
	PermissionType []PermissionType `bson:"permission_type,omitempty"`
}

type TeamVisibility struct {
	Team              TeamMini  `bson:"team,omitempty"`
	ActionPerformedBy MiniUser  `bson:"action_performed_by,omitempty"`
	ActionPerformedAt time.Time `bson:"action_performed_at,omitempty"`
}

type RecordVisibility struct {
	VisibilityType VisibilityType `bson:"visibility_type,omitempty"`
	//  user or contacId, provided when visibility type is individual. Doing this way, we dont
	// have to maintain ACL or RecordLevelSharing Data
	// In case team, we can fetch the team info to see if the current context user is in the
	// same team as contact
	IndividualVisibility []MiniUser           `bson:"individual_visibility,omitempty"`
	TeamVisibility       []TeamMini           `bson:"team_visibility,omitempty"`
	PartnershipSetting   []PartnershipSetting `bson:"partnership_setting,omitempty"`
}
