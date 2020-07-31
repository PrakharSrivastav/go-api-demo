package breakingbad

// Each entity can be further broken into subpackages
type ApiClient interface {
	// Characters
	AllCharacters() ([]Response, error)
	OneCharacter(ID int) (Response, error)

	// Episodes
	AllEpisodes() ([]string, error)
	OneEpisode(ID int) (string, error)

	// Quotes
	AllQuotes() ([]string, error)
	OneQuote(ID int) (string, error)
}
