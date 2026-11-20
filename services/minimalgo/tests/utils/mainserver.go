package utils

import (
	"myboardgames/minimalgo/internal/http"
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

func HelperHttpServer(t *testing.T) *http.HttpApp {
	t.Helper()
	return &http.HttpApp{}
}

func HelperFullServer(t *testing.T) (*server.ServerApp, *http.HttpApp, chan models.InternalMessage) {
	ch := make(chan models.InternalMessage, 10)

	s := &server.ServerApp{
		ChInternal: ch,
	}

	h := &http.HttpApp{}
	h.SetServer(s)

	return s, h, ch
}
