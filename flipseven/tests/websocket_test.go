package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Zadigo/flipseven/internal"
	"github.com/gorilla/websocket"
)

func newWebsocketConnection(t *testing.T, handler http.HandlerFunc) (*websocket.Conn, *httptest.Server) {
	t.Helper()

	server := httptest.NewServer(handler)
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}
	
	return conn, server
}

func writeWsMessage(t *testing.T, connection *websocket.Conn, message internal.WebsocketMessage) {
	t.Helper()

	err := connection.WriteJSON(message)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}
}

func TestLiveHandlerInitialConnection(t *testing.T) {
	conn, server := newWebsocketConnection(t, internal.LiveGameHandler)

	defer server.Close()
	defer conn.Close()

	message := internal.ReadWsMessage(conn)

	writeWsMessage(t, conn, internal.WebsocketMessage{
		Action:  "waiting_lobby",
		TableId: "", // missing table ID
	})

	if message.Action != "error" {
		t.Errorf("expected action 'error', got '%s'", message.Action)
	}

	if message.Message != "Table ID is required" {
		t.Errorf("expected 'Table ID is required', got '%s'", message.Message)
	}
}
