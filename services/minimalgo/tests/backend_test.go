package tests

import (
	"myboardgames/minimalgo/internal/backend"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	t.Run("should load routes", func(t *testing.T) {
		router, err := backend.NewRouter(t.Context(), "..")
		assert.NoError(t, err)
		assert.NotNil(t, router)
		assert.IsType(t, router, &backend.Router{})
	})

	t.Run("should find route", func(t *testing.T) {
		r, _ := backend.NewRouter(t.Context(), "..")
		assert.NotNil(t, r.Find("some name"))
	})

	t.Run("should return valid domain", func(t *testing.T) {
		r, _ := backend.NewRouter(t.Context(), "..")
		assert.NotNil(t, r)

		d, err := r.FindDomain("some name")
		assert.NoError(t, err)
		assert.Equal(t, "localhost:9001", d)
	})
}
