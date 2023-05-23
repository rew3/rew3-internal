package crm

/*
 A class to represent ''Email''

 @field EmailType email type information, from the lookup
 @field value email number
 @field isPrimary determines if the the email is primary or secondary

 @author rew3 on 2023/05/11

*/

type Email struct {
	EmailType string `bson:"email_type,omitempty"`
	value     string `bson:"value,omitempty"`
	IsPrimary int    `bson:"is_primary,omitempty"`
}
