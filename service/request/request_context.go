package request

import (
	"math/big"

	a "github.com/rew3/rew3-internal/app/account"
	ac "github.com/rew3/rew3-internal/app/account/constants"
	c "github.com/rew3/rew3-internal/app/common"
	cc "github.com/rew3/rew3-internal/app/common/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	MemberId         string               `bson:"member_id,omitempty"`
	Users            []string             `bson:"users,omitempty"`
	AccessType       SharedMetaAccessType `bson:"access_type,omitempty"`
	AccountTypeAlias AccountTypeAlias     `bson:"account_type_alias,omitempty"`
	SharedToNetwork  bool                 `bson:"shared_to_network,omitempty"`
}*/

type Rew3UserPersonalAlias struct {
	FirstName  string `bson:"first_name,omitempty"`
	MiddleName string `bson:"middle_name,omitempty"`
	LastName   string `bson:"last_name,omitempty"`
	FullName   string `bson:"full_name,omitempty"`
	Salutation string `bson:"salutation,omitempty"`
	Gender     string `bson:"gender,omitempty"`
}

type CompanyAlias struct {
	Id            primitive.ObjectID `bson:"_id"`
	NoOfEmployees string             `bson:"no_of_employees,omitempty"`
	Website       string             `bson:"website,omitempty"`
	AnnualRevenue big.Float          `bson:"annual_revenue,omitempty"`
	Industry      string             `bson:"industry,omitempty"`
	Name          string             `bson:"name,omitempty"`
}

type UserInfo struct {
	Email       []c.Email             `bson:"email,omitempty"`
	PhoneNumber []c.Phone             `bson:"phone_number,omitempty"`
	Personal    Rew3UserPersonalAlias `bson:"personal,omitempty"`
	Company     CompanyAlias          `bson:"company,omitempty"`
}

type TeamMiniAlias struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
}

/**
* The model represents the information extracted from a defined sharing rule involving a user
* @field entity entity defined in sharing rule
* @field visibleUsers the users whose data is visible to the current user. It is calculated based on the Role assigned
 */

type SharingRuleContext struct {
	Entity       cc.Entity `bson:"entity,omitempty"`
	VisibleUsers []string  `bson:"visible_users,omitempty"`
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
* @field teamActions only defined if account_type = TEAM
 */
type UserActions struct {
	TeamActions TeamActions `bson:"team_actions,omitempty"`
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
	Member           string               `bson:"first_name,omitempty"`
	User             a.MiniUser           `bson:"mini_user,omitempty"`
	FullName         string               `bson:"full_name,omitempty"`
	Lang             string               `bson:"lang" i18n:"lang,default=en"` //Default is en. Later we need to unmarshal this using i18n.NewBundle(nil)
	Command          string               `bson:"command,omitempty"`
	ETag             string               `bson:"e_tag,omitempty"`
	Entity           string               `bson:"entity,omitempty"`
	Module           string               `bson:"module,omitempty"`
	IsExternal       bool                 `bson:"is_external,omitempty"`
	IsAdmin          bool                 `bson:"is_admin,omitempty"`
	Info             UserInfo             `bson:"user_info,omitempty"`
	Timezone         string               `bson:"timezone,omitempty"`          // values must be are TZInfo identifiers for eg: Europe/Prague, // America/Guatemala for more info: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	Teams            []TeamMiniAlias      `bson:"teams,omitempty"`             // list of teams to which user of the current context belongs to
	SubordinateUsers []string             `bson:"subordinate_users,omitempty"` // id of the users who lie below the hierarchy of the current user
	RulesContext     []SharingRuleContext `bson:"rules_context,omitempty"`
	//GlobalSharedMeta GlobalSharedMeta     `bson:"global_shared_meta,omitempty"` // this should be populated from the application layer, based on this, shared meta will be constructed
	AccountType  ac.AccountTypeAlias  `bson:"account_type,omitempty"`
	SecurityType ac.SecurityTypeAlias `bson:"security_type,omitempty"`
	Actions      UserActions          `bson:"actions,omitempty"`
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
