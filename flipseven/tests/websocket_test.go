package tests

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/Zadigo/flipseven/internal/handlers"
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
	// pubSub := backend.NewSubscription(redisConn(), t.Context())
	// registry := backend.NewRegistry(pubSub, t.Context())

	conn, server, cancel := LiveGameWebsocketConnection(t)

	defer cancel()
	defer conn.Close()
	defer server.Close()

	message, _ := handlers.ReadWsMessage(conn)

	if message.Action != "initial_connection" {
		t.Error("Action should be initial connection")
	}

	time.Sleep(3 * time.Second)

	handlers.WriteWsMessage(conn, handlers.WebsocketMessage{
		Action: "initiate_table",
	})

	message, _ = handlers.ReadWsMessage(conn)
	log.Print(message)

	// handlers.WriteWsMessage(conn, handlers.WebsocketMessage{
	// 	Action:   "waiting_lobby",
	// 	PlayerId: "test-player",
	// })

	// message, _ = handlers.ReadWsMessage(conn)
	// log.Print(message)

	cancel()
}

func TestWaitingLobby(t *testing.T) {
	conn, server, cancel := LiveGameWebsocketConnection(t)

	defer cancel()
	defer conn.Close()
	defer server.Close()

	message, _ := handlers.ReadWsMessage(conn)

	time.Sleep(3 * time.Second)

	handlers.WriteWsMessage(conn, handlers.WebsocketMessage{
		Action:   "waiting_lobby",
		PlayerId: "test-player",
	})

	message, _ = handlers.ReadWsMessage(conn)
	log.Print(message)

	cancel()
}
