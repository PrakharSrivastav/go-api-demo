package breakingbad

import "syreclabs.com/go/faker"

type mock struct {
	list []Response
}

func (m *mock) AllCharacters() ([]Response, error) {
	panic("implement me")
}

func (m *mock) OneCharacter(ID int) (Response, error) {
	panic("implement me")
}

func (m *mock) AllEpisodes() ([]string, error) {
	panic("implement me")
}

func (m *mock) OneEpisode(ID int) (string, error) {
	panic("implement me")
}

func (m *mock) AllQuotes() ([]string, error) {
	panic("implement me")
}

func (m *mock) OneQuote(ID int) (string, error) {
	panic("implement me")
}

func NewMockApiClient() ApiClient {
	res := make([]Response, 0, 10)
	for i := 0; i < 10; i++ {
		res = append(res, Response{
			Name:      faker.Name().Name(),
			Birthday:  faker.Date().Birthday(10, 55).String(),
			Img:       faker.Internet().Url(),
			Status:    faker.RandomChoice([]string{"Alive", "Deceseased"}),
			Nickname:  faker.Name().FirstName(),
			Portrayed: faker.Name().Name(),
		})
	}
	return &mock{list: res}
}
