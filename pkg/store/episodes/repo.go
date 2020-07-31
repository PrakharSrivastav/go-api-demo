package episodes

import "context"

type Repository interface {
	Get(ctx context.Context, ID int) (*Episode, error)
	Find(ctx context.Context, s Search ) ([]Episode, error)
	Insert(ctx context.Context, ch *Episode) error
	Update(ctx context.Context, ID int, ch *Episode) error
	Delete(ctx context.Context, ID int) error
}
