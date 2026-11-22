package tests

import (
	"context"
	"myboardgames/minimalgo/internal/backend"
	"myboardgames/minimalgo/internal/models"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	rootDir, _ := filepath.Abs("..")

	ctx := context.WithValue(t.Context(), "rootDir", rootDir)
	redisDb, _ := backend.CreateRedisConnection(ctx)

	t.Run("should load routes", func(t *testing.T) {
		router, err := models.NewRouter(ctx, redisDb)
		assert.NoError(t, err)
		assert.NotNil(t, router)
		assert.IsType(t, router, &models.Router{})
	})

	t.Run("should find route", func(t *testing.T) {
		r, _ := models.NewRouter(ctx, redisDb)
		assert.NotNil(t, r.Find("some name"))
	})
}
