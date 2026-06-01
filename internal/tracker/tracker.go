package tracker

import (
	"fmt"
	"strings"
)

type Item struct {
	ID   string
	Name string
}

func (i Item) toString() string {
	return fmt.Sprintf("%s\t%s", i.ID, i.Name)
}

type Tracker struct {
	items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) (Item, error) {
	_, ok := t.indexOf(item.ID)
	if ok {
		return Item{}, ErrAlreadyExists
	}
	t.items = append(t.items, item)
	return item, nil
}

func (t *Tracker) GetItems() []Item {
	items := make([]Item, len(t.items))
	copy(items, t.items)
	return items
}

func (t *Tracker) UpdateItem(item Item) error {
	index, ok := t.indexOf(item.ID)
	if !ok {
		return ErrNotFound
	}
	t.items[index] = item
	return nil
}

func (t *Tracker) indexOf(id string) (int, bool) {
	for index, item := range t.items {
		if item.ID == id {
			return index, true
		}
	}
	return 0, false
}

func (t *Tracker) DeleteItem(id string) {
	for index, item := range t.items {
		if item.ID == id {
			t.items = append(t.items[:index], t.items[index+1:]...)
		}
	}
}

func (t *Tracker) SearchItems(name string) []Item {
	items := make([]Item, 0)
	for _, item := range t.items {
		if strings.Contains(item.Name, name) {
			items = append(items, item)
		}
	}
	return items
}
