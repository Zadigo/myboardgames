package tests

import (
	"net/http"
	"testing"
	"time"

	"github.com/Zadigo/flipseven2/internal/backend/broadcasting"
	"github.com/Zadigo/flipseven2/internal/handlers"
	"github.com/Zadigo/flipseven2/internal/models"
)

func init() {
	// Allow all origins in tests since ports are dynamic
	handlers.RequestUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	// handlers.BaseRegistry = models.CreateBaseRegistry()
}

func TestInitialConnection(t *testing.T) {
	redisClient := NewredisConn(t)

	s := broadcasting.NewSubscription(redisClient, t.Context())
	br := broadcasting.NewBroadcastingRegistry(s, t.Context())

	baseRegistry := models.CreateBaseRegistry()

	conn, server := NewWebsocketConnection(t, redisClient, br, baseRegistry)

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
	redisClient := NewredisConn(t)

	s := broadcasting.NewSubscription(redisClient, t.Context())
	br := broadcasting.NewBroadcastingRegistry(s, t.Context())

	baseRegistry := models.CreateBaseRegistry()

	// Conn 1: Host

	conn, server := NewWebsocketConnection(t, redisClient, br, baseRegistry)

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
		Username: "pauline95",
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

	// Conn 2: Player joining the table

	conn2, server2 := NewWebsocketConnection(t, redisClient, br, baseRegistry)

	defer server2.Close()

	message2 := &models.WebsocketMessage{}
	err = conn2.ReadJSON(message2)

	if err != nil {
		t.Fatalf("Failed to read response message: %v", err)
	}

	err = conn2.WriteJSON(models.WebsocketMessage{
		Action:   "accept_player",
		TableId:  message.TableId,
		Username: "pauline88",
	})
}
