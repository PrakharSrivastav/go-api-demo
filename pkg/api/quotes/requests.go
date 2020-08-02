package quotes

import (
	"github.com/PrakharSrivastav/go-api-demo/pkg/api/helpers"
	"github.com/PrakharSrivastav/go-api-demo/pkg/store/quotes"
)

type Request struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
	Series string `json:"series"`
}

// Validate input request
func (r *Request) Validate() (field string, err error) {
	if *r == (Request{}) {
		return "body", helpers.ErrorRequestBodyEmpty
	}

	// Lets say quote is always mandatory
	if r.Quote == "" {
		return "quote", helpers.ErrorRequestFieldMissing
	}

	return "", nil
}

// ToEntity func converts input request to db entity
// We can have similar converters for other stuff like external http clients etc
func (r *Request) ToEntity() (*quotes.Quote, error) {
	if *r == (Request{}) {
		return nil, helpers.ErrorConvesionToEntity
	}

	return &quotes.Quote{
		Quote:  r.Quote,
		Author: r.Author,
		Series: r.Series,
	}, nil
}
