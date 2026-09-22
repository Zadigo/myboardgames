package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Zadigo/goflipseven/internal/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := server.NewApp()
	app.Start(ctx)

	// serverRegistry := backend.NewServerRegistry()
	// router := chi.NewRouter()

	// router.Get("/wsgosplendorlive", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.LiveGameHandler(w, r, serverRegistry)
	// })

	// http.ListenAndServe(":9000", router)
}
