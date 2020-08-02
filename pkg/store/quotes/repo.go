package quotes

import "context"

type Repository interface {
	Get(ctx context.Context, ID int) (*Quote, error)
	Find(ctx context.Context, s Search ) ([]Quote, error)
	Insert(ctx context.Context, ch *Quote) error
	Update(ctx context.Context, ID int, ch *Quote) error
}