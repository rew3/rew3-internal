package common

/*
 A class to represent ''Phone''

 @field phoneType phone type information, from the lookup
 @field value phone number
 @field isPrimary determines if the the phone is primary or secondary

 @author rew3 on 2023/05/11

*/

type Phone struct {
	PhoneType string `bson:"phone_type,omitempty"`
	value     string `bson:"value,omitempty"`
	IsPrimary int    `bson:"is_primary,omitempty"`
}
