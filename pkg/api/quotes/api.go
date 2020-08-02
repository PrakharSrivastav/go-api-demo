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
	repoQuotes  quotes.Repository
	clientApi breakingbad.ApiClient
}

func New(repo quotes.Repository, client breakingbad.ApiClient) *API {
	return &API{
		repoQuotes:  repo,
		clientApi: client,
	}
}

func (a *API) Routes(r chi.Router) {
	r.Get("/quotes", a.find)
	r.Get("/quotes/{id}", a.get)
	r.Put("/quotes/{id}", a.edit)
	log.Println("quote routes registered")
}

