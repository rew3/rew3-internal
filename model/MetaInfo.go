package model

import (
	"rew3.com/app-core/model/utility"
)

/**
 * Class to represent the ''Meta Information'' of an entity
 * This class will be required by all model class because security related things are also handled by this class.
 *
 * @param _version version of the entity
 * @param _created created date and time
 * @param _createdBy user who created the entity.
 * @param _lastModified last modified date
 * @param _modifiedBy user who modified the entity the last time .
 * @param _owner user who is the owner of the entity .
 * @param _member id of the member i.e the organization to whom the users belong
 * @param _master
 * @param _deleted deleted date and time
 * @param _deletedBy user who deleted the record.
 * @param _archivedAt the record archived date.
 * @param _archivedBy user who archived the record.
 * @param _entity Entity
 * @param _module Module
 * @param _agentId Agent Id
 * @param _accountType The type of logged in Account
 * @param _globalSharedMeta Global Shared Meta Information
 *
 * @author rew3 on 2023/05/10
 * @version 1.0.0
 */

type MetaInfo struct {
	Version          int64
	Created          utility.DateTime
	CreatedBy        MiniUser
	LastModified     utility.DateTime
	ModifiedBy       utility.DateTime
	Owner            MiniUser
	Member           string
	Master           string
	Deleted          string
	DeletedBy        string
	ArchivedAt       utility.DateTime
	ArchivedBy       string
	Entity           string
	Module           string
	AgentId          string // If this exists, entire meta should belongs to an Agent,the field is populated if, the user in current context is an Agent
	AccountType      AccountTypeAlias
	AlobalSharedMeta GlobalSharedMeta
}
