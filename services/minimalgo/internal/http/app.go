package app

import (
	"context"
	"fmt"
	"myboardgames/minimalgo/internal/models"
)

type HttpApp struct {
	models.BaseHttpApp
}

func (h *HttpApp) Start(ctx context.Context) {
	fmt.Print("Initializing HTTP server")
}

func NewHttpApp(ctx context.Context) models.AppInterface {
	return &HttpApp{}
}
