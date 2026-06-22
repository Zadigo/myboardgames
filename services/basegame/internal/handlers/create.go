package handlers

import (
	"net/http"
)

type GenericHandler struct {
	BaseHandler
}

func (g *GenericHandler) CreateGameHandler(w http.ResponseWriter, r *http.Request) {
}

func (g *GenericHandler) JoinGameHandler(w http.ResponseWriter, r *http.Request) {

}

func (g *GenericHandler) ObserveGameHandler(w http.ResponseWriter, r *http.Request) {

}
