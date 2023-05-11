package model

/*
 A class to represent ''Phone''

 @param phoneType phone type information, from the lookup
 @param value phone number
 @param isPrimary determines if the the phone is primary or secondary

 @author rew3 on 2023/05/11
 @version 1.0.0
*/

type Phone struct {
	phoneType *string `bson:"phone_type,omitempty"`
	value     *string `bson:"value,omitempty"`
	isPrimary *int    `bson:"is_primary,omitempty"`
}
