package handlers

import (
	"log"
	"net/http"
)

func (g *GenericHandler) CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the logic for creating a game
	log.Print("CreateGameHandler called")
	// g.app.GetServerApp().CreateGame(games.STANDARD)
	// g.app.GetServerApp().CreateGame(games.EXTENSION)
}
