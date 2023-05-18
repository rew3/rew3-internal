package email

import (
	. "github.com/rew3/rew3-base/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
  A class to represent ''Email''

 @param emailType email type information, from the lookup
 @param value the email address
 @param isPrimary determines if the email is primary

 @author rew3 on 2023/05/11
 @version 1.0.0
*/

type Email struct {
	Id          primitive.ObjectID `bson:"_id"`
	To          []string           `bson:"to,omitempty"`
	BodyHtml    string             `bson:"body_html,omitempty"`
	Body        string             `bson:"body,omitempty"`
	Subject     string             `bson:"subject,omitempty"`
	Attachments []string           `bson:"attachments,omitempty"`
	JobTitle    string             `bson:"job_title,omitempty"`
	CompanyName string             `bson:"company_name,omitempty"`
}

type Attachment struct {
	Id               primitive.ObjectID `bson:"_id"`
	Meta             string             `bson:"meta,omitempty"`
	Reference        Reference          `bson:"reference,omitempty"`
	FileName         string             `bson:"file_name,omitempty"`
	MimeType         string             `bson:"mime_type,omitempty"`
	Description      string             `bson:"description,omitempty"`
	Path             string             `bson:"path,omitempty"`
	IsFavourite      bool               `bson:"is_favourite,omitempty"`
	DocumentMetadata DocMetadata        `bson:"document_metadata,omitempty"`
}
