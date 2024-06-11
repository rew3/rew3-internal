package misc

/*
 A class to represent ''Phone''

 @field phoneType phone type information, from the lookup
 @field value phone number
 @field isPrimary determines if the the phone is primary or secondary

 @author rew3 on 2023/05/11

*/

type Phone struct {
	PhoneType string `json:"phone_type,omitempty" bson:"phone_type,omitempty"`
	Value     string `json:"value,omitempty" bson:"value,omitempty"`
	IsPrimary int    `json:"is_primary,omitempty" bson:"is_primary,omitempty"`
}
