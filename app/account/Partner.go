package account

import (
	. "github.com/rew3/rew3-internal/app/account/constants"
)

type Partner struct {
	UserId      string          `json:"user_id" bson:"user_id"`
	Email       string          `json:"email,omitempty" bson:"email,omitempty"`
	UserType    UserType        `json:"user_type,omitempty" bson:"user_type,omitempty"`
	MemberId    string          `json:"member_id" bson:"member_id"`
	AccountType Rew3AccountType `json:"account_type" bson:"account_type"`
}
