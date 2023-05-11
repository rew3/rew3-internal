package model

import (
	"math/big"

	"go.mongodb.org/mongo-driver/bson/primitive"
	. "rew3.com/app-core/conf"
)

type Rew3UserPersonalAlias struct {
	FirstName  *string `bson:"first_name,omitempty"`
	MiddleName *string `bson:"middle_name,omitempty"`
	LastName   *string `bson:"last_name,omitempty"`
	FullName   *string `bson:"full_name,omitempty"`
	Salutation *string `bson:"salutation,omitempty"`
	Gender     *string `bson:"gender,omitempty"`
}

type CompanyAlias struct {
	Id            primitive.ObjectID `bson:"_id"`
	NoOfEmployees *string            `bson:"noOfEmployees,omitempty"`
	Website       *string            `bson:"website,omitempty"`
	AnnualRevenue *big.Float         `bson:"annualRevenue,omitempty"`
	Industry      *string            `bson:"industry,omitempty"`
	Name          *string            `bson:"name,omitempty"`
}

type UserInfo struct {
	email       *[]Email               `bson:"email,omitempty"`
	phoneNumber *[]Phone               `bson:"phone_number,omitempty"`
	personal    *Rew3UserPersonalAlias `bson:"personal,omitempty"`
	company     *CompanyAlias          `bson:"company,omitempty"`
}

type TeamMiniAlias struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	name string             `bson:"name,omitempty"`
}

/**
* The model represents the information extracted from a defined sharing rule involving a user
* @param entity entity defined in sharing rule
* @param visibleUsers the users whose data is visible to the current user. It is calculated based on the Role assigned
 */

type SharingRuleContext struct {
	entity       Entity   `bson:"entity,omitempty"`
	visibleUsers []string `bson:"visible_users,omitempty"`
}

type TeamActions struct {
	PostSignUpCompleted bool `bson:"post_sign_up_completed, omitempty"`
}

/**
* The model represents status of the actions that are specific to a ORGANIZATION account
 */
type OrganizationActions struct{}

/**
* Model used to represent different user actions in different account_types
* @param teamActions only defined if account_type = TEAM
 */
type UserActions struct {
	TeamActions TeamActions `bson:"team_actions,omitempty"`
}

/**
* This class represents the [[RequestContext)]]
*
* @param member member information of the user in the current context
* @param user user in the current context
* @param fullName full name of the user
* @param memberId unique identifier of the member
* @param lang the language used in the current context
* @param command command in the current context
* @param eTag etag of the entity in the current context
*
* @author rew3 on 2016/08/01
* @version 1.2.2
 */
type RequestContext struct {
	Member     string   `bson:"first_name,omitempty"`
	User       MiniUser `bson:"mini_user,omitempty"`
	FullName   string   `bson:"full_name,omitempty"`
	Lang       string   `bson:"lang" i18n:"lang,default=en"` //Default is en. Later we need to unmarshal this using i18n.NewBundle(nil)
	Command    string   `bson:"command,omitempty"`
	ETag       string   `bson:"e_tag,omitempty"`
	Entity     string   `bson:"entity,omitempty"`
	Module     string   `bson:"module,omitempty"`
	IsExternal bool     `bson:"is_external,omitempty"`
	IsAdmin    bool     `bson:"is_admin,omitempty"`
	Info       UserInfo `bson:"user_info,omitempty"`
	Timezone   string   `bson:"timezone,omitempty"` // values must be are TZInfo identifiers for eg: Europe/Prague,
	// America/Guatemala for more info: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	Teams            []TeamMiniAlias      `bson:"teams,omitempty"`             // list of teams to which user of the current context belongs to
	SubordinateUsers []string             `bson:"subordinate_users,omitempty"` // id of the users who lie below the hierarchy of the current user
	RulesContext     []SharingRuleContext `bson:"rules_context,omitempty"`
	GlobalSharedMeta GlobalSharedMeta     `bson:"global_shared_meta,omitempty"` // this should be populated from the application layer, based on this, shared meta will be constructed
	AccountType      AccountTypeAlias     `bson:"account_type,omitempty"`
	securityType     SecurityTypeAlias    `bson:"security_type,omitempty"`
	actions          UserActions          `bson:"actions,omitempty"`
}

func (r *RequestContext) WithCommand(command string) *RequestContext {
	r.Command = command
	return r
}

func (r *RequestContext) WithETag(eTag string) *RequestContext {
	r.ETag = eTag
	return r
}

func (r *RequestContext) WithEntity(entity string) *RequestContext {
	r.Entity = entity
	return r
}

func (r *RequestContext) WithModule(module string) *RequestContext {
	r.Module = module
	return r
}

/**
* I think we don't need this
* TODO Remove this after accountservice implementation
func NewRequestContext(member string, user string) *RequestContext {
    return &RequestContext{
        Member:       member, //Default Value "5a7989d20e00002f009687ce"
        User:         user,
        FullName:     "Proprio Direct",
        AccountType:  "ORGANIZATION",
        SecurityType: "ADVANCE"
    }
}
**/
