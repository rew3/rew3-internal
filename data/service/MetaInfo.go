package service

import (
	"time"

	. "github.com/rew3/rew3-base/data/entity"
	. "github.com/rew3/rew3-base/data/service/constants"
)

/*
 Class to represent the ''Meta Information'' of an entity
 This class will be required by all model class because security related things are also handled by this class.

 @field _version version of the entity
 @field _created created date and time
 @field _createdBy user who created the entity.
 @field _lastModified last modified date
 @field _modifiedBy user who modified the entity the last time .
 @field _owner user who is the owner of the entity .
 @field _member id of the member i.e the organization to whom the users belong
 @field _master
 @field _deleted deleted date and time
 @field _deletedBy user who deleted the record.
 @field _archivedAt the record archived date.
 @field _archivedBy user who archived the record.
 @field _entity Entity
 @field _module Module
 @field _agentId Agent Id
 @field _accountType The type of logged in Account
 @field _globalSharedMeta Global Shared Meta Information

 @author rew3 on 2023/05/10

*/

type MetaInfo struct {
	Version          int64            `bson:"_version"`
	Created          time.Time        `bson:"_created,omitempty"`
	CreatedBy        MiniUser         `bson:"_created_by,omitempty"`
	LastModified     time.Time        `bson:"_last_modified,omitempty"`
	ModifiedBy       time.Time        `bson:"_modified_by,omitempty"`
	Owner            MiniUser         `bson:"_owner,omitempty"`
	Member           string           `bson:"_member,omitempty"`
	Master           string           `bson:"_master,omitempty"`
	Deleted          string           `bson:"_deleted,omitempty"`
	DeletedBy        string           `bson:"_deleted_by,omitempty"`
	ArchivedAt       time.Time        `bson:"_archived_at,omitempty"`
	ArchivedBy       string           `bson:"_archived_by,omitempty"`
	Entity           string           `bson:"_entity,omitempty"`
	Module           string           `bson:"_module,omitempty"`
	AgentId          string           `bson:"_agent_id,omitempty"` // If this exists, entire meta should belongs to an Agent,the field is populated if, the user in current context is an Agent
	AccountType      AccountTypeAlias `bson:"_account_type,omitempty"`
	GlobalSharedMeta GlobalSharedMeta `bson:"_global_shared_meta,omitempty"`
}
