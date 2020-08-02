package episodes

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

func (r repo) Get(ctx context.Context, ID int) (*Episode, error) {
	e := Episode{}
	err := r.db.GetContext(ctx, &e, "select * from episodes where id = $1", ID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r repo) Find(ctx context.Context, s Search) ([]Episode, error) {
	var e []Episode
	err := r.db.Select(&e, "select * from characters")
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r repo) Insert(ctx context.Context, ch *Episode) error {
	q := "Insert into episodes (id,title,season,episode,air_date,series) values (:id,:title,:season,:episode,:air_date,:series)"
	_, err := r.db.NamedExecContext(ctx, q, ch)
	if err != nil {
		return err
	}
	return nil
}

func (r repo) Update(ctx context.Context, ID int, ch *Episode) error {
	q := `Update episodes set 
title=:title,
season=:season,
episode=:episode,
air_date=:air_date,
series=:series
where id = :id`

	ch.ID = ID
	tx := r.db.MustBegin()
	r.db.MustExecContext(ctx, q, ch)
	return tx.Commit()
}
