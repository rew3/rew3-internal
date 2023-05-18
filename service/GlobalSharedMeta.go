package service

import (
	. "github.com/rew3/rew3-base/service/constants"
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

type MetaInfoModel interface{}

type GlobalSharedMeta struct {
	MemberId         string               `bson:"member_id,omitempty"`
	Users            []string             `bson:"users,omitempty"`
	AccessType       SharedMetaAccessType `bson:"access_type,omitempty"`
	AccountTypeAlias AccountTypeAlias     `bson:"account_type_alias,omitempty"`
	SharedToNetwork  bool                 `bson:"shared_to_network,omitempty"`
}

//extends MetaInfoModel
