package pkg

import (
	"github.com/PrakharSrivastav/go-api-demo/pkg/api/characters"
	"github.com/PrakharSrivastav/go-api-demo/pkg/api/episodes"
	"github.com/PrakharSrivastav/go-api-demo/pkg/api/quotes"
	"github.com/PrakharSrivastav/go-api-demo/pkg/clients/breakingbad"
	"github.com/PrakharSrivastav/go-api-demo/pkg/store"
	cr "github.com/PrakharSrivastav/go-api-demo/pkg/store/characters"
	sr "github.com/PrakharSrivastav/go-api-demo/pkg/store/episodes"
	qr "github.com/PrakharSrivastav/go-api-demo/pkg/store/quotes"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	character *characters.API
	quote     *quotes.API
	episode   *episodes.API
}

func New() (*App, error) {

	db, err := sqlx.Open("sqlite3", "./db/demo.db")
	if err != nil {
		return nil, err
	}

	err = store.CreateSchema(db)
	if err != nil {
		return nil, err
	}

	// could be setup in a struct
	charRepo := cr.NewSqliteRepo(db)
	episRepo := sr.NewSqliteRepo(db)
	quotRepo := qr.NewSqliteRepo(db)
	client := breakingbad.NewHttpClient()

	return &App{
		character: characters.New(charRepo, client),
		quote:     quotes.New(quotRepo, client),
		episode:   episodes.New(episRepo, client),
	}, nil
}

func (a *App) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		a.character.Routes(r)
		a.episode.Routes(r)
		a.quote.Routes(r)
	})
	return r
}
