package model

/**
 * Model created to represent information about the targeted shared partner, for cases where:
 *    -> An agent shares an entity with ORGANIZATION
 *    -> ORGANIZATION shares an entity with TEAM/Agent
 *    -> TEAM/ORGANIZATION shares the record to Network
 *
 * @param memberId id of the TEAM or ORGANIZATION
 * @param users ids of the INTERNAL(in case of ORGANIZATION) or EXTERNAL USER(in case of TEAM), not a required field
 * @param accessType if accessType = ALL, all the users of the organization/team have access to it, else visibility will be limited to .users field
 * @param accountTypeAlias TEAM or ORGANIZATION
 */

type MetaInfoModel interface{}

type GlobalSharedMeta struct {
	MemberId         string
	Users            []string
	AccessType       SharedMetaAccessType
	AccountTypeAlias AccountTypeAlias
	SharedToNetwork  bool
}

//extends MetaInfoModel
