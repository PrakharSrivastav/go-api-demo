package breakingbad

import (
	"syreclabs.com/go/faker"
	"time"
)

type mock struct {
	chars    []Character
	episodes []Episode
	quotes   []Quote
}

func (m *mock) Characters() ([]Character, error) {
	return m.chars, nil
}

func (m *mock) Character(ID int) (Character, error) {
	for i := range m.chars {
		if m.chars[i].ID == ID {
			return m.chars[i], nil
		}
	}
	return Character{}, nil
}

func (m *mock) Episodes() ([]Episode, error) {
	return m.episodes, nil
}

func (m *mock) Episode(ID int) (Episode, error) {
	for i := range m.episodes {
		if m.episodes[i].ID == ID {
			return m.episodes[i], nil
		}
	}
	return Episode{}, nil
}

func (m *mock) Quotes() ([]Quote, error) {
	return m.quotes, nil
}

func (m *mock) Quote(ID int) (Quote, error) {
	for i := range m.quotes {
		if m.quotes[i].ID == ID {
			return m.quotes[i], nil
		}
	}
	return Quote{}, nil
}

func NewMockApiClient() ApiClient {
	chars := make([]Character, 0, 10)
	for i := 0; i < 10; i++ {
		chars = append(chars, Character{
			ID:        i,
			Name:      faker.Name().Name(),
			Birthday:  faker.Date().Birthday(10, 55).String(),
			Img:       faker.Internet().Url(),
			Status:    faker.RandomChoice([]string{"Alive", "Deceseased"}),
			Nickname:  faker.Name().FirstName(),
			Portrayed: faker.Name().Name(),
		})
	}

	episodes := make([]Episode, 0, 10)
	for i := 0; i < 10; i++ {
		episodes = append(episodes, Episode{
			ID:      i,
			Title:   faker.Team().Name(),
			Season:  faker.RandomInt(1, 5),
			Episode: faker.RandomInt(1, 23),
			AirDate: faker.Date().Between(time.Now().AddDate(-10, 0, 0), time.Now()).String(),
			Series:  "",
		})
	}

	quotes := make([]Quote, 0, 10)
	for i := 0; i < 10; i++ {
		quotes = append(quotes, Quote{
			ID:     i,
			Quote:  faker.Lorem().Sentence(15),
			Author: faker.Name().Name(),
			Series: "",
		})
	}
	return &mock{chars: chars, episodes: episodes, quotes: quotes}
}
