package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Zadigo/splendor/internal/logic"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type App struct {
	ctx         context.Context
	router      *chi.Mux
	redisClient *redis.Client
	gameEngine  *logic.GameEngine
}

func (a *App) Start(ctx context.Context) error {
	a.ctx = ctx

	err := a.redisClient.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("🔴 Could not connect to Redis: %w", err)
	}

	defer func() {
		a.redisClient.Close()
		log.Println("✅ Redis client closed.")
	}()

	// Engine setup
	a.gameEngine = logic.NewGameEngine(ctx)

	server := http.Server{
		Addr: ":9000",
	}

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("🔴 Could not start server %w", err)
		}
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		log.Println("⚡️ Shutting down server...")
		timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		return server.Shutdown(timeoutCtx)
	}
}

func NewApp() *App {
	app := &App{
		redisClient: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
	}
	app.loadUrls()
	return app
}
