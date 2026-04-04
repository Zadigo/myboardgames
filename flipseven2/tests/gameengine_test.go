package tests

import (
	"fmt"
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

func TestGameEngine(t *testing.T) {
	conn, server := NewWebsocketConnection(t)

	defer server.Close()

	message := &models.WebsocketMessage{}
	err := conn.ReadJSON(message)

	if err != nil {
		t.Fatalf("Failed to read initial connection message: %v", err)
	}

	time.Sleep(2 * time.Second)

	fmt.Print(message)

	err = conn.WriteJSON(models.WebsocketMessage{
		Action:  "initiate_table",
		Username: "julie95",
	})

	if err != nil {
		t.Fatalf("Failed to send initiate_table message: %v", err)
	}

	err = conn.ReadJSON(message)

	if err != nil {
		t.Fatalf("Failed to read response message: %v", err)
	}

	fmt.Print(message)

	// msg = conn.ReadJSON(models.WebsocketMessage{})
	// fmt.Print(msg)

	// conn.WriteJSON(models.WebsocketMessage{
	// 	Action:  "initiate_table",
	// 	TableId: "test-table-id",
	// })

}
