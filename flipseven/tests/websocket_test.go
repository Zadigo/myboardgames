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
	"github.com/Zadigo/flipseven/internal/logic"
	"github.com/gorilla/websocket"
)

func init() {
	// Allow all origins in tests since ports are dynamic
	internal.RequestUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	internal.Tables = make(map[string]*logic.PlayersTable)
	internal.Tables["test-table-id"] = &logic.PlayersTable{}
}

func newWebsocketConnection(t *testing.T, handler internal.ContextHandlerFunc) (*websocket.Conn, *httptest.Server, context.CancelFunc) {
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

	time.Sleep(90 * time.Millisecond)

	return conn, server, cancel
}

func TestInitialConnection(t *testing.T) {
	conn, server, cancel := newWebsocketConnection(t, internal.LiveGameHandler)

	defer cancel()
	defer server.Close()
	defer conn.Close()

	message := internal.ReadWsMessage(conn)

	if message.Action != "initial_connection" {
		t.Error("Action should be initial connection")
	}
}

func TestWaitingLobby(t *testing.T) {
	conn, server, cancel := newWebsocketConnection(t, internal.LiveGameHandler)

	defer cancel()
	defer server.Close()
	defer conn.Close()

	internal.ReadWsMessage(conn)

	internal.WriteWsMessage(conn, internal.WebsocketMessage{
		Action:  "waiting_lobby",
		TableId: "test-table-id",
	})

	message := internal.ReadWsMessage(conn)
	fmt.Print(message)

	_, ok := internal.Tables["test-table-id"]
	if !ok {
		t.Error("Table should exist in the server's table map")
	}

	table := internal.Tables["test-table-id"]
	if table.NumberOfPlayers != 1 {
		t.Error("Table should have 1 player after joining")
	}

	if len(table.Clients) != 1 {
		t.Error("Table should have 1 client connection after joining")
	}

	message = internal.ReadWsMessage(conn)
	fmt.Print(message)

	if message.Action != "waiting_lobby" {
		t.Error("Action should be waiting_lobby")
	}

	cancel()
}

func TestGetDeck(t *testing.T) {
	conn, server, cancel := newWebsocketConnection(t, internal.LiveGameHandler)

	defer cancel()
	defer server.Close()
	defer conn.Close()

	internal.ReadWsMessage(conn)

	internal.WriteWsMessage(conn, internal.WebsocketMessage{
		Action:  "waiting_lobby",
		TableId: "test-table-id",
	})

	internal.ReadWsMessage(conn)

	internal.WriteWsMessage(conn, internal.WebsocketMessage{
		Action:  "get_deck",
		TableId: "test-table-id",
	})

	message := internal.ReadWsMessage(conn)
	fmt.Print(message)

	if message.Action != "deck_created" {
		t.Error("Action should be deck_created")
	}

	if message.Deck == nil {
		t.Error("Deck data should not be empty")
	}

	cancel()
}
