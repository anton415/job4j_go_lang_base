package base_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Tracker(t *testing.T) {
	t.Parallel()

	t.Run("add item", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}

		tracker.AddItem(item)
		res := tracker.GetItems()
		expected := []base.Item{item}

		assert.Equal(t, expected, res)
	})

	t.Run("empty tracker", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		res := tracker.GetItems()
		expected := []base.Item{}

		assert.Equal(t, expected, res)
	})

	t.Run("get items without link leak", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()
		res[0].Name = "Second Item"
		expected := []base.Item{item}

		assert.Equal(t, expected, tracker.GetItems())
	})

	t.Run("append to items copy without link leak", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		items := []base.Item{
			{ID: "1", Name: "First Item"},
			{ID: "2", Name: "Second Item"},
			{ID: "3", Name: "Third Item"},
		}
		for _, item := range items {
			tracker.AddItem(item)
		}

		res := tracker.GetItems()
		res = append(res, base.Item{ID: "4", Name: "External Item"})

		tracker.AddItem(base.Item{ID: "4", Name: "Fourth Item"})
		res[3].Name = "Changed Item"
		expected := []base.Item{
			{ID: "1", Name: "First Item"},
			{ID: "2", Name: "Second Item"},
			{ID: "3", Name: "Third Item"},
			{ID: "4", Name: "Fourth Item"},
		}

		assert.Equal(t, expected, tracker.GetItems())
	})
}
