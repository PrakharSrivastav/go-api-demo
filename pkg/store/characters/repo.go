package characters

import "context"

type Repository interface {
	Get(ctx context.Context, ID int) (*Character, error)
	Find(ctx context.Context, s SearchCriteria) ([]Character, error)
	Insert(ctx context.Context, ch *Character) error
	Update(ctx context.Context, ID int, ch *Character) error
	Delete(ctx context.Context, ID int) error
}

