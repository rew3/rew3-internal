package meta

import (
	"time"

	ac "github.com/rew3/rew3-internal/app/common"
	mUtil "github.com/rew3/rew3-internal/db/utils"
	s "github.com/rew3/rew3-internal/service/request"
	"go.mongodb.org/mongo-driver/bson"
)

type MetaDataWriter struct{}

func (writer *MetaDataWriter) WriteNewMeta(data *bson.M, context *s.RequestContext) bson.M {
	mc := MetaContext{
		ContextUser: context.User,
		MemberId:    context.Member,
		Entity:      context.Entity,
		Module:      context.Module,
		AccountType: context.AccountType,
	}
	// Write Append a meta field
	doc := bson.M(*data)
	doc["meta"] = writer.newMeta(mc)
	return doc
}

func (writer *MetaDataWriter) WriteUpdateMeta(data *bson.M, context *s.RequestContext) bson.M {
	cUser := context.User
	// Write Update a meta field
	doc := bson.M(*data)

	originalVersion := doc["meta._version"].(int)
	newVersion := originalVersion + 1
	doc["meta._version"] = newVersion
	doc["meta._last_modified"] = time.Now()
	doc["meta._modified_by"] = bson.M{"_id": cUser.Id, "first_name": cUser.FirstName, "last_name": cUser.LastName}
	return doc
}

func (writer *MetaDataWriter) WriteSoftDeleteMeta(data *bson.M, context *s.RequestContext) bson.M {
	cUser := context.User
	// Write Delete meta field
	doc := bson.M(*data)
	originalVersion := doc["meta._version"].(int)
	newVersion := originalVersion + 1
	doc["meta._version"] = newVersion
	doc["meta._deleted"] = time.Now()
	doc["meta._deleted_by"] = bson.M{"_id": cUser.Id, "first_name": cUser.FirstName, "last_name": cUser.LastName}
	return doc
}

func (writer *MetaDataWriter) newMeta(mc MetaContext) bson.M {
	meta := ac.MetaInfo{
		Version:      1,
		Created:      time.Now(),
		CreatedBy:    mc.ContextUser,
		LastModified: time.Now(),
		ModifiedBy:   mc.ContextUser,
		Owner:        mc.ContextUser,
		Member:       mc.MemberId,
		Entity:       mc.Entity,
		Module:       mc.Module,
		AccountType:  mc.AccountType,
	}
	if bsonMeta, err := mUtil.EntityToBson[ac.MetaInfo](&meta); err == nil {
		return bsonMeta
	}
	return bson.M{}
}
