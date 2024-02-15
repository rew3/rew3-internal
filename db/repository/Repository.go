package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Repository[Entity any] interface {
	/* Writes */
	Insert(ctx context.Context, data *Entity) (*Entity, error)
	Update(ctx context.Context, id string, data *Entity) (*Entity, error)
	UpdateOnly(ctx context.Context, id string, data *Entity) (bool, error)
	FindAndUpdate(ctx context.Context, selector bson.D, data *Entity, updateMany bool) (bool, error)
	FindAndUpdateBson(ctx context.Context, selector bson.D, data bson.D, updateMany bool) (bool, error)
	UnsetFields(ctx context.Context, selector bson.D, fields []string, multipleDoc bool) (bool, error)
	Archive(ctx context.Context, id string) (bool, error)
	UnArchive(ctx context.Context, id string) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
	FindAndDelete(ctx context.Context, selector bson.D, multipleDoc bool) (bool, error)
	BulkInsert(ctx context.Context, data []*Entity) (bool, error)
	BulkUpdate(ctx context.Context, data map[string]*Entity) (bool, error)
	BulkDelete(ctx context.Context, ids []string) (bool, error)
	AppendToArrayField(ctx context.Context, docId string, arrFieldPath string, elementSelectorName string, elementSelectorValue string, elementUpdateValue bson.D) (bool, error)
	RemoveFromArrayField(ctx context.Context, docId string, arrFieldPath string, elementSelectorName string, elementSelectorValue string) (bool, error)

	/* Reads */
	FindById(ctx context.Context, id string) *Entity
	Find(ctx context.Context, filters bson.D, offset int64, limit int64, sort bson.D) []*Entity
	Count(ctx context.Context, filters bson.D) int64
	Aggregate(ctx context.Context, pipelines []bson.D) []*Entity

	/*Public Reads*/
	FindByIdPublic(ctx context.Context, id string) *Entity
	FindPublic(ctx context.Context, filters bson.D, offset int64, limit int64, sort bson.D) []*Entity
	CountPublic(ctx context.Context, filters bson.D) int64
}
