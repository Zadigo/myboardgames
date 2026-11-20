package tests

import (
	"context"
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/internal/server"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMainserver(t *testing.T) {
	t.Run("should instantiate a new server", func(t *testing.T) {
		s := server.NewServerApp(t.Context(), &models.ServerAppOptions{})
		assert.NotNil(t, s)
		assert.IsType(t, s, &server.ServerApp{})
	})

	t.Run("should start server", func(t *testing.T) {
		tCtx, cancel := context.WithTimeout(t.Context(), 60*time.Second)
		defer cancel()

		s := server.NewServerApp(t.Context(), &models.ServerAppOptions{})
		s.Start(tCtx)

		assert.True(t, s.GetState())
	})
}
