package breakingbad

import (
	"net/http"
	"time"
)

func NewHttpClient() ApiClient {
	h := httpClient{
		baseURL: "https://www.breakingbadapi.com/api/",
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	return h
}

type httpClient struct {
	client  *http.Client
	baseURL string
}

func (h httpClient) Characters() ([]Character, error) {
	panic("implement me")
}

func (h httpClient) Character(ID int) (Character, error) {
	panic("implement me")
}

func (h httpClient) Episodes() ([]Episode, error) {
	panic("implement me")
}

func (h httpClient) Episode(ID int) (Episode, error) {
	panic("implement me")
}

func (h httpClient) Quotes() ([]Quote, error) {
	panic("implement me")
}

func (h httpClient) Quote(ID int) (Quote, error) {
	panic("implement me")
}

