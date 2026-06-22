package tests

import (
	"maps"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Zadigo/basegame/internal/game/cards"
	"github.com/Zadigo/basegame/internal/handlers"
	"github.com/Zadigo/basegame/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
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
