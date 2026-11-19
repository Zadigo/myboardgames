package models

import "context"

type AppInterface interface {
	Start(ctx context.Context)
}

type WebsocketClientInterface interface {
	IsActive() bool
}
