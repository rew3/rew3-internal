package crm

type LeadMini struct {
	ID        string `json:"_id" bson:"_id"`
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Title     string `json:"title,omitempty" bson:"title,omitempty"`
}
