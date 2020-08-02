package episodes

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (a *API) edit(w http.ResponseWriter, r *http.Request) {
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

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var model Request
	if err = json.Unmarshal(req, &model); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	field, err := model.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("err :: %v , field :: %s", err, field)
		return
	}

	entity, err := model.ToEntity()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = a.repoEpisode.Update(r.Context(), ID, entity); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
