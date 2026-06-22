package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Zadigo/basegame/internal/handlers"
)

func NewRecorder(t *testing.T, method, path string, handler http.HandlerFunc) *httptest.ResponseRecorder {
	server := http.NewServeMux()
	server.HandleFunc(path, handler)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, httptest.NewRequest(method, path, nil))
	return recorder
}

func NewCreateRecorder(t *testing.T) *httptest.ResponseRecorder {
	handle := handlers.GenericHandler{}
	return NewRecorder(t, "POST", "/create", handle.CreateGameHandler)
}

func NewJoinRecorder(t *testing.T) *httptest.ResponseRecorder {
	handle := handlers.GenericHandler{}
	return NewRecorder(t, "GET", "/join", handle.JoinGameHandler)
}

func NewObserveRecorder(t *testing.T) *httptest.ResponseRecorder {
	handle := handlers.GenericHandler{}
	return NewRecorder(t, "GET", "/observe", handle.ObserveGameHandler)
}
