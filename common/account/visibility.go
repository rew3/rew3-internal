package account

import (
	"time"

	ac "github.com/rew3/rew3-pkg/common/account/constants"
)

type TeamMini struct {
	ID   string `json:"_id" bson:"_id"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

type IndividualVisibility struct {
	IndividualId   string              `json:"individual_id,omitempty" bson:"individual_id,omitempty"`
	PermissionType []ac.PermissionType `json:"permission_type,omitempty" bson:"permission_type,omitempty"`
}

type GroupVisibility struct {
	EntityId       string              `json:"entity_id,omitempty" bson:"entity_id,omitempty"` // teamId, groupId etc
	PermissionType []ac.PermissionType `json:"permission_type,omitempty" bson:"permission_type,omitempty"`
}

type TeamVisibility struct {
	Team              TeamMini  `json:"team,omitempty" bson:"team,omitempty"`
	ActionPerformedBy MiniUser  `json:"action_performed_by,omitempty" bson:"action_performed_by,omitempty"`
	ActionPerformedAt time.Time `json:"action_performed_at,omitempty" bson:"action_performed_at,omitempty"`
}

type RecordVisibility struct {
	VisibilityType ac.VisibilityType `json:"visibility_type,omitempty" bson:"visibility_type,omitempty"`
	//  user or contacId, provided when visibility type is individual. Doing this way, we dont
	// have to maintain ACL or RecordLevelSharing Data
	// In case team, we can fetch the team info to see if the current context user is in the
	// same team as contact
	IndividualVisibility []MiniUser `json:"individual_visibility,omitempty" bson:"individual_visibility,omitempty"`
	TeamVisibility       []TeamMini `json:"team_visibility,omitempty" bson:"team_visibility,omitempty"`
}
