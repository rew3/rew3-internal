package entity

/*
 * Type to represent a [[AccountMini]] . This class contains few important fields of [[Account]]
 * This class is used as member of various Classes like [[com.rew3.crm.ticket.Case]]
 * This class is also used in attach and detach operation related to ''Account''
 *
 * @field _id           unique identifier of the account
 * @field noOfEmployees no of employees of the account
 * @field website       website of the account
 *                      contains full url i.e. http://www.rew3.com and domain url i.e. www.rew3.com of website
 *                      Only domain url is displayed in UI but filtering is done on full url
 *                      User can provide either domain or full url in Account Form
 *                      So to support filtering and sorting in standard format and to support UI displayable url we use [[Rew3Url]]
 * @field annualRevenue annual revenue of account
 * @field industry      industry of account, from lookup
 * @field name          name of the account
 *
 * @author              Rew3 on 2023/05/15

 */

type AccountMini struct {
	ID            string  `bson:"_id,omitempty"`
	NoOfEmployees string  `bson:"no_of_employees,omitempty"`
	Website       Rew3Url `bson:"website,omitempty"`
	AnnualRevenue int     `bson:"annual_revenue,omitempty"` //It was BigDecimal
	Industry      string  `bson:"industry,omitempty"`
	Name          string  `bson:"name,omitempty"`
}
