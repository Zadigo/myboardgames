package http

import (
	"fmt"
	"myboardgames/minimalgo/internal/models"
)

func authenticationMessage(client models.WebsocketClientInterface, message models.BaseWebsocketMessage) {
	switch message.Action {
	case models.MustIdentify:
		fmt.Println(message)
	}
}

func gameRelatedMessage(client models.WebsocketClientInterface, message models.BaseWebsocketMessage) {

}
