package repository

import (
	"context"
	"encoding/json"

	s "github.com/rew3/rew3-base/app/service"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository[Entity any] interface {
	/* Writes */
	Insert(ctx context.Context, data *Entity) (*Entity, error)
	Update(ctx context.Context, id string, data *Entity) (*Entity, error)
	UpdateDataOnly(ctx context.Context, id string, data *Entity) (*Entity, error)
	FindAndUpdate(ctx context.Context, selector json.RawMessage, data *Entity) (bool, error)
	UpdateWithRawData(ctx context.Context, selector json.RawMessage, update bson.D) (bool, error)
	UnsetFields(ctx context.Context, selector json.RawMessage, fields []string, multipleDoc bool) (bool, error)
	Archive(ctx context.Context, id string) (bool, error)
	UnArchive(ctx context.Context, id string) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
	FindAndDelete(ctx context.Context, selector json.RawMessage) (bool, error)
	BulkInsert(ctx context.Context, data []*Entity) (bool, error)
	BulkUpdate(ctx context.Context, data []*Entity) (bool, error)

	/* Reads */
	FindById(ctx context.Context, id string) *Entity
	Find(ctx context.Context, param s.RequestParam) []*Entity
	FindBySelector(ctx context.Context, selector json.RawMessage, offset int, limit int, sort json.RawMessage) []*Entity
	Count(ctx context.Context, param s.RequestParam) int64
	CountBySelector(ctx context.Context, selector json.RawMessage, offset int, limit int, sort json.RawMessage) int64
	Aggregate(ctx context.Context) []*Entity
	AggregateWithLookupJoin(ctx context.Context) []*Entity

	/*Public Reads*/
	FindByIdPublic(ctx context.Context, id string) *Entity
	FindPublic(ctx context.Context, param s.RequestParam) []*Entity
	FindBySelectorPublic(ctx context.Context, selector json.RawMessage, offset int, limit int, sort json.RawMessage) []*Entity
	CountPublic(ctx context.Context, param s.RequestParam) int64
	CountBySelectorPublic(ctx context.Context, selector json.RawMessage, offset int, limit int, sort json.RawMessage) int64
}