package quotes

import (
	"encoding/json"
	"github.com/PrakharSrivastav/go-api-demo/pkg/store/quotes"
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
	ch, err := a.repoQuotes.Get(r.Context(), ID)
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
	q, err := a.clientApi.Quote(ID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	entity := &quotes.Quote{
		ID:     q.ID,
		Quote:  q.Quote,
		Author: q.Author,
		Series: q.Series,
	}

	// save to db before returning
	if err = a.repoQuotes.Insert(r.Context(), entity); err != nil {
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
