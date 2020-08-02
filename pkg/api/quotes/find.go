package quotes

import (
	"encoding/json"
	"net/http"
)

func (a *API) find(w http.ResponseWriter, _ *http.Request) {
	c, err := a.clientApi.Quotes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(b)
	w.WriteHeader(http.StatusOK)
	return
}
