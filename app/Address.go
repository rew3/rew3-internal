package app

/*
 * Tyoe to represent ''Address''
 *
 * @param state sate information, from the lookup
 * @param location location information, from the lookup
 * @param city city information, from the lookup
 * @param country country information, form the lookup
 * @param addressNumber address number
 * @param zip zip code
 * @param streetNumber street number
 * @param district name of the district
 * @param zone name of the zone
 * @param province name of the province
 * @param office name of the office
 * @param department name of the department
 * @param `type` type of the address
 * @param isPrimary determined if the address is primary secondary or primary
 * @param apt name of the apartment
 *
 * @author rew3 on 2023/05/15
 * @version 1.0.0
 */
type Address struct {
	State         *string `bson:"state,omitempty"`
	Location      *string `bson:"location,omitempty"`
	City          *string `bson:"city,omitempty"`
	Country       *string `bson:"country,omitempty"`
	Region        *string `bson:"region,omitempty"`
	Sector        *string `bson:"sector,omitempty"`
	AddressNumber *string `bson:"address_number,omitempty"`
	Zip           *string `bson:"zip,omitempty"`
	StreetNumber  *string `bson:"street_number,omitempty"`
	StreetName    *string `bson:"street_name,omitempty"`
	District      *string `bson:"district,omitempty"`
	Zone          *string `bson:"zone,omitempty"`
	Province      *string `bson:"province,omitempty"`
	Office        *string `bson:"office,omitempty"`
	Department    *string `bson:"department,omitempty"`
	Type          *string `bson:"type,omitempty"`
	IsPrimary     *int    `bson:"is_primary,omitempty"`
	Apt           *string `bson:"apt,omitempty"`
	MailingStreet *string `bson:"mailing_street,omitempty"`
}
