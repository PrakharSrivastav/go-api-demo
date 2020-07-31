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

func (h httpClient) AllCharacters() ([]Response, error) {
	panic("implement me")
}

func (h httpClient) OneCharacter(ID int) (Response, error) {
	panic("implement me")
}

func (h httpClient) AllEpisodes() ([]string, error) {
	panic("implement me")
}

func (h httpClient) OneEpisode(ID int) (string, error) {
	panic("implement me")
}

func (h httpClient) AllQuotes() ([]string, error) {
	panic("implement me")
}

func (h httpClient) OneQuote(ID int) (string, error) {
	panic("implement me")
}
