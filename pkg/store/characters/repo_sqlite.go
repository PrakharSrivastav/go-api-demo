package characters

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func (r repo) Get(ctx context.Context, ID int) (*Character, error) {
	e := Character{}
	err := r.db.GetContext(ctx, &e, "select * from characters where id = $1", ID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r repo) Find(ctx context.Context, s SearchCriteria) ([]Character, error) {
	var e []Character
	err := r.db.Select(&e, "select * from characters")
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r repo) Insert(ctx context.Context, ch *Character) error {
	q := "Insert into characters (id,name,birthday,img,status,nickname,portrayed) values (:id,:name,:birthday,:img,:status,:nickname,:portrayed)"
	_, err := r.db.NamedExecContext(ctx, q, ch)
	if err != nil {
		return err
	}
	return nil
}

func (r repo) Update(ctx context.Context, ID int, ch *Character) error {

	q := `Update characters set 
name=:name,
birthday=:birthday,
img=:img,
status=:status,
nickname=:nickname,
portrayed=:portrayed
where id = :id`

	ch.ID = ID
	tx := r.db.MustBegin()
	r.db.MustExecContext(ctx, q, ch)
	return tx.Commit()
}

func NewSqliteRepo(db *sqlx.DB) Repository {
	return repo{db: db}
}
