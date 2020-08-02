package pkg

import (
	"github.com/PrakharSrivastav/go-api-demo/pkg/api/characters"
	"github.com/PrakharSrivastav/go-api-demo/pkg/api/episodes"
	"github.com/PrakharSrivastav/go-api-demo/pkg/api/quotes"
	"github.com/PrakharSrivastav/go-api-demo/pkg/clients/breakingbad"
	cr "github.com/PrakharSrivastav/go-api-demo/pkg/store/characters"
	sr "github.com/PrakharSrivastav/go-api-demo/pkg/store/episodes"
	qr "github.com/PrakharSrivastav/go-api-demo/pkg/store/quotes"
	"github.com/go-chi/chi"
)

type App struct {
	character *characters.API
	quote     *quotes.API
	episode   *episodes.API
}

func New() *App {
	// could be setup in another function and can even be in a struct
	charRepo := cr.NewMockRepo()
	episRepo := sr.NewMockRepo()
	quotRepo := qr.NewMockRepo()
	client := breakingbad.NewMockApiClient()
	return &App{
		character: characters.New(charRepo, client),
		quote:     quotes.New(quotRepo, client),
		episode:   episodes.New(episRepo, client),
	}
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
