package characters

import (
	"context"
	"syreclabs.com/go/faker"
)

type mock struct {
	list []Character
}



func NewMockRepo() Repository {
	list := make([]Character, 0, 10)
	for i := 0; i < 10; i++ {
		list = append(list, Character{
			ID:        i,
			Name:      faker.Name().Name(),
			Birthday:  faker.Date().Birthday(40, 55).String(),
			Img:       faker.Internet().Url(),
			Status:    faker.RandomChoice([]string{statusAlive, statusDeceased}),
			Nickname:  faker.Name().FirstName(),
			Portrayed: faker.Name().Name(),
		})
	}
	return mock{list: list}
}

func (m mock) Get(ctx context.Context, ID int) (*Character, error) {
	for _, item := range m.list {
		if item.ID == ID {
			return &item, nil
		}
	}
	return &Character{}, nil
}

func (m mock) Find(ctx context.Context, s SearchCriteria) ([]Character, error) {
	// return a random subset
	ch := make([]Character, 0, 5)
	for i := 0; i < 5; i++ {
		ch = append(ch, m.list[i])
	}
	return ch, nil
}

func (m mock) Insert(ctx context.Context, ch *Character) error {
	m.list = append(m.list, *ch)
	return nil
}

func (m mock) Update(ctx context.Context, ID int, ch *Character) error {
	for i, item := range m.list {
		if item.ID == ID {
			m.list[i] = *ch
		}
	}
	return nil
}

func (m mock) Delete(ctx context.Context, ID int) error {
	return nil
}
