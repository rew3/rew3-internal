package platform

import (
	. "github.com/rew3/rew3-base/platform/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PartnershipSetting struct {
	Id                primitive.ObjectID `bson:"_id"`
	INVITEE           string             `bson:"invitee,omitempty"`
	INVITEDBY         UserType           `bson:"invited_by,omitempty"`
	PARTNERSHIPSTATUS PartnershipStatus  `bson:"partnership_status,omitempty"`
	META              MetaInfo           `bson:"meta,omitempty"`
}
