package server

import (
	"context"
	"fmt"
	"myboardgames/minimalgo/internal/games"
	httpapp "myboardgames/minimalgo/internal/http"
	"myboardgames/minimalgo/internal/models"
	"time"
)

type ServerApp struct {
	models.BaseServerApp
}

func (s *ServerApp) Start(ctx context.Context) {
	chError := make(chan error)

	serverCtx := context.WithValue(ctx, "state", true)

	fmt.Println("Initializing main server")

	http := httpapp.NewHttpApp(ctx, s)
	go func(ctx context.Context) {
		http.Start(ctx)
	}(ctx)

	gameApp := games.NewGameServerApp(ctx)
	go func(ctx context.Context) {
		gameApp.Start(ctx)
	}(ctx)

	select {
	case err := <-chError:
		fmt.Println(err)
	case <-serverCtx.Done():
		timeoutCtx, cancel := context.WithTimeout(ctx, 10000*time.Second)
		defer cancel()

		fmt.Printf("server stopped %v", timeoutCtx.Err())
	}
}

func NewServerApp(ctx context.Context, options *models.ServerAppOptions) models.ServerAppInterface {
	return &ServerApp{}
}
