package characters

import (
	"github.com/PrakharSrivastav/go-api-demo/pkg/clients/breakingbad"
	"github.com/PrakharSrivastav/go-api-demo/pkg/store/characters"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockApi() API {
	return API{
		repoChar:  characters.NewMockRepo(),
		clientApi: breakingbad.NewMockApiClient(),
	}
}

func TestFind(t *testing.T) {
	api := mockApi()
	req := httptest.NewRequest(http.MethodGet, "/characters", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.find)
	handler.ServeHTTP(rr, req)

	require.Equal(t, rr.Code, http.StatusOK)
}
