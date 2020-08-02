package episodes

import (
	"context"
	"syreclabs.com/go/faker"
	"time"
)

type mock struct {
	list []Episode
}

func (m mock) Get(ctx context.Context, ID int) (*Episode, error) {
	for _, item := range m.list {
		if item.ID == ID {
			return &item, nil
		}
	}
	return &Episode{}, nil
}

func (m mock) Find(ctx context.Context, s Search) ([]Episode, error) {
	// return a random subset
	ch := make([]Episode, 0, 5)
	for i := 0; i < 5; i++ {
		ch = append(ch, m.list[i])
	}
	return ch, nil
}

func (m mock) Insert(ctx context.Context, ch *Episode) error {
	m.list = append(m.list, *ch)
	return nil
}

func (m mock) Update(ctx context.Context, ID int, ch *Episode) error {
	for i, item := range m.list {
		if item.ID == ID {
			m.list[i] = *ch
		}
	}
	return nil
}

func NewMockRepo() Repository {
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
	return &mock{list: episodes}
}
