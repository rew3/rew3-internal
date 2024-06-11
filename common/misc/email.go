package misc

/*
 A class to represent ''Email''

 @field EmailType email type information, from the lookup
 @field value email number
 @field isPrimary determines if the the email is primary or secondary

 @author rew3 on 2023/05/11

*/

type Email struct {
	EmailType string `json:"email_type,omitempty" bson:"email_type,omitempty"`
	Value     string `json:"value,omitempty" bson:"value,omitempty"`
	IsPrimary int    `json:"is_primary,omitempty" bson:"is_primary,omitempty"`
}
