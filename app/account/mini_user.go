package account

/*
 A class to represent an ''Mini User''

 @param _id unique identifier of the User
 @param firstName first name of the User
 @param lastName last name of the User

 @author rew3 on 2023/05/10
 @version 1.0.0

*/

type MiniUser struct {
	Id          string `json:"_id" bson:"_id"`
	FirstName   string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Title       string `json:"title,omitempty" bson:"title,omitempty"`
	Userpic     string `json:"userpic,omitempty" bson:"userpic,omitempty"`
	RoleName    string `json:"role_name,omitempty" bson:"role_name,omitempty"`
	JobTitle    string `json:"job_title,omitempty" bson:"job_title,omitempty"`
	CompanyName string `json:"company_name,omitempty" bson:"company_name,omitempty"`
}
