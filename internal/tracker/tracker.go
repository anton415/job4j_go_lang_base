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

func (t *Tracker) AddItem(item Item) {
	t.items = append(t.items, item)
}

func (t *Tracker) GetItems() []Item {
	items := make([]Item, len(t.items))
	copy(items, t.items)
	return items
}

func (t *Tracker) UpdateItem(id string, name string) {
	for index := range t.items {
		if t.items[index].ID == id {
			t.items[index].Name = name
		}
	}
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
