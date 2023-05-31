package model

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
	State         string `bson:"state,omitempty"`
	Location      string `bson:"location,omitempty"`
	City          string `bson:"city,omitempty"`
	Country       string `bson:"country,omitempty"`
	Region        string `bson:"region,omitempty"`
	Sector        string `bson:"sector,omitempty"`
	AddressNumber string `bson:"address_number,omitempty"`
	Zip           string `bson:"zip,omitempty"`
	StreetNumber  string `bson:"street_number,omitempty"`
	StreetName    string `bson:"street_name,omitempty"`
	District      string `bson:"district,omitempty"`
	Zone          string `bson:"zone,omitempty"`
	Province      string `bson:"province,omitempty"`
	Office        string `bson:"office,omitempty"`
	Department    string `bson:"department,omitempty"`
	Type          string `bson:"type,omitempty"`
	IsPrimary     int    `bson:"is_primary,omitempty"`
	Apt           string `bson:"apt,omitempty"`
	MailingStreet string `bson:"mailing_street,omitempty"`
}
