package crm

/*
   Class to represent a [[MiniProperty]] . This class contains few important fields of ''Property''
   This class is used as one of the member field of ''Opportunity Information''. See [[OpportunityInformation]]

   @field _id    unique identifier of property
   @field title  title of property

   @author rew3 2023/05/18
*/

type MiniProperty struct {
	ID    string `json:"_id" bson:"_id"`
	Title string `json:"title,omitempty" bson:"title,omitempty"`
}
