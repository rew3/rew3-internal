package crm

/*
   Class to represent a [[MiniCampaign]] . This class contains few important fields of ''Campaign''
   This class is used as one of the member field of ''Opportunity Information''. See [[OpportunityInformation]]

   @field _id           unique identifier of campaign
   @field campaignName  campaign name

   @author rew3 2023/05/18
*/

type MiniCampaign struct {
	ID           string `json:"_id" bson:"_id"`
	CampaignName string `json:"campaign_name,omitempty" bson:"campaign_name,omitempty"`
}
