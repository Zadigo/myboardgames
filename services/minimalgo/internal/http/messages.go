package http

import (
	"context"
	"myboardgames/minimalgo/internal/models"
)

func authenticationMessage(app *HttpApp, client models.WebsocketClientInterface, message models.BaseWebsocketMessage) {
	switch message.Action {
	case models.HelloPing:
		client.GetConn().WriteMessage(1, []byte("Hello"))
	default:
		return
	}
}

func globalGameMessage(app *HttpApp, client models.WebsocketClientInterface, message models.BaseWebsocketMessage) {
	// handle global game messages here
}

// gameRelatedMessage must implement all the related websocket logic to interact with the
// logic of the game in the backend
func gameRelatedMessage(app *HttpApp, client models.WebsocketClientInterface, message models.BaseWebsocketMessage) {
	// handle game messages here
	switch message.Action {
	case models.CreateTable:
		app.server.SendInstruction(context.TODO(), models.InternalInstruction{
			Client: client,
			Action: models.CreateTable,
			Sender: models.HttpAppSender,
		})
	default:
		return
	}
}
