package episodes

import (
	"encoding/json"
	"github.com/PrakharSrivastav/go-api-demo/pkg/store/episodes"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

func (a *API) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if in db, get it from there first
	ch, err := a.repoEpisode.Get(r.Context(), ID)
	if err != nil {
		// if db is in error, we can still get it from api
		// so no need to return http 400
		log.Println(err)
	}

	if ch != nil {
		b, err := json.Marshal(ch)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)
		return
	}

	// else get it from api
	epi, err := a.clientApi.Episode(ID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	entity := &episodes.Episode{
		ID:      epi.ID,
		Title:   epi.Title,
		Season:  epi.Season,
		Episode: epi.Episode,
		AirDate: epi.AirDate,
		Series:  epi.Series,
	}

	// save to db before returning
	if err = a.repoEpisode.Insert(r.Context(), entity); err != nil {
		log.Println(err)
	}

	b, err := json.Marshal(entity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(b)
	w.WriteHeader(http.StatusOK)
	return

}
