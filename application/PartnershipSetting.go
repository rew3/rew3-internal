package application

import (
	. "github.com/rew3/app-core/application/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PartnershipSetting struct {
	Id                primitive.ObjectID `bson:"_id"`
	INVITEE           string             `bson:"invitee,omitempty"`
	INVITEDBY         UserType           `bson:"invited_by,omitempty"`
	PARTNERSHIPSTATUS PartnershipStatus  `bson:"partnership_status,omitempty"`
	META              MetaInfo           `bson:"meta,omitempty"`
}
