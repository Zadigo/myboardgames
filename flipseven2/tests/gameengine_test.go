package tests

import (
	"fmt"
	"net/http"
	"testing"

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
	conn, server, cancel := NewWebsocketConnection(t)

	defer server.Close()
	defer cancel()

	msg := conn.ReadJSON(models.WebsocketMessage{})
	fmt.Print(msg)

	conn.WriteJSON(models.WebsocketMessage{
		Action:  "initiate_table",
		TableId: "test-table-id",
	})

	msg = conn.ReadJSON(models.WebsocketMessage{})
	fmt.Print(msg)

	conn.WriteJSON(models.WebsocketMessage{
		Action:  "initiate_table",
		TableId: "test-table-id",
	})

}
