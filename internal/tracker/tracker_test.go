package tracker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrackerAddItem(t *testing.T) {
	t.Parallel()

	t.Run("success add", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}

		actual, err := tracker.AddItem(item)
		assert.NoError(t, err)
		assert.Equal(t, item, actual)
		assert.Equal(t, []Item{item}, tracker.GetItems())
	})

	t.Run("error add - already exists", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}
		_, err := tracker.AddItem(item)
		assert.NoError(t, err)

		actual, err := tracker.AddItem(Item{ID: "1", Name: "Second Item"})
		assert.ErrorIs(t, err, ErrAlreadyExists)
		assert.Equal(t, Item{}, actual)
		assert.Equal(t, []Item{item}, tracker.GetItems())
	})
}

func TestTrackerUpdateItem(t *testing.T) {
	t.Parallel()

	t.Run("error update - not found", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tracker.UpdateItem(item)
		assert.ErrorIs(t, err, ErrNotFound)
	})
}
