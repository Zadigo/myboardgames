package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Zadigo/flipseven2/internal/backend/broadcasting"
	"github.com/Zadigo/flipseven2/internal/handlers"
	"github.com/Zadigo/flipseven2/internal/models"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

func NewredisConn(t *testing.T) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func NewWebsocketConnection(t *testing.T) (*websocket.Conn, *httptest.Server, context.CancelFunc) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	redisClient := NewredisConn(t)

	s := broadcasting.NewSubscription(redisClient, t.Context())
	b := broadcasting.NewBroadcastingRegistry(s, t.Context())

	br := models.CreateNewBaseRegistry()

	wrappedHanlder := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handler(w, r, ctx, _redisClient, b, br)
		handlers.GameEngine(w, r, ctx, redisClient, b, br)
	})

	server := httptest.NewServer(wrappedHanlder)
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}

	time.Sleep(50 * time.Millisecond)

	return conn, server, cancel
}
