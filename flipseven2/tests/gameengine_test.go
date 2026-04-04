package tests

import (
	"net/http"
	"testing"
	"time"

	"github.com/Zadigo/flipseven2/internal/handlers"
	"github.com/Zadigo/flipseven2/internal/models"
)

func init() {
	// Allow all origins in tests since ports are dynamic
	handlers.RequestUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	// handlers.Tables = make(map[string]*logic.TableLayer)
	// handlers.Tables["test-table-id"] = &logic.TableLayer{
	// 	Layer: &logic.PlayersTable{},
	// }
}

func TestInitialConnection(t *testing.T) {
	conn, server := NewWebsocketConnection(t)

	defer server.Close()

	message := &models.WebsocketMessage{}
	err := conn.ReadJSON(message)

	if err != nil {
		t.Fatalf("Failed to read initial connection message: %v", err)
	}

	if message.Action != "initial_connection" {
		t.Fatalf("Expected initial_connection action, got %s", message.Action)
	}

	time.Sleep(2 * time.Second)

	err = conn.WriteJSON(models.WebsocketMessage{
		Action:   "initiate_table",
		Username: "julie95",
	})

	if err != nil {
		t.Fatalf("Failed to send initiate_table message: %v", err)
	}

	err = conn.ReadJSON(message)

	if err != nil {
		t.Fatalf("Failed to read response message: %v", err)
	}

	if message.Action != "table_initiated" {
		t.Fatalf("Expected table_initiated action, got %s", message.Action)
	}
}

func TestAcceptPlayer(t *testing.T) {
	conn, server := NewWebsocketConnection(t)

	defer server.Close()

	message := &models.WebsocketMessage{}
	err := conn.ReadJSON(message)

	if err != nil {
		t.Fatalf("Failed to read initial connection message: %v", err)
	}

	if message.Action != "initial_connection" {
		t.Fatalf("Expected initial_connection action, got %s", message.Action)
	}

	time.Sleep(2 * time.Second)

	err = conn.WriteJSON(models.WebsocketMessage{
		Action: "initiate_table",
	})

	message2 := &models.WebsocketMessage{}
	err = conn.ReadJSON(message2)
	if err != nil {
		t.Fatalf("Failed to read response message: %v", err)
	}

	if message2.Action != "table_initiated" {
		t.Fatalf("Expected table_initiated action, got %s", message2.Action)
	}

	err = conn.WriteJSON(models.WebsocketMessage{
		Action:   "accept_player",
		TableId:  message2.TableId,
		Username: "pauline88",
	})
}
