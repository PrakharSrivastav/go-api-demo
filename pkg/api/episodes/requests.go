package episodes

import (
	"github.com/PrakharSrivastav/go-api-demo/pkg/api"
	"github.com/PrakharSrivastav/go-api-demo/pkg/store/episodes"
)

// Request is input request structure .
// It would be in a separate pkg of itself if generated via swagger or graphql schema
type Request struct {
	Title   string `json:"title"`
	Season  int    `json:"season"`
	Episode int    `json:"episode"`
	AirDate string `json:"air_date"`
	Series  string `json:"series"`
}


// Validate input request
func (r *Request) Validate() (field string, err error) {
	if *r == (Request{}) {
		return "body", api.ErrorRequestBodyEmpty
	}

	// Lets say title is always mandatory
	if r.Title == "" {
		return "title", api.ErrorRequestFieldMissing
	}

	return "", nil
}

// ToEntity func converts input request to db entity
// We can have similar converters for other stuff like external http clients etc
func (r *Request) ToEntity() (*episodes.Episode, error) {
	if *r == (Request{}) {
		return nil, api.ErrorConvesionToEntity
	}

	return &episodes.Episode{
		Title:      r.Title,
		Season:     r.Season,
		Episode:    r.Episode,
		AirDate:    r.AirDate,
		Series:     r.Series,
	}, nil
}
