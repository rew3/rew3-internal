package meta

import (
	s "github.com/rew3/rew3-base/app/service"
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
	// Append a meta field
	doc := bson.M(*data)
	doc["meta"] = writer.newMeta(mc)
	return doc
}

func (writer *MetaDataWriter) WriteUpdateMeta(data *bson.M, originalMeta *s.MetaInfo, context *s.RequestContext) bson.M {
	mc := MetaContext{
		ContextUser: context.User,
		MemberId:    context.Member,
		Entity:      context.Entity,
		Module:      context.Module,
		AccountType: context.AccountType,
	}
	// Update a meta field
	// Todo need to update existing meta.. not write new meta later. updateOriginalMeta.
	// Remove meta from data, and add updated original meta.
	doc := bson.M(*data)
	doc["meta"] = writer.updateMeta(mc)
	return doc
}

func (writer *MetaDataWriter) newMeta(mc MetaContext) bson.M {
	// To do write core meta.
	return bson.M{}
}

func (writer *MetaDataWriter) updateMeta(mc MetaContext) bson.M {
	// todo update core meta.
	return bson.M{}
}
