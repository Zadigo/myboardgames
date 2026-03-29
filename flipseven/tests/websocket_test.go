package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func newWebsocketConnection(t *testing.T, handler http.HandlerFunc) (*websocket.Conn, *httptest.Server) {
	t.Helper()

	server := httptest.NewServer()
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}
	return conn, server
}

func TestLiveHandlerInitialConnection(t *testing.T) {
	conn, server := newWebsocketConnection(t, liveGameHandler)

	defer server.Close()
	defer conn.Close()

	message := ReadWebsocketMessage(conn)

	sendMessage(t, conn, WebsocketMessage{
		Action:  "waiting_lobby",
		TableId: "", // missing table ID
	})

	msg := readMessage(t, conn)

	if msg.Action != "error" {
		t.Errorf("expected action 'error', got '%s'", msg.Action)
	}

	if msg.Message != "Table ID is required" {
		t.Errorf("expected 'Table ID is required', got '%s'", msg.Message)
	}
}
