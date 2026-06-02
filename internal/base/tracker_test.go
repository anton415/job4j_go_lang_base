package base_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Tracker(t *testing.T) {
	t.Run("tracker fields are unexported", func(t *testing.T) {
		trackerType := reflect.TypeOf(*base.NewTracker())

		for i := range trackerType.NumField() {
			assert.False(t, trackerType.Field(i).IsExported())
		}
	})

	t.Run("get items without link leak", func(t *testing.T) {
		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()
		res[0].Name = "Second Item"

		assert.Equal(t, []base.Item{item}, tracker.GetItems())
	})
}
