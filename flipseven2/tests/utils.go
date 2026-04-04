package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Zadigo/flipseven2/internal/backend/broadcasting"
	"github.com/Zadigo/flipseven2/internal/handlers"
	"github.com/Zadigo/flipseven2/internal/models"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

func NewredisConn(t *testing.T) *redis.Client {
	options, _ := redis.ParseURL("redis://:@localhost:6379/0")
	return redis.NewClient(options)
}

func NewWebsocketConnection(t *testing.T, redisClient *redis.Client, broadcastingRegistry *broadcasting.BroadcasterRegistry, baseRegistry *models.BaseRegistry) (*websocket.Conn, *httptest.Server) {
	t.Helper()

	wrappedHanlder := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GameEngine(w, r, redisClient, broadcastingRegistry, baseRegistry)
	})

	server := httptest.NewServer(wrappedHanlder)
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}

	return conn, server
}
