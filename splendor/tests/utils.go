package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Zadigo/splendor/internal/backend"
	"github.com/Zadigo/splendor/internal/handlers"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func CreateWebsocketConnection(t *testing.T) (*websocket.Conn, *httptest.Server, error) {
	serverRegistry := backend.NewServerRegistry()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Origin", "http://localhost:3000")
		handlers.LiveGameHandler(w, r, serverRegistry)
	})

	server := httptest.NewServer(handler)
	url := "ws" + server.URL[4:] + "/splendor/live"

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	assert.NotNil(t, conn, "Expected connection to be established")

	return conn, server, err
}
