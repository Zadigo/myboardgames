package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Zadigo/flipseven/internal"
	"github.com/gorilla/websocket"
)

func init() {
	// Allow all origins in tests since ports are dynamic
	internal.RequestUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
}

func newWebsocketConnection(t *testing.T, handler http.HandlerFunc) (*websocket.Conn, *httptest.Server) {
	t.Helper()

	server := httptest.NewServer(handler)
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

	if err != nil {
		t.Fatalf("Failed to connect to test server: %v", err)
	}

	time.Sleep(90 * time.Millisecond)

	return conn, server
}

func TestLiveHandlerInitialConnection(t *testing.T) {
	conn, server := newWebsocketConnection(t, internal.LiveGameHandler)

	defer server.Close()
	defer conn.Close()

	message := internal.ReadWsMessage(conn)
	log.Printf("Result: %v", message)

	if message.Action != "initial_connection" {
		t.Error("Action should be initial connection")
	}
}
