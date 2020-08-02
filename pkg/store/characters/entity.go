package characters

const (
	statusAlive    string = "Alive"
	statusDeceased string = "Deceased"
)

type Character struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Birthday  string `json:"birthday" db:"birthday"`
	Img       string `json:"img" db:"img"`
	Status    string `json:"status" db:"status"`
	Nickname  string `json:"nickname" db:"nickname"`
	Portrayed string `json:"portrayed" db:"portrayed"`
}
