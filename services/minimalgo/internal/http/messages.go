package http

import (
	"context"
	pb "myboardgames/gen/minimalgame/v1"
	"myboardgames/minimalgo/internal/models"
)

func authenticationMessage(app *HttpApp, client models.WebsocketClientInterface, message models.BaseWebsocketMessage) {
	switch message.Action {
	case models.PingRpc:
		rpc, ok := app.server.GetGrpcConnections()[message.Rpc]
		if !ok {
			client.GetConn().WriteJSON(models.BaseWebsocketMessage{
				ClientID:   client.GetID(),
				StrMessage: "rpc not found",
			})
			return
		} else {
			rpc.Connect()
			client.GetConn().WriteJSON(models.BaseWebsocketMessage{
				ClientID:   client.GetID(),
				StrMessage: "ok",
			})
		}
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
		// Run action
		req := &pb.CreateTableRequest{ClientId: client.GetID()}

		conns := app.server.GetGrpcConnections()
		conn, _ := conns[message.Rpc]

		err := conn.Invoke(context.Background(), pb.TableService_CreateTable_FullMethodName, req, nil)
		if err != nil {
			client.GetConn().WriteMessage(1, []byte(err.Error()))
		}

		client.GetConn().WriteJSON(models.BaseWebsocketMessage{
			ClientID:   client.GetID(),
			StrMessage: "ok",
		})

	case models.TableEventStealCard:
		// Run action

	default:
		return
	}
}
