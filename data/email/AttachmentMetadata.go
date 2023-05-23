package email

type AttachmentMetadata struct {
	Title        string `bson:"title,omitempty"`
	Creator      string `bson:"creator,omitempty"`
	Subject      string `bson:"subject,omitempty"`
	Description  string `bson:"description,omitempty"`
	Publisher    string `bson:"publisher,omitempty"`
	Contributor  string `bson:"contributor,omitempty"`
	Date         string `bson:"date,omitempty"`
	DocumentType string `bson:"document_type,omitempty"`
	Format       string `bson:"format,omitempty"`
	Identifier   string `bson:"identifier,omitempty"`
	Source       string `bson:"source,omitempty"`
	Language     string `bson:"language,omitempty"`
	Relation     string `bson:"relation,omitempty"`
	Coverage     string `bson:"coverage,omitempty"`
	Rights       string `bson:"rights,omitempty"`
}
