package queries

import "context"

type Storage interface {
	Create(ctx context.Context, query Query) (string, error)
	Read(ctx context.Context, id string) (Query, error)
	Update(ctx context.Context, query Query) error
	Delete(ctx context.Context, is string) error
}
