package flipseven3

import (
	"net/http"

	"github.com/Zadigo/flipseven3/internal/backend/logic"
	"github.com/Zadigo/flipseven3/internal/backend/redisclient"
	"github.com/Zadigo/flipseven3/internal/handlers"
)

func main() {
	redisClient := redisclient.NewRedisClient("redis://@localhost:5679/0")
	serverRegistry := logic.NewServerRegistry(redisClient)

	http.HandleFunc("/flip7/live", func(w http.ResponseWriter, r *http.Request) {
		handlers.LiveGameHandler(w, r, serverRegistry)
	})

	http.ListenAndServe(":8080", nil)
}
