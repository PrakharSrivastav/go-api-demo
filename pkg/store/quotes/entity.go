package quotes

type Quote struct {
	ID     int    `json:"id" db:"id"`
	Quote  string `json:"quote" db:"quote"`
	Author string `json:"author" db:"author"`
	Series string `json:"series" db:"series"`
}
