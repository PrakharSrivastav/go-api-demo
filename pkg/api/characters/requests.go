package characters

import (
	"github.com/PrakharSrivastav/go-api-demo/pkg/api/helpers"
	"github.com/PrakharSrivastav/go-api-demo/pkg/store/characters"
)

// Request is input request structure .
// It would be in a separate pkg of itself if generated via swagger or graphql schema
type Request struct {
	Name       string `json:"name"`
	Birthday   string `json:"birthday"`
	Img        string `json:"img"`
	Status     string `json:"status"`
	Nickname   string `json:"nickname"`
	Portrayed  string `json:"portrayed"`
}

// Validate input request
func (r *Request) Validate() (field string, err error) {
	if *r == (Request{}) {
		return "body", helpers.ErrorRequestBodyEmpty
	}

	// Lets say name is always mandatory
	if r.Name == "" {
		return "name", helpers.ErrorRequestFieldMissing
	}

	// other validations for birth date format, image URL format etc

	return "", nil
}

// ToEntity func converts input request to db entity
// We can have similar converters for other stuff like external http clients etc
func (r *Request) ToEntity() (*characters.Character, error) {
	if *r == (Request{}) {
		return nil, helpers.ErrorConvesionToEntity
	}

	return &characters.Character{
		Name:       r.Name,
		Birthday:   r.Birthday,
		Img:        r.Img,
		Status:     r.Status,
		Nickname:   r.Nickname,
		Portrayed:  r.Portrayed,
	}, nil
}
