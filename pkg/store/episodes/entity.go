package episodes

type Episode struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Season     int      `json:"season"`
	Episode    int      `json:"episode"`
	AirDate    string   `json:"air_date"`
	Series     string   `json:"series"`
}
