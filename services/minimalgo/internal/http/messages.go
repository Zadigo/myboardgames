package http

import (
	"context"
	pb "myboardgames/gen/minimalgame/v1"
	"myboardgames/minimalgo/internal/grpcclient"
	"myboardgames/minimalgo/internal/models"
)

// func authenticationMessage(app *HttpApp, client models.WebsocketClientInterface, message models.BaseWebsocketMessage) {
// 	switch message.Action {
// 	case models.PingRpc:
// 		rpc, ok := app.server.GetGrpcConnections()[message.Rpc]
// 		if !ok {
// 			client.GetConn().WriteJSON(models.BaseWebsocketMessage{
// 				ClientID:   client.GetID(),
// 				StrMessage: "rpc not found",
// 			})
// 			return
// 		} else {
// 			rpc.Connect()
// 			client.GetConn().WriteJSON(models.BaseWebsocketMessage{
// 				ClientID:   client.GetID(),
// 				StrMessage: "ok",
// 			})
// 		}
// 	default:
// 		return
// 	}
// }

// func globalGameMessage(app *HttpApp, client models.WebsocketClientInterface, message models.BaseWebsocketMessage) {
// 	// handle global game messages here
// }

// // gameRelatedMessage must implement all the related websocket logic to interact with the
// // logic of the game in the backend
// func gameRelatedMessage(app *HttpApp, client models.WebsocketClientInterface, message models.BaseWebsocketMessage) {
// 	// handle game messages here
// 	switch message.Action {
// 	case models.CreateTable:
// 		// Run action
// 		req := &pb.CreateTableRequest{ClientId: client.GetID()}

// 		conns := app.server.GetGrpcConnections()
// 		conn, _ := conns[message.Rpc]

// 		err := conn.Invoke(context.Background(), pb.TableService_CreateTable_FullMethodName, req, nil)
// 		if err != nil {
// 			client.GetConn().WriteMessage(1, []byte(err.Error()))
// 		}

// 		client.GetConn().WriteJSON(models.BaseWebsocketMessage{
// 			ClientID:   client.GetID(),
// 			StrMessage: "ok",
// 		})

// 	case models.TableEventStealCard:
// 		// Run action

// 	default:
// 		return
// 	}
// }

type MessagesInterface interface {
	execute(ctx context.Context, client models.WebsocketClientInterface, message *models.BaseWebsocketMessage)
	setNext(MessagesInterface)
}

type BaseMessageBroker struct {
	app  models.AppInterface
	next MessagesInterface
}

func (b *BaseMessageBroker) setNext(next MessagesInterface) {
	b.next = next
}

// Create a base message for client which always includes the client iD field
func (b *BaseMessageBroker) createBaseMessage(client models.WebsocketClientInterface) models.BaseWebsocketMessage {
	return models.BaseWebsocketMessage{ClientID: client.GetID()}
}

type AuthenticationMessage struct {
	BaseMessageBroker
}

func (a *AuthenticationMessage) execute(ctx context.Context, client models.WebsocketClientInterface, message *models.BaseWebsocketMessage) {
	switch message.Action {
	case models.PingRpc:
		// Create a goroutine to not block
		// the read loop
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), models.WriteTimeout)
			defer cancel()

			ctx = grpcclient.WithPlayer(ctx, client.GetID())

			// err := client.WriteJSON(models.BaseWebsocketMessage{
			// 	Action:   message.Action,
			// 	ClientID: client.GetID(),
			// 	Rpc:      msg.Rpc,
			// 	TableID:  resp.GetEvent().GetTableId(),
			// })

			// tryLogWriteError(client, err)
		}()
	case models.MustIdentify:
		_ = &pb.CreateTableRequest{}
		// handle ping
	default:
		a.next.execute(ctx, client, message)
	}
}

type GameRelatedMessage struct {
	BaseMessageBroker
}

func (g *GameRelatedMessage) execute(ctx context.Context, client models.WebsocketClientInterface, message *models.BaseWebsocketMessage) {
	switch message.Action {
	default:
		g.next.execute(ctx, client, message)
	}
}

// NewMessageBroker creates the chain of message brokers
// for handling incoming messages
func NewMessageBroker(app *HttpApp) MessagesInterface {
	auth := &AuthenticationMessage{app: app}

	game := &GameRelatedMessage{}
	auth.setNext(game)

	return auth
}
