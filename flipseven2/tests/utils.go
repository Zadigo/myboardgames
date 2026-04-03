package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Zadigo/flipseven2/internal/handlers"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

func NewredisConn(t *testing.T) *redis.Client {
	options, _ := redis.ParseURL("redis://:@localhost:6379/0")
	return redis.NewClient(options)
}

func NewWebsocketConnection(t *testing.T) (*websocket.Conn, *httptest.Server) {
	t.Helper()

	redisClient := NewredisConn(t)

	// s := broadcasting.NewSubscription(redisClient, t.Context())
	// b := broadcasting.NewBroadcastingRegistry(s, t.Context())

	// br := models.CreateNewBaseRegistry()

	wrappedHanlder := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GameEngine(w, r, redisClient, nil, nil)
	})

	server := httptest.NewServer(wrappedHanlder)
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}

	return conn, server
}
