package account

import (
	. "github.com/rew3/rew3-base/app/service/constants"
)

type Partner struct {
	UserId      string          `bson:"user_id"`
	Email       string          `bson:"email,omitempty"`
	UserType    UserType        `bson:"user_type,omitempty"`
	MemberId    string          `bson:"member_id"`
	AccountType Rew3AccountType `bson:"account_type"`
}
