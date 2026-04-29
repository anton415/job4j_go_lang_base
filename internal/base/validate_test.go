package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserID:      "1",
			Title:       "title",
			Description: "description",
		}
		res := base.Validate(req)

		assert.Equal(t, []string([]string{}), res)
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			Title:       "title",
			Description: "description",
		}
		res := base.Validate(req)

		assert.Equal(t, []string{"user id is empty"}, res)
	})

	t.Run("invalid empty title", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserID:      "1",
			Description: "description",
		}
		res := base.Validate(req)

		assert.Equal(t, []string{"title is empty"}, res)
	})

	t.Run("invalid empty description", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserID: "1",
			Title:  "title",
		}
		res := base.Validate(req)

		assert.Equal(t, []string{"description is empty"}, res)
	})

	t.Run("invalid empty fields", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{}
		res := base.Validate(req)

		assert.Equal(t, []string{"user id is empty", "title is empty", "description is empty"}, res)
	})

	t.Run("invalid nil request", func(t *testing.T) {
		t.Parallel()

		res := base.Validate(nil)

		assert.Equal(t, []string{"request is nil"}, res)
	})
}
