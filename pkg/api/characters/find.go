package characters

import (
	"encoding/json"
	"net/http"
)

func (a *API) find(w http.ResponseWriter, _ *http.Request) {
	c, err := a.clientApi.Characters()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
	return
}
