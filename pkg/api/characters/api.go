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
	r.Get("/characters", a.find)
	r.Get("/characters/{id}", a.get)
	r.Post("/characters", a.add)
	r.Put("/characters/{id}", a.edit)
	r.Delete("/characters/{id}", a.delete)
	log.Println("characters routes registered")
}
