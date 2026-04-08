package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Zadigo/oriflamme/internal/backend"
	"github.com/Zadigo/oriflamme/internal/handlers"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func createLiveGameConn(t *testing.T) (*websocket.Conn, *httptest.Server) {
	t.Helper()

	redisClient := backend.NewRedisClient("redis://@localhost:6379/0")
	sr := backend.NewServerRegistry(redisClient)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Origin", "http://localhost:3000")
		handlers.LiveGameHandler(w, r, sr)
	})

	server := httptest.NewServer(handler)

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/oriflamme/live"
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

	assert.NotNil(t, conn)
	assert.NoError(t, err)

	time.Sleep(50 * time.Millisecond)

	return conn, server
}

func TestLiveGameHandler(t *testing.T) {
	conn, server := createLiveGameConn(t)

	defer conn.Close()
	defer server.Close()

	t.Run("Initial connection", func(t *testing.T) {

	})
}
