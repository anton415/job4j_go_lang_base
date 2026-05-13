package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewTracker(t *testing.T) {
	t.Parallel()

	t.Run("check link leak", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()
		res[0].Name = "Second Item"

		assert.Equal(t,
			[]Item{item},
			tracker.GetItems(),
		)
	})
}
