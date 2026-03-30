package tests

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Zadigo/flipseven/internal"
	"github.com/Zadigo/flipseven/internal/backend"
	"github.com/Zadigo/flipseven/internal/handlers"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

func init() {
	// Allow all origins in tests since ports are dynamic
	handlers.RequestUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	handlers.Tables = make(map[string]*internal.PlayersTable)
	handlers.Tables["test-table-id"] = &internal.PlayersTable{}
}

func redisConn() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func newWebsocketConnection(t *testing.T, handler handlers.ContextHandlerFunc) (*websocket.Conn, *httptest.Server, context.CancelFunc) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wrappedHanlder := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, ctx)
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

func TestInitialConnection(t *testing.T) {
	pubSub := backend.NewSubscription(redisConn(), t.Context())
	registry := backend.NewRegistry(pubSub, t.Context())

	conn, server, cancel := newWebsocketConnection(t, func(w http.ResponseWriter, r *http.Request, ctx context.Context) {
		handlers.LiveGameHandler(w, r, context.Background(), redisConn(), registry)
	})

	defer cancel()
	defer server.Close()
	defer conn.Close()

	message := handlers.ReadWsMessage(conn)

	if message.Action != "initial_connection" {
		t.Error("Action should be initial connection")
	}

	message = handlers.ReadWsMessage(conn)
	fmt.Print(message)
}

func TestWaitingLobby(t *testing.T) {
	pubSub := backend.NewSubscription(redisConn(), t.Context())
	registry := backend.NewRegistry(pubSub, t.Context())

	conn, server, cancel := newWebsocketConnection(t, func(w http.ResponseWriter, r *http.Request, ctx context.Context) {
		handlers.LiveGameHandler(w, r, context.Background(), redisConn(), registry)
	})

	defer cancel()
	defer server.Close()
	defer conn.Close()

	handlers.ReadWsMessage(conn)

	handlers.WriteWsMessage(conn, handlers.WebsocketMessage{
		Action:  "waiting_lobby",
		TableId: "test-table-id",
	})

	message := handlers.ReadWsMessage(conn)
	fmt.Print(message)

	_, ok := handlers.Tables["test-table-id"]
	if !ok {
		t.Error("Table should exist in the server's table map")
	}

	table := handlers.Tables["test-table-id"]
	if table.NumberOfPlayers != 1 {
		t.Error("Table should have 1 player after joining")
	}

	if len(table.Clients) != 1 {
		t.Error("Table should have 1 client connection after joining")
	}

	message = handlers.ReadWsMessage(conn)
	fmt.Print(message)

	if message.Action != "waiting_lobby" {
		t.Error("Action should be waiting_lobby")
	}

	cancel()
}

func TestGetDeck(t *testing.T) {
	pubSub := backend.NewSubscription(redisConn(), t.Context())
	registry := backend.NewRegistry(pubSub, t.Context())

	conn, server, cancel := newWebsocketConnection(t, func(w http.ResponseWriter, r *http.Request, ctx context.Context) {
		handlers.LiveGameHandler(w, r, context.Background(), redisConn(), registry)
	})

	defer cancel()
	defer server.Close()
	defer conn.Close()

	handlers.ReadWsMessage(conn)

	handlers.WriteWsMessage(conn, handlers.WebsocketMessage{
		Action:  "waiting_lobby",
		TableId: "test-table-id",
	})

	handlers.ReadWsMessage(conn)

	handlers.WriteWsMessage(conn, handlers.WebsocketMessage{
		Action:  "get_deck",
		TableId: "test-table-id",
	})

	message := handlers.ReadWsMessage(conn)
	fmt.Print(message)

	if message.Action != "deck_created" {
		t.Error("Action should be deck_created")
	}

	if message.Deck == nil {
		t.Error("Deck data should not be empty")
	}

	cancel()
}
