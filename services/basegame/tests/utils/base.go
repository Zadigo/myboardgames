package utils

import (
	"fmt"
	"maps"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Zadigo/basegame/internal/game/cards"
	"github.com/Zadigo/basegame/internal/game/games"
	"github.com/Zadigo/basegame/internal/handlers"
	"github.com/Zadigo/basegame/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

func NewRecorder(t *testing.T, method, path string, handler http.HandlerFunc) (*httptest.Server, *httptest.ResponseRecorder) {
	server := httptest.NewServer(handler)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, path, nil)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Origin", "http://localhost:3000")

	server.Config.Handler.ServeHTTP(recorder, request)
	return server, recorder
}

func NewWebsocketRecorder(t *testing.T, path string, handler http.HandlerFunc) (*websocket.Conn, *httptest.Server, *httptest.ResponseRecorder, error) {
	recorder := httptest.NewRecorder()

	originSequence := maps.Keys(utils.AllowedOrigins)

	var origin string
	for value := range originSequence {
		if origin == "" {
			origin = value
			break
		}
	}

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("Origin", origin)
			handler(w, r)
		}),
	)

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + path
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	return conn, server, recorder, err
}

func NewCreateRecorder(t *testing.T) (*httptest.Server, *httptest.ResponseRecorder) {
	handle := handlers.GenericHandler{}
	return NewRecorder(t, "POST", "/create", handle.CreateGameHandler)
}

func NewJoinRecorder(t *testing.T) (*websocket.Conn, *httptest.Server, *httptest.ResponseRecorder, error) {
	handle := handlers.GenericHandler{}
	return NewWebsocketRecorder(t, "/join", handle.JoinGameHandler)
}

func NewObserveRecorder(t *testing.T) (*websocket.Conn, *httptest.Server, *httptest.ResponseRecorder, error) {
	handle := handlers.GenericHandler{}
	return NewWebsocketRecorder(t, "/observe", handle.ObserveGameHandler)
}

func CreateTestCards() []cards.BaseCard {
	return []cards.BaseCard{
		{
			Uuid: uuid.NewString(),
		},
	}
}

func CreateRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func CreateClientPlayer(t *testing.T) *games.WebsocketClient {
	return games.NewWebsocketClient(nil)
}

func CreatePlayer(t *testing.T) *games.Player {
	return &games.Player{
		Username:    "some-username",
		IsInitiator: true,
	}
}

// CreatePileFlipOnTop creates a card pile with a Flip3 special card on top,
// followed by standard cards with values 0 to 3.
func CreatePileFlipOnTop(t *testing.T) (pile []cards.CardInterface) {
	c1 := &cards.BaseCard{
		Uuid:         "card-1",
		CardType:     cards.STANDARD_SPECIAL_CARD,
		CardValue:    0,
		CardOperator: cards.OPERATION_FLIP3,
	}

	for i := range 4 {
		item := &cards.BaseCard{
			Uuid:         fmt.Sprintf("card-2-%d", i),
			CardType:     cards.STANDARD_CARD,
			CardValue:    i,
			CardOperator: cards.OPERATION_ADD,
		}
		pile = append(pile, item)
	}

	return append([]cards.CardInterface{c1}, pile...)
}
