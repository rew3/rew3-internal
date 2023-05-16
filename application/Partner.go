package application

import (
	. "github.com/rew3/app-core/application/constants"
)

type Partner struct {
	USERID      string          `bson:"user_id"`
	EMAIL       string          `bson:"email,omitempty"`
	USERTYPE    UserType        `bson:"user_type,omitempty"`
	MEMBERID    string          `bson:"member_id"`
	AccountType Rew3AccountType `bson:"account_type"`
}
