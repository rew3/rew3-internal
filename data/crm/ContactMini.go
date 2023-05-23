package crm

type ContactMini struct {
	ID        string `bson:"_id"`
	FirstName string `bson:"first_name,omitempty"`
	LastName  string `bson:"last_name,omitempty"`
	Title     string `bson:"title,omitempty"`
}
