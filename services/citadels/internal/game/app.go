package game

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	router *chi.Mux
}

func (a *Application) Start(ctx context.Context) {
	errCh := make(chan error)
	defer close(errCh)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8000),
		Handler: a.router,
	}

	go func() {
		errCh <- server.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		fmt.Println(err)
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		
		fmt.Print("🔴 Stopping HTTP application...")
		server.Shutdown(ctx)
	}
}

func NewHttpApplication() *Application {
	app := &Application{}
	app.loadRoutes()
	return app
}
