package queries

import (
	"context"
	_ "github.com/lib/pq"
)

type Storage interface {
	Create(Query) error
	Read(ctx context.Context, id string) (Query, error)
	Update(ctx context.Context, query Query) error
	Delete(ctx context.Context, is string) error
}
