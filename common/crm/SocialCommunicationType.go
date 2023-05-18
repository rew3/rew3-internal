package crm

/**
  Class to represent ''Social Communication Type''

  @field _id unique identifier of the social communication type
  @field title title of the social communication type

  @author rew3 on 2023/05/18
*/

type SocialCommunicationType struct {
	ID    string `bson:"_id,omitempty"`
	Title string `bson:"title,omitempty"`
}
