package quotes

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func NewSqliteRepo(db *sqlx.DB) Repository {
	return repo{db: db}
}

func (r repo) Get(ctx context.Context, ID int) (*Quote, error) {
	e := Quote{}
	err := r.db.GetContext(ctx, &e, "select * from quotes where id = $1", ID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r repo) Find(ctx context.Context, s Search) ([]Quote, error) {
	var e []Quote
	err := r.db.Select(&e, "select * from quotes")
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r repo) Insert(ctx context.Context, ch *Quote) error {
	q := "Insert into episodes (id,quote,author,series) values (:id,:quote,:author,:series)"
	_, err := r.db.NamedExecContext(ctx, q, ch)
	if err != nil {
		return err
	}
	return nil
}

func (r repo) Update(ctx context.Context, ID int, ch *Quote) error {
	q := `Update episodes set 
quote=:quote,
author=:author,
series=:series
where id = :id`

	ch.ID = ID
	tx := r.db.MustBegin()
	r.db.MustExecContext(ctx, q, ch)
	return tx.Commit()
}
