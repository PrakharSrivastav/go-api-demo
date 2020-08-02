package breakingbad

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strconv"
	"time"
)

func NewHttpClient() ApiClient {
	r := resty.New().
		SetRetryCount(3).
		SetRetryMaxWaitTime(time.Second * 5)

	h := httpClient{
		baseURL: "https://www.breakingbadapi.com/api/",
		c:       r,
	}
	return h
}

type httpClient struct {
	c       *resty.Client
	baseURL string
}

func (h httpClient) Characters() ([]Character, error) {
	var result []Character

	resp, err := h.c. //SetDebug(true).
		R().
		SetResult(&result).
		Get(h.baseURL + "characters")

	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		return result, nil
	default:
		return nil, errors.New("not found")
	}

}

func (h httpClient) Character(ID int) ([]Character, error) {
	var result []Character

	resp, err := h.c.//SetDebug(true).
		R().
		SetResult(&result).
		SetPathParams(map[string]string{"id": strconv.Itoa(ID)}).
		Get(h.baseURL + "characters/{id}")

	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		return result, nil
	default:
		return nil, errors.New("not found")
	}
}

func (h httpClient) Episodes() ([]Episode, error) {
	var result []Episode

	resp, err := h.c. //SetDebug(true).
		R().
		SetResult(&result).
		Get(h.baseURL + "episodes")

	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		return result, nil
	default:
		return nil, errors.New("not found")
	}
}

func (h httpClient) Episode(ID int) (*Episode, error) {
	var result Episode

	resp, err := h.c. //SetDebug(true).
		R().
		SetResult(&result).
		SetPathParams(map[string]string{"id": strconv.Itoa(ID)}).
		Get(h.baseURL + "episodes/{id}")

	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		return &result, nil
	default:
		return nil, errors.New("not found")
	}
}

func (h httpClient) Quotes() ([]Quote, error) {
	var result []Quote

	resp, err := h.c. //SetDebug(true).
		R().
		SetResult(&result).
		Get(h.baseURL + "quotes")

	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		return result, nil
	default:
		return nil, errors.New("not found")
	}
}

func (h httpClient) Quote(ID int) (*Quote, error) {
	var result Quote

	resp, err := h.c. //SetDebug(true).
		R().
		SetResult(&result).
		SetPathParams(map[string]string{"id": strconv.Itoa(ID)}).
		Get(h.baseURL + "quotes/{id}")

	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		return &result, nil
	default:
		return nil, errors.New("not found")
	}
}
