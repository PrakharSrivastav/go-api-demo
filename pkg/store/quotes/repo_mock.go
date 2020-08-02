package quotes

import (
	"context"
	"syreclabs.com/go/faker"
)

type mock struct {
	list []Quote
}

func (m mock) Get(ctx context.Context, ID int) (*Quote, error) {
	for _, item := range m.list {
		if item.ID == ID {
			return &item, nil
		}
	}
	return &Quote{},nil
}

func (m mock) Find(ctx context.Context, s Search) ([]Quote, error) {
	// return a random subset
	ch := make([]Quote, 0, 5)
	for i := 0; i < 5; i++ {
		ch = append(ch, m.list[i])
	}
	return ch, nil
}

func (m mock) Insert(ctx context.Context, ch *Quote) error {
	m.list = append(m.list, *ch)
	return nil
}

func (m mock) Update(ctx context.Context, ID int, ch *Quote) error {
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

func NewMockRepo() Repository{
	quotes := make([]Quote, 0, 10)
	for i := 0; i < 10; i++ {
		quotes = append(quotes, Quote{
			ID:     i,
			Quote:  faker.Lorem().Sentence(15),
			Author: faker.Name().Name(),
			Series: "",
		})
	}
	return &mock{list: quotes}
}