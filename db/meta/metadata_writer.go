package meta

import (
	"time"

	ac "github.com/rew3/rew3-internal/app/common"
	mUtil "github.com/rew3/rew3-internal/db/utils"
	s "github.com/rew3/rew3-internal/service/common/request"
	"go.mongodb.org/mongo-driver/bson"
)

type MetaDataWriter struct{}

func (writer *MetaDataWriter) WriteNewMeta(docId string, doc bson.D, context *s.RequestContext) bson.D {
	mc := MetaContext{
		ContextUser: context.User,
		MemberId:    context.Member,
		Entity:      context.Entity,
		Module:      context.Module,
		AccountType: context.AccountType,
	}
	meta := writer.newMeta(mc)
	meta = mUtil.AddBsonDNewFieldValue(meta, "_sid", docId)
	doc = mUtil.AddBsonDNewFieldValue(doc, "meta", meta)
	return doc
}

func (writer *MetaDataWriter) WriteUpdateMeta(doc bson.D, context *s.RequestContext) (bson.D, bson.M) {
	cUser := context.User
	meta := bson.D{
		{Key: "meta", Value: bson.M{
			"_last_modified": time.Now(),
			"_modified_by": bson.M{
				"_id":         cUser.Id,
				"first_name": cUser.FirstName,
				"last_name":  cUser.LastName,
			},
		}},
	}
	version := bson.M{
		"$inc": bson.M{
			"meta._version": 1,
		},
	}
	doc = append(doc, meta...)
	return doc, version
}

func (writer *MetaDataWriter) WriteSoftDeleteMeta(doc bson.D, context *s.RequestContext) bson.D {
	cUser := context.User
	originalVersion := mUtil.GetBsonDFieldValueWithDefault[int64](doc, "meta._version", 1)
	mUtil.SetBsonDFieldValue(doc, "meta._version", originalVersion+1)
	mUtil.SetBsonDFieldValue(doc, "meta._deleted", time.Now())
	mUtil.SetBsonDFieldValue(doc, "meta._deleted_by", bson.M{"_id": cUser.Id, "first_name": cUser.FirstName, "last_name": cUser.LastName})
	return doc
}

func (writer *MetaDataWriter) newMeta(mc MetaContext) bson.D {
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
	if bsonMeta, err := mUtil.EntityToBsonD[ac.MetaInfo](&meta, false, true); err == nil {
		return bsonMeta
	}
	return bson.D{}
}
