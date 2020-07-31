package characters

const (
	statusAlive    string = "Alive"
	statusDeceased string = "Deceased"
)

type Character struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Birthday  string `json:"birthday"`
	Img       string `json:"img"`
	Status    string `json:"status"`
	Nickname  string `json:"nickname"`
	Portrayed string `json:"portrayed"`
}
