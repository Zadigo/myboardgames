package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Zadigo/basegame/internal/game/games"
	"github.com/Zadigo/basegame/internal/models"
	"github.com/Zadigo/basegame/internal/utils"
)

type RequestData struct {
	GameType string `json:"gameType"`
}

func (g *GenericHandler) CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	errors := &HandlerErrors{}
	
	body := &RequestData{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		utils.JsonResponse(w, errors.InvalidRequestBody(), http.StatusBadRequest)
		return
	}

	if body.GameType == "" {
		utils.JsonResponse(w, errors.MissingGameType(), http.StatusBadRequest)
		return
	}

	if body.GameType != games.STANDARD && body.GameType != games.EXTENSION {
		utils.JsonResponse(w, map[string]string{"error": "Invalid game type"}, http.StatusBadRequest)
		return
	}

	var newGame models.GameInterface
	var err error

	if body.GameType == games.STANDARD {
		newGame, err = g.app.GetServerApp().CreateGame(games.STANDARD)
	}

	if body.GameType == games.EXTENSION {
		newGame, err = g.app.GetServerApp().CreateGame(games.EXTENSION)
	}

	if err != nil {
		utils.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	if newGame == nil {
		utils.JsonResponse(w, map[string]string{"error": "Failed to create game"}, http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, map[string]string{"gameUUID": newGame.GetUuid()}, http.StatusOK)
}
