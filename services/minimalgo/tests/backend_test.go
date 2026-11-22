package tests

import (
	"myboardgames/minimalgo/internal/backend"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	redisDb, _ := backend.CreateRedisConnection(t.Context())

	t.Run("should load routes", func(t *testing.T) {
		router, err := backend.NewRouter(t.Context(), redisDb, "..")
		assert.NoError(t, err)
		assert.NotNil(t, router)
		assert.IsType(t, router, &backend.Router{})
	})

	t.Run("should find route", func(t *testing.T) {
		r, _ := backend.NewRouter(t.Context(), redisDb, "..")
		assert.NotNil(t, r.Find("some name"))
	})
}
