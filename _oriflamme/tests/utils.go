package tests

import (
	"fmt"
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

func CreateLiveGameConn(t *testing.T) (*websocket.Conn, *httptest.Server) {
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

func CreateMultipleGameConn(t *testing.T, usernames ...string) (map[string]*websocket.Conn, *httptest.Server) {
	t.Helper()

	redisClient := redisclient.NewRedisClient("redis://@localhost:6379/0")
	sr := backend.NewServerRegistry(redisClient)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Origin", "http://localhost:3000")
		handlers.LiveGameHandler(w, r, sr)
	})

	server := httptest.NewServer(handler)

	conns := make(map[string]*websocket.Conn)

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/oriflamme/live"
	for _, name := range usernames {
		conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

		time.Sleep(5 * time.Millisecond)

		assert.NotNil(t, conn)
		assert.NoError(t, err)
		conns[fmt.Sprintf("%s", name)] = conn
	}

	return conns, server
}

func ConvertToWebsocketClient(t *testing.T, conns map[string]*websocket.Conn) []*backend.WebsocketClient {
	t.Helper()

	totalConns := len(conns)

	clients := make([]*backend.WebsocketClient, 0, totalConns)
	for name, conn := range conns {
		client := backend.NewWebsocketClient(conn)
		client.Username = name
		clients = append(clients, client)
	}

	return clients
}
