package tests

import (
	"context"
	"myboardgames/minimalgo/internal/backend"
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/internal/server"
	"myboardgames/minimalgo/internal/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsProduction(t *testing.T) {
	t.Run("should return false", func(t *testing.T) {
		result := utils.IsProduction()
		assert.False(t, result)
	})
}

func TestMainserver(t *testing.T) {
	client, err := backend.CreateRedisConnection(t.Context())
	assert.NoError(t, err)
	defer client.Close()

	t.Run("should instantiate a new server", func(t *testing.T) {
		s, err := server.NewServerApp(t.Context(), &models.ServerAppOptions{
			RedisClient: client,
		})
		assert.NoError(t, err)
		assert.NotNil(t, s)
	})

	t.Run("should start server", func(t *testing.T) {

		tCtx, cancel := context.WithTimeout(t.Context(), 15*time.Second)
		defer cancel()

		s, err := server.NewServerApp(t.Context(), &models.ServerAppOptions{
			RedisClient: client,
		})
		assert.NotNil(t, s)
		assert.NoError(t, err)

		err = s.Start(tCtx)

		assert.NoError(t, err)
		assert.True(t, s.GetState())
	})
}
