package quotes

type Quote struct {
	ID     int    `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
	Series string `json:"series"`
}
