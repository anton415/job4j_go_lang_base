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
			// UserID не заполнен
			Title:       "title",
			Description: "description",
		}
		res := base.Validate(req)

		assert.Equal(t, []string{"validate message"}, res)
	})
}
