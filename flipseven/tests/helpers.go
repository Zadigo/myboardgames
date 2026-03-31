package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Zadigo/flipseven/internal/handlers"
	"github.com/Zadigo/flipseven/internal/logic"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

const testTableId string = "ffc11d2a-cb17-4ceb-b01b-8c23c1cf47a8"

func NewWebsocketConnection(t *testing.T, handler handlers.ContextHandlerFunc) (*websocket.Conn, *httptest.Server, context.CancelFunc) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wrappedHanlder := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, ctx)
	})

	server := httptest.NewServer(wrappedHanlder)
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	conn, _, err := websocket.DefaultDialer.Dial(wsURL+"?table="+testTableId, nil)

	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}

	time.Sleep(50 * time.Millisecond)

	return conn, server, cancel
}

func LiveGameWebsocketConnection(t *testing.T) (*websocket.Conn, *httptest.Server, context.CancelFunc) {
	baseRegistry := handlers.NewBaseRegistry()
	baseRegistry.Set(testTableId, &logic.TableLayer{
		Layer: &logic.PlayersTable{
			TableId: testTableId,
		},
	})

	return NewWebsocketConnection(t, func(w http.ResponseWriter, r *http.Request, ctx context.Context) {
		handlers.LiveGameHandler(w, r, ctx, NewredisConn(t), nil, baseRegistry)
	})
}

func NewredisConn(t *testing.T) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
