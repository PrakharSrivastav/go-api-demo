package breakingbad

type Character struct {
	ID        int    `json:"char_id"`
	Name      string `json:"name"`
	Birthday  string `json:"birthday"`
	Img       string `json:"img"`
	Status    string `json:"status"`
	Nickname  string `json:"nickname"`
	Portrayed string `json:"portrayed"`
	Category  string `json:"category,omitempty"`
}

type Episode struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Season  int    `json:"season"`
	Episode int    `json:"episode"`
	AirDate string `json:"air_date"`
	Series  string `json:"series"`
}

type Quote struct {
	ID     int    `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
	Series string `json:"series"`
}
