package breakingbad

// Each entity can be further broken into subpackages
type ApiClient interface {
	// Character
	Characters() ([]Character, error)
	Character(ID int) ([]Character, error)

	// Episode
	Episodes() ([]Episode, error)
	Episode(ID int) (*Episode, error)

	// Quote
	Quotes() ([]Quote, error)
	Quote(ID int) (*Quote, error)
}
