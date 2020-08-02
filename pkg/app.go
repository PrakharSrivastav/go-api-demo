package pkg

import (
	"github.com/PrakharSrivastav/go-api-demo/pkg/api/characters"
	"github.com/PrakharSrivastav/go-api-demo/pkg/api/episodes"
	"github.com/PrakharSrivastav/go-api-demo/pkg/api/quotes"
	"github.com/PrakharSrivastav/go-api-demo/pkg/clients/breakingbad"
	"github.com/PrakharSrivastav/go-api-demo/pkg/store"
	cr "github.com/PrakharSrivastav/go-api-demo/pkg/store/characters"
	er "github.com/PrakharSrivastav/go-api-demo/pkg/store/episodes"
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

	return application()
	// or work with mock implementation
	// return mockApplication()
}

func application() (*App, error) {
	db, err := sqlx.Open("sqlite3", "./db/demo.db")
	if err != nil {
		return nil, err
	}

	err = store.CreateSchema(db)
	if err != nil {
		return nil, err
	}

	cc := cr.NewSqliteRepo(db)
	qq := qr.NewSqliteRepo(db)
	ee := er.NewSqliteRepo(db)

	cl := breakingbad.NewHttpClient()

	return &App{
		character: characters.New(cc, cl),
		quote:     quotes.New(qq, cl),
		episode:   episodes.New(ee, cl),
	},nil
}

func mockApplication() (*App, error) {
	cc := cr.NewMockRepo()
	qq := qr.NewMockRepo()
	ee := er.NewMockRepo()
	cl := breakingbad.NewMockApiClient()

	return &App{
		character: characters.New(cc, cl),
		quote:     quotes.New(qq, cl),
		episode:   episodes.New(ee, cl),
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
