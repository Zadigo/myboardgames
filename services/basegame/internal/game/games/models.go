package games

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
)

type BaseGame struct {
	ctx     context.Context
	players map[string]*WebsocketClient
}

func (b *BaseGame) Start() error {
	scheduler := gocron.NewScheduler(time.UTC)

	// Start a taks that will run every X seconds in order
	// to check if the game is still active, that the players
	// are still connected, etc.
	_, err := scheduler.Every(60).Second().Do(func() {})

	if err != nil {
		// Handle error
		return err
	}

	scheduler.StartBlocking()
	return nil
}

func (b *BaseGame) CreateCards() error {
	return nil
}

func (b *BaseGame) Prepare() error {
	return nil
}

func (b *BaseGame) AddPlayer(player *WebsocketClient) error {
	return nil
}

func (b *BaseGame) RemovePlayer(player *WebsocketClient) error {
	return nil
}

func (b *BaseGame) GetPlayer(playerUuid string) (*WebsocketClient, error) {
	return nil, nil
}
