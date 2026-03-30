package tests

import (
	"net/http"

	"github.com/Zadigo/flipseven/internal/handlers"
	"github.com/Zadigo/flipseven/internal/logic"
)

func init() {
	// Allow all origins in tests since ports are dynamic
	handlers.RequestUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	handlers.Tables = make(map[string]*logic.TableLayer)
	handlers.Tables["test-table-id"] = &logic.TableLayer{}
}

// func redisConn() *redis.Client {
// 	return redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 	})
// }

// func TestInitialConnection(t *testing.T) {
// 	pubSub := backend.NewSubscription(redisConn(), t.Context())
// 	registry := backend.NewRegistry(pubSub, t.Context())

// 	conn, server, cancel := newWebsocketConnection(t, func(w http.ResponseWriter, r *http.Request, ctx context.Context) {
// 		handlers.LiveGameHandler(w, r, context.Background(), redisConn(), registry)
// 	})

// 	defer cancel()
// 	defer server.Close()
// 	defer conn.Close()

// 	message := handlers.ReadWsMessage(conn)

// 	if message.Action != "initial_connection" {
// 		t.Error("Action should be initial connection")
// 	}

// 	message = handlers.ReadWsMessage(conn)
// 	fmt.Print(message)
// }

// func TestWaitingLobby(t *testing.T) {
// 	pubSub := backend.NewSubscription(redisConn(), t.Context())
// 	registry := backend.NewRegistry(pubSub, t.Context())

// 	conn, server, cancel := newWebsocketConnection(t, func(w http.ResponseWriter, r *http.Request, ctx context.Context) {
// 		handlers.LiveGameHandler(w, r, context.Background(), redisConn(), registry)
// 	})

// 	defer cancel()
// 	defer server.Close()
// 	defer conn.Close()

// 	handlers.ReadWsMessage(conn)

// 	handlers.WriteWsMessage(conn, handlers.WebsocketMessage{
// 		Action:  "waiting_lobby",
// 		TableId: "test-table-id",
// 	})

// 	message := handlers.ReadWsMessage(conn)
// 	fmt.Print(message)

// 	_, ok := handlers.Tables["test-table-id"]
// 	if !ok {
// 		t.Error("Table should exist in the server's table map")
// 	}

// 	table := handlers.Tables["test-table-id"]
// 	if table.NumberOfPlayers != 1 {
// 		t.Error("Table should have 1 player after joining")
// 	}

// 	if len(table.Clients) != 1 {
// 		t.Error("Table should have 1 client connection after joining")
// 	}

// 	message = handlers.ReadWsMessage(conn)
// 	fmt.Print(message)

// 	if message.Action != "waiting_lobby" {
// 		t.Error("Action should be waiting_lobby")
// 	}

// 	cancel()
// }

// func TestGetDeck(t *testing.T) {
// 	pubSub := backend.NewSubscription(redisConn(), t.Context())
// 	registry := backend.NewRegistry(pubSub, t.Context())

// 	conn, server, cancel := newWebsocketConnection(t, func(w http.ResponseWriter, r *http.Request, ctx context.Context) {
// 		handlers.LiveGameHandler(w, r, context.Background(), redisConn(), registry)
// 	})

// 	defer cancel()
// 	defer server.Close()
// 	defer conn.Close()

// 	handlers.ReadWsMessage(conn)

// 	handlers.WriteWsMessage(conn, handlers.WebsocketMessage{
// 		Action:  "waiting_lobby",
// 		TableId: "test-table-id",
// 	})

// 	handlers.ReadWsMessage(conn)

// 	handlers.WriteWsMessage(conn, handlers.WebsocketMessage{
// 		Action:  "get_deck",
// 		TableId: "test-table-id",
// 	})

// 	message := handlers.ReadWsMessage(conn)
// 	fmt.Print(message)

// 	if message.Action != "deck_created" {
// 		t.Error("Action should be deck_created")
// 	}

// 	if message.Deck == nil {
// 		t.Error("Deck data should not be empty")
// 	}

// 	cancel()
// }
