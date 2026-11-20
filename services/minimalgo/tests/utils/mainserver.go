package utils

import (
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/internal/server"
	"testing"
)

func HelperMainServer(t *testing.T) (*server.ServerApp, chan models.InternalMessage) {
	t.Helper()

	ch := make(chan models.InternalMessage, 10)

	s := &server.ServerApp{
		ChInternal: ch,
	}

	return s, ch
}
