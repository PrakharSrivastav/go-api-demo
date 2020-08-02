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
	repoEpisode  episodes.Repository
	clientApi breakingbad.ApiClient
}

func New(repo episodes.Repository, client breakingbad.ApiClient) *API {
	return &API{
		repoEpisode:  repo,
		clientApi: client,
	}
}

func (a *API) Routes(r chi.Router) {
	r.Get("/episodes", a.find)
	r.Get("/episodes/{id}", a.get)
	r.Put("/episodes/{id}", a.edit)
	log.Println("episode routes registered")
}
