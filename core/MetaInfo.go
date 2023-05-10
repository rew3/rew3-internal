package core

import (
	"time"
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
	_version          int64
	_created          time.Time
	_createdBy        MiniUser
	_lastModified     time.Time
	_modifiedBy       time.Time
	_owner            MiniUser
	_member           string
	_master           string
	_deleted          string
	_deletedBy        string
	_archivedAt       time.Time
	_archivedBy       string
	_entity           string
	_module           string
	_agentId          string // If this exists, entire meta should belongs to an Agent,the field is populated if, the user in current context is an Agent
	_accountType      AccountTypeAlias
	_globalSharedMeta GlobalSharedMeta
}
