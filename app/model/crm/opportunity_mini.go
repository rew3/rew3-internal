package crm

type OpportunityMini struct {
	ID              string      `json:"_id,omitempty" bson:"_id,omitempty"`
	OpportunityName string      `json:"opportunity_name,omitempty" bson:"opportunity_name,omitempty"`
	Company         AccountMini `json:"company,omitempty" bson:"company,omitempty"`
}
