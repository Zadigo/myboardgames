package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Zadigo/oriflamme/internal/backend"
	"github.com/Zadigo/oriflamme/internal/backend/redisclient"
	"github.com/Zadigo/oriflamme/internal/handlers"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func createLiveGameConn(t *testing.T) (*websocket.Conn, *httptest.Server) {
	t.Helper()

	redisClient := redisclient.NewRedisClient("redis://@localhost:6379/0")
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
		message := backend.WebsocketMessage{}
		err := conn.ReadJSON(&message)
		assert.NoError(t, err, err)
		assert.Equal(t, "must_identify", message.Action)
		assert.NotEmpty(t, message.PlayerUuid)

		conn.WriteJSON(backend.WebsocketMessage{
			Action:     "identify",
			PlayerUuid: message.PlayerUuid,
			Username:   "pauline",
		})
	})

	t.Run("Create game", func(t *testing.T) {})

	t.Run("Another player joins the game", func(t *testing.T) {})

	t.Run("Start game", func(t *testing.T) {})

	t.Run("Players perform actions play card place card", func(t *testing.T) {})

	t.Run("Player A performs play card action reveal", func(t *testing.T) {})

	t.Run("Player B performs play card action place token", func(t *testing.T) {})

	t.Run("Resolve queue", func(t *testing.T) {})

	t.Run("Player A performs play card action stack card", func(t *testing.T) {})

	t.Run("Player B performs unrecognized card action", func(t *testing.T) {})

	t.Run("Player B performs unrecognized action", func(t *testing.T) {})

	t.Run("Player disconnects", func(t *testing.T) {})

	t.Run("Player reconnects within a given time limit", func(t *testing.T) {})

	t.Run("End game", func(t *testing.T) {})
}
