package model

/*
 A class to represent ''Email''

 @param emailType email type information, from the lookup
 @param value email number
 @param isPrimary determines if the the email is primary or secondary

 @author rew3 on 2023/05/11
 @version 1.0.0
*/

type Email struct {
	emailType *string `bson:"phone_type,omitempty"`
	value     *string `bson:"value,omitempty"`
	isPrimary *int    `bson:"is_primary,omitempty"`
}
