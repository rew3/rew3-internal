package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Repository[Entity any] interface {
	/* Writes */
	Insert(ctx context.Context, data *Entity) (*Entity, error)
	Update(ctx context.Context, id string, data *Entity) (*Entity, error)
	UpdateDataOnly(ctx context.Context, id string, data *Entity) (*Entity, error)
	FindAndUpdate(ctx context.Context, selector bson.M, data *Entity) (bool, error)
	UpdateWithRawData(ctx context.Context, selector bson.M, update bson.D) (bool, error)
	UnsetFields(ctx context.Context, selector bson.M, fields []string, multipleDoc bool) (bool, error)
	Archive(ctx context.Context, id string) (bool, error)
	UnArchive(ctx context.Context, id string) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
	FindAndDelete(ctx context.Context, selector bson.M) (bool, error)
	BulkInsert(ctx context.Context, data []*Entity) (bool, error)
	BulkUpdate(ctx context.Context, data []*Entity) (bool, error)

	/* Reads */
	FindById(ctx context.Context) *Entity
	Find(ctx context.Context) []*Entity
	FindBySelector(ctx context.Context) []*Entity
	Count(ctx context.Context) int64
	CountBySelector(ctx context.Context) int64
	Aggregate(ctx context.Context) []*Entity
	AggregateWithLookupJoin(ctx context.Context) []*Entity

	/*Public Reads*/
	FindByIdPublic(ctx context.Context) *Entity
	FindPublic(ctx context.Context) []*Entity
	FindBySelectorPublic(ctx context.Context) []*Entity
	CountPublic(ctx context.Context) int64
	CountBySelectorPublic(ctx context.Context) int64
}
