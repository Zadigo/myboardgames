package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Zadigo/gosplendor/internal/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := server.NewApp()
	app.Start(ctx)

	// serverRegistry := backend.NewServerRegistry()
	// router := chi.NewRouter()

	// router.Get("/ws/splendor/live", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.LiveGameHandler(w, r, serverRegistry)
	// })

	// http.ListenAndServe(":9000", router)
}
