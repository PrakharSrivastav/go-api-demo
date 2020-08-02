package characters

import (
	"encoding/json"
	"github.com/PrakharSrivastav/go-api-demo/pkg/store/characters"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

func (a *API) get(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if in db, get it from there first
	ch, err := a.repoChar.Get(r.Context(), ID)
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
	char, err := a.clientApi.Character(ID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	entity := &characters.Character{
		ID:        char.ID,
		Name:      char.Name,
		Birthday:  char.Birthday,
		Img:       char.Img,
		Status:    char.Status,
		Nickname:  char.Name,
		Portrayed: char.Portrayed,
	}

	// save to db before returning
	if err = a.repoChar.Insert(r.Context(), entity); err != nil {
		log.Println(err)
	}

	b, err := json.Marshal(entity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
	w.WriteHeader(http.StatusOK)
	return

}
