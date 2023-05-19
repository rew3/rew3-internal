package crm

/*
   Class to represent a [[CaseMini]] . This class contains few important fields of [[Case]]
   This class is also used in attach and detach operation related to ''Case''

   @field _id      unique identifier of Case
   @field subject  subject of case
   @field title    title of case

   @author rew3 on 2023/05/18
*/

type CaseMini struct {
	ID      string `bson:"_id"`
	Subject string `bson:"subject,omitempty"`
	Title   string `bson:"title,omitempty"`
}
