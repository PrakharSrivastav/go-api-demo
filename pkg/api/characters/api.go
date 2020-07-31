package characters

import (
	"github.com/PrakharSrivastav/go-api-demo/pkg/clients/breakingbad"
	"github.com/PrakharSrivastav/go-api-demo/pkg/store/characters"
	"github.com/go-chi/chi"
	"log"
)

// API represents the handler for characters
type API struct {
	// I prefer to keep these as interfaces
	// so as to have some room for mocking to test complicated scenarios
	repoChar  characters.Repository
	clientApi breakingbad.ApiClient
}

func New(repo characters.Repository, client breakingbad.ApiClient) *API {
	return &API{
		repoChar:  repo,
		clientApi: client,
	}
}

func (a *API) Routes(r chi.Router) {
	r.Get("/characters", nil)
	r.Get("/characters/:id", nil)
	r.Post("/characters", nil)
	r.Put("/characters/:id", nil)
	r.Delete("/characters/:id", nil)
	log.Println("characters routes registered")
}
