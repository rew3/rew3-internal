package misc

/*
Type to represent ”Address”

@field state sate information, from the lookup
@field location location information, from the lookup
@field city city information, from the lookup
@field country country information, form the lookup
@field addressNumber address number
@field zip zip code
@field streetNumber street number
@field district name of the district
@field zone name of the zone
@field province name of the province
@field office name of the office
@field department name of the department
@field `type` type of the address
@field isPrimary determined if the address is primary secondary or primary
@field apt name of the apartment

@author rew3 on 2023/05/15
*/
type Address struct {
	State         string `json:"state,omitempty" bson:"state,omitempty"`
	Location      string `json:"location,omitempty" bson:"location,omitempty"`
	City          string `json:"city,omitempty" bson:"city,omitempty"`
	Country       string `json:"country,omitempty" bson:"country,omitempty"`
	Region        string `json:"region,omitempty" bson:"region,omitempty"`
	Sector        string `json:"sector,omitempty" bson:"sector,omitempty"`
	AddressNumber string `json:"address_number,omitempty" bson:"address_number,omitempty"`
	Zip           string `json:"zip,omitempty" bson:"zip,omitempty"`
	StreetNumber  string `json:"street_number,omitempty" bson:"street_number,omitempty"`
	StreetName    string `json:"street_name,omitempty" bson:"street_name,omitempty"`
	District      string `json:"district,omitempty" bson:"district,omitempty"`
	Zone          string `json:"zone,omitempty" bson:"zone,omitempty"`
	Province      string `json:"province,omitempty" bson:"province,omitempty"`
	Office        string `json:"office,omitempty" bson:"office,omitempty"`
	Department    string `json:"department,omitempty" bson:"department,omitempty"`
	Type          string `json:"type,omitempty" bson:"type,omitempty"`
	IsPrimary     int    `json:"is_primary,omitempty" bson:"is_primary,omitempty"`
	Apt           string `json:"apt,omitempty" bson:"apt,omitempty"`
	MailingStreet string `json:"mailing_street,omitempty" bson:"mailing_street,omitempty"`
}
