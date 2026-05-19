package game

import "github.com/Zadigo/gosplendor/internal/models"

type AuthenticationActions struct {
	Client *models.WebsocketClient
}

func (h *AuthenticationActions) ParseMessage(message models.WebsocketMessage[models.AuthenticationMessage]) error {
	switch message.Action {
	case "identify":
		return h.Identify(message)
	default:
		return nil
	}
}

func (h *AuthenticationActions) Identify(message models.WebsocketMessage[models.AuthenticationMessage]) error {
	username := message.Username
	if username == "" {
		message := models.WebsocketMessage[models.AuthenticationMessage]{
			Action:  "error",
			Message: "Username cannot be empty.",
		}
		return message.SendJsonMessage()
	}
	return nil
}
