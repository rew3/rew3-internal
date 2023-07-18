package common

import (
	"time"

	a "github.com/rew3/rew3-internal/app/account"
	ac "github.com/rew3/rew3-internal/app/account/constants"
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
	Version      int64               `json:"_version" bson:"_version"`
	Created      time.Time           `json:"_created,omitempty" bson:"_created,omitempty"`
	CreatedBy    a.MiniUser          `json:"_created_by,omitempty" bson:"_created_by,omitempty"`
	LastModified time.Time           `json:"_last_modified,omitempty" bson:"_last_modified,omitempty"`
	ModifiedBy   a.MiniUser          `json:"_modified_by,omitempty" bson:"_modified_by,omitempty"`
	Owner        a.MiniUser          `json:"_owner,omitempty" bson:"_owner,omitempty"`
	Member       string              `json:"_member,omitempty" bson:"_member,omitempty"`
	Master       string              `json:"_master,omitempty" bson:"_master,omitempty"`
	Deleted      time.Time           `json:"_deleted,omitempty" bson:"_deleted,omitempty"`
	DeletedBy    string              `json:"_deleted_by,omitempty" bson:"_deleted_by,omitempty"`
	ArchivedAt   time.Time           `json:"_archived_at,omitempty" bson:"_archived_at,omitempty"`
	ArchivedBy   string              `json:"_archived_by,omitempty" bson:"_archived_by,omitempty"`
	Entity       string              `json:"_entity,omitempty" bson:"_entity,omitempty"`
	Module       string              `json:"_module,omitempty" bson:"_module,omitempty"`
	AccountType  ac.AccountTypeAlias `json:"_account_type,omitempty" bson:"_account_type,omitempty"`
	//GlobalSharedMeta GlobalSharedMeta `json:"_global_shared_meta,omitempty" bson:"_global_shared_meta,omitempty"`
}
