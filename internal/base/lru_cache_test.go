package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Lru_Cache(t *testing.T) {
	t.Parallel()

	t.Run("get exists", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(2)
		cache.Put("1", "one")

		rsl := cache.Get("1")

		assert.Equal(t, "one", *rsl)
	})

	t.Run("get missing", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(2)
		cache.Put("1", "one")

		rsl := cache.Get("2")

		assert.Nil(t, rsl)
	})

	t.Run("evicts least recently used", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(2)
		cache.Put("1", "one")
		cache.Put("2", "two")
		cache.Put("3", "three")

		rsl1 := cache.Get("1")
		rsl2 := cache.Get("2")
		rsl3 := cache.Get("3")

		assert.Nil(t, rsl1)
		assert.Equal(t, "two", *rsl2)
		assert.Equal(t, "three", *rsl3)
	})

	t.Run("get moves item to head", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(2)
		cache.Put("1", "one")
		cache.Put("2", "two")
		cache.Get("1")
		cache.Put("3", "three")

		rsl1 := cache.Get("1")
		rsl2 := cache.Get("2")
		rsl3 := cache.Get("3")

		assert.Equal(t, "one", *rsl1)
		assert.Nil(t, rsl2)
		assert.Equal(t, "three", *rsl3)
	})
}
