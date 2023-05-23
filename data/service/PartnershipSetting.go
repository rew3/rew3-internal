package service

import (
	. "github.com/rew3/rew3-base/data/service/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PartnershipSetting struct {
	ID                primitive.ObjectID `bson:"_id"`
	Invitee           string             `bson:"invitee,omitempty"`
	InvitedBy         UserType           `bson:"invited_by,omitempty"`
	PartnershipStatus PartnershipStatus  `bson:"partnership_status,omitempty"`
	Meta              MetaInfo           `bson:"meta,omitempty"`
}
