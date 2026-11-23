package wsclient

import (
	"myboardgames/minimalgo/internal/constants"
	"myboardgames/minimalgo/internal/models"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var CustomRequestUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(request *http.Request) bool {
		origin := request.Header.Get("Origin")

		_, ok := constants.AllowedOrigins[origin]
		if !ok {
			return false
		}

		return constants.AllowedOrigins[origin]
	},
}

func WsMiddleware(conn *websocket.Conn) {
	conn.SetReadLimit(models.MaxMessageSize)
	conn.SetReadDeadline(time.Now().Add(models.PingPeriod))

	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(models.ReadTimeout))
		return nil
	})
}
