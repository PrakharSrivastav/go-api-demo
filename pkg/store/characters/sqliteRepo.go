package characters

import (
	"context"
)

type repo struct {
	db string
}

func (r repo) Get(ctx context.Context, ID int) (*Character, error) {
	panic("implement me")
}

func (r repo) Find(ctx context.Context, s SearchCriteria) ([]Character, error) {
	panic("implement me")
}

func (r repo) Insert(ctx context.Context, ch *Character) error {
	panic("implement me")
}

func (r repo) Update(ctx context.Context, ID int, ch *Character) error {
	panic("implement me")
}

func (r repo) Delete(ctx context.Context, ID int) error {
	panic("implement me")
}

func NewSqliteRepo() Repository {
	return repo{}
}
