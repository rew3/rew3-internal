package model

type OpportunityMini struct {
	ID              string      `bson:"_id"`
	OpportunityName string      `bson:"opportunity_name,omitempty"`
	Company         AccountMini `bson:"company,omitempty"`
}
