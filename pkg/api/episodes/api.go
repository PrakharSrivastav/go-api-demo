package episodes

import (
	"github.com/PrakharSrivastav/go-api-demo/pkg/clients/breakingbad"
	"github.com/PrakharSrivastav/go-api-demo/pkg/store/episodes"
	"github.com/go-chi/chi"
	"log"
)

// API represents the handler for episodes
type API struct {
	// I prefer to keep these as interfaces
	// so as to have some room for mocking to test complicated scenarios
	repoChar  episodes.Repository
	clientApi breakingbad.ApiClient
}

func New(repo episodes.Repository, client breakingbad.ApiClient) *API {
	return &API{
		repoChar:  repo,
		clientApi: client,
	}
}

func (a *API) Routes(r chi.Router) {
	r.Get("/episodes", nil)
	r.Get("/episodes/:id", nil)
	r.Post("/episodes", nil)
	r.Put("/episodes/:id", nil)
	r.Delete("/episodes/:id", nil)
	log.Println("episode routes registered")

}
