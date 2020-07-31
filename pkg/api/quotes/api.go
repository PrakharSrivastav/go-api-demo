package quotes

import (
	"github.com/PrakharSrivastav/go-api-demo/pkg/clients/breakingbad"
	"github.com/PrakharSrivastav/go-api-demo/pkg/store/quotes"
	"github.com/go-chi/chi"
	"log"
)

// API represents the handler for quotes
type API struct {
	// I prefer to keep these as interfaces
	// so as to have some room for mocking to test complicated scenarios
	repoChar  quotes.Repository
	clientApi breakingbad.ApiClient
}

func New(repo quotes.Repository, client breakingbad.ApiClient) *API {
	return &API{
		repoChar:  repo,
		clientApi: client,
	}
}

func (a *API) Routes(r chi.Router) {
	r.Get("/quotes", nil)
	r.Get("/quotes/:id", nil)
	r.Post("/quotes", nil)
	r.Put("/quotes/:id", nil)
	r.Delete("/quotes/:id", nil)
	log.Println("quote routes registered")
}

