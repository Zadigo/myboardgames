package http

import (
	"log"
	"myboardgames/minimalgo/internal/constants"
	"myboardgames/minimalgo/internal/models"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func Cors(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		origin := r.Header.Get("Origin")

		if _, ok := constants.AllowedOrigins[origin]; !ok {
			log.Println("❌ Origin not allowed:", origin)
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

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
