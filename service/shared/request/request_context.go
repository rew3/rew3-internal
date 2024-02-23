package request

import (
	"math/big"

	a "github.com/rew3/rew3-internal/app/account"
	ac "github.com/rew3/rew3-internal/app/account/constants"
	c "github.com/rew3/rew3-internal/app/common"
	cc "github.com/rew3/rew3-internal/app/common/constants"
)

/*
 Model created to represent information about the targeted shared partner, for cases where:
     -> An agent shares an entity with ORGANIZATION
     -> ORGANIZATION shares an entity with TEAM/Agent
     -> TEAM/ORGANIZATION shares the record to Network

  @field MemberId id of the TEAM or ORGANIZATION
  @field Users ids of the INTERNAL(in case of ORGANIZATION) or EXTERNAL USER(in case of TEAM), not a required field
  @field AccessType if accessType = ALL, all the users of the organization/team have access to it, else visibility will be limited to .users field
  @field AccountTypeAlias TEAM or ORGANIZATION
*/

//type MetaInfoModel interface{}

// Deprecated.
/*type GlobalSharedMeta struct {
	MemberId         string               `json:"member_id,omitempty" bson:"member_id,omitempty"`
	Users            []string             `json:"users,omitempty" bson:"users,omitempty"`
	AccessType       SharedMetaAccessType `json:"access_type,omitempty" bson:"access_type,omitempty"`
	AccountTypeAlias AccountTypeAlias     `json:"account_type_alias,omitempty" bson:"account_type_alias,omitempty"`
	SharedToNetwork  bool                 `json:"shared_to_network,omitempty" bson:"shared_to_network,omitempty"`
}*/

type Rew3UserPersonalAlias struct {
	FirstName  string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty" bson:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	FullName   string `json:"full_name,omitempty" bson:"full_name,omitempty"`
	Salutation string `json:"salutation,omitempty" bson:"salutation,omitempty"`
	Gender     string `json:"gender,omitempty" bson:"gender,omitempty"`
}

type CompanyAlias struct {
	Id            string    `json:"_id" bson:"_id"`
	NoOfEmployees string    `json:"no_of_employees,omitempty" bson:"no_of_employees,omitempty"`
	Website       string    `json:"website,omitempty" bson:"website,omitempty"`
	AnnualRevenue big.Float `json:"annual_revenue,omitempty" bson:"annual_revenue,omitempty"`
	Industry      string    `json:"industry,omitempty" bson:"industry,omitempty"`
	Name          string    `json:"name,omitempty" bson:"name,omitempty"`
}

type UserInfo struct {
	Email       []c.Email             `json:"email,omitempty" bson:"email,omitempty"`
	PhoneNumber []c.Phone             `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
	Personal    Rew3UserPersonalAlias `json:"personal,omitempty" bson:"personal,omitempty"`
	Company     CompanyAlias          `json:"company,omitempty" bson:"company,omitempty"`
}

type TeamMiniAlias struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

/**
* The model represents the information extracted from a defined sharing rule involving a user
* @field entity entity defined in sharing rule
* @field visibleUsers the users whose data is visible to the current user. It is calculated based on the Role assigned
 */

type SharingRuleContext struct {
	Entity       cc.Entity `json:"entity,omitempty" bson:"entity,omitempty"`
	VisibleUsers []string  `json:"visible_users,omitempty" bson:"visible_users,omitempty"`
}

type TeamActions struct {
	PostSignUpCompleted bool `json:"post_sign_up_completed, omitempty" bson:"post_sign_up_completed, omitempty"`
}

/**
* The model represents status of the actions that are specific to a ORGANIZATION account
 */
type OrganizationActions struct{}

/**
* Model used to represent different user actions in different account_types
* @field teamActions only defined if account_type = TEAM
 */
type UserActions struct {
	TeamActions TeamActions `json:"team_actions,omitempty" bson:"team_actions,omitempty"`
}

/*
This class represents the [[RequestContext)]]

@field member member information of the user in the current context
@field user user in the current context
@field fullName full name of the user
@field memberId unique identifier of the member
@field lang the language used in the current context
@field command command in the current context
@field eTag etag of the entity in the current context

@author rew3 on 2023/05/15
*/
type RequestContext struct {
	Member           string               `json:"first_name,omitempty" bson:"first_name,omitempty"`
	User             a.MiniUser           `json:"mini_user,omitempty" bson:"mini_user,omitempty"`
	FullName         string               `json:"full_name,omitempty" bson:"full_name,omitempty"`
	Lang             string               `bson:"lang" i18n:"lang,default=en"` //Default is en. Later we need to unmarshal this using i18n.NewBundle(nil)
	Command          string               `json:"command,omitempty" bson:"command,omitempty"`
	ETag             string               `json:"e_tag,omitempty" bson:"e_tag,omitempty"`
	Entity           string               `json:"entity,omitempty" bson:"entity,omitempty"`
	Module           string               `json:"module,omitempty" bson:"module,omitempty"`
	IsExternal       bool                 `json:"is_external,omitempty" bson:"is_external,omitempty"`
	IsAdmin          bool                 `json:"is_admin,omitempty" bson:"is_admin,omitempty"`
	Info             UserInfo             `json:"user_info,omitempty" bson:"user_info,omitempty"`
	Timezone         string               `json:"timezone,omitempty" bson:"timezone,omitempty"`                   // values must be are TZInfo identifiers for eg: Europe/Prague, // America/Guatemala for more info: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	Teams            []TeamMiniAlias      `json:"teams,omitempty" bson:"teams,omitempty"`                         // list of teams to which user of the current context belongs to
	SubordinateUsers []string             `json:"subordinate_users,omitempty" bson:"subordinate_users,omitempty"` // id of the users who lie below the hierarchy of the current user
	RulesContext     []SharingRuleContext `json:"rules_context,omitempty" bson:"rules_context,omitempty"`
	//GlobalSharedMeta GlobalSharedMeta     `json:"global_shared_meta,omitempty" bson:"global_shared_meta,omitempty"` // this should be populated from the application layer, based on this, shared meta will be constructed
	AccountType  ac.AccountTypeAlias  `json:"account_type,omitempty" bson:"account_type,omitempty"`
	SecurityType ac.SecurityTypeAlias `json:"security_type,omitempty" bson:"security_type,omitempty"`
	Actions      UserActions          `json:"actions,omitempty" bson:"actions,omitempty"`
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
