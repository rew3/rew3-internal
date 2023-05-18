package crm

/*
Class to represent ”Social Communications”

@field value value of the social communication
@field socialCommunicationType See [[com.rew3.platform.core.common.SocialCommunicationType]]
@field isPrimary determines the social communication type is primary or secondary

@author rew3 on 2023/05/18
*/
type SocialCommunications struct {
	Value                   string                  `bson:"value,omitempty"`
	SocialCommunicationType SocialCommunicationType `bson:"social_communication_type,omitempty"`
	IsPrimary               int                     `bson:"is_primary,omitempty"`
}
