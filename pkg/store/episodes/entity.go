package episodes

type Episode struct {
	ID      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Season  int    `json:"season" db:"season"`
	Episode int    `json:"episode" db:"episode"`
	AirDate string `json:"air_date" db:"air_date"`
	Series  string `json:"series" db:"series"`
}
