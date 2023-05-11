package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
 A class to represent an ''Mini User''

 @param _id unique identifier of the User
 @param firstName first name of the User
 @param lastName last name of the User

 @author rew3 on 2023/05/10
 @version 1.0.0

*/

type MiniUser struct {
	Id          primitive.ObjectID `bson:"_id"`
	FirstName   *string            `bson:"first_name,omitempty"`
	LastName    *string            `bson:"last_name,omitempty"`
	Title       *string            `bson:"title,omitempty"`
	Userpic     *string            `bson:"userpic,omitempty"`
	RoleName    *string            `bson:"role_name,omitempty"`
	JobTitle    *string            `bson:"job_title,omitempty"`
	CompanyName *string            `bson:"company_name,omitempty"`
}
