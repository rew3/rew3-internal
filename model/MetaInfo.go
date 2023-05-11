package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
 Class to represent the ''Meta Information'' of an entity
 This class will be required by all model class because security related things are also handled by this class.

 @param _version version of the entity
 @param _created created date and time
 @param _createdBy user who created the entity.
 @param _lastModified last modified date
 @param _modifiedBy user who modified the entity the last time .
 @param _owner user who is the owner of the entity .
 @param _member id of the member i.e the organization to whom the users belong
 @param _master
 @param _deleted deleted date and time
 @param _deletedBy user who deleted the record.
 @param _archivedAt the record archived date.
 @param _archivedBy user who archived the record.
 @param _entity Entity
 @param _module Module
 @param _agentId Agent Id
 @param _accountType The type of logged in Account
 @param _globalSharedMeta Global Shared Meta Information

 @author rew3 on 2023/05/10
 @version 1.0.0
*/

type MetaInfo struct {
	Version          int64               `bson:"_version"`
	Created          *primitive.DateTime `bson:"_created,omitempty"`
	CreatedBy        *MiniUser           `bson:"_created_by,omitempty"`
	LastModified     *primitive.DateTime `bson:"_last_modified,omitempty"`
	ModifiedBy       *primitive.DateTime `bson:"_modified_by,omitempty"`
	Owner            *MiniUser           `bson:"_owner,omitempty"`
	Member           *string             `bson:"_member,omitempty"`
	Master           *string             `bson:"_master,omitempty"`
	Deleted          *string             `bson:"_deleted,omitempty"`
	DeletedBy        *string             `bson:"_deleted_by,omitempty"`
	ArchivedAt       *primitive.DateTime `bson:"_archived_at,omitempty"`
	ArchivedBy       *string             `bson:"_archived_by,omitempty"`
	Entity           *string             `bson:"_entity,omitempty"`
	Module           *string             `bson:"_module,omitempty"`
	AgentId          *string             `bson:"_agent_id,omitempty"` // If this exists, entire meta should belongs to an Agent,the field is populated if, the user in current context is an Agent
	AccountType      *AccountTypeAlias   `bson:"_account_type,omitempty"`
	GlobalSharedMeta *GlobalSharedMeta   `bson:"_global_shared_meta,omitempty"`
}
