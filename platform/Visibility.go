package platform

import (
	"time"

	. "github.com/rew3/rew3-base/data"
	. "github.com/rew3/rew3-base/platform/constants"
)

type TeamMini struct {
	Id   string `bson:"_id"`
	name string `bson:"name,omitempty"`
}

type IndividualVisibility struct {
	INDIVIDUALID   string           `bson:"individual_id,omitempty"`
	PERMISSIONTYPE []PermissionType `bson:"permission_type,omitempty"`
}

type GroupVisibility struct {
	ENTITYID       string           `bson:"entity_id,omitempty"` // teamId, groupId etc
	PERMISSIONTYPE []PermissionType `bson:"permission_type,omitempty"`
}

type TeamVisibility struct {
	Team              TeamMini  `bson:"entity_id,omitempty"`
	ACTIONPERFORMEDBY MiniUser  `bson:"action_performed_by,omitempty"`
	ACTIONPERFORMEDAT time.Time `bson:"action_performed_at,omitempty"`
}

type RecordVisibility struct {
	visibilityType VisibilityType `bson:"entity_id,omitempty"`
	//  user or contacId, provided when visibility type is individual. Doing this way, we dont
	// have to maintain ACL or RecordLevelSharing Data
	// In case team, we can fetch the team info to see if the current context user is in the
	// same team as contact
	individualVisibility []MiniUser           `bson:"action_performed_by,omitempty"`
	TEAMVISIBILITY       []TeamMini           `bson:"team_visibility,omitempty"`
	partnership_setting  []PartnershipSetting `bson:"partnership_setting,omitempty"`
}
