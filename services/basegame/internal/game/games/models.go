package games

import (
	"context"
	"log"
	"time"

	"github.com/Zadigo/basegame/internal/models"
	"github.com/go-co-op/gocron"
)

type BaseGame struct {
	ctx     context.Context
	players map[string]*WebsocketClient
	options models.GameAppOptions
}

func (b *BaseGame) SetOptions(options models.GameAppOptions) {
	b.options = options
}

func (b *BaseGame) GetOptions() models.GameAppOptions {
	return b.options
}

func (b *BaseGame) Start() error {
	scheduler := gocron.NewScheduler(time.UTC)

	if b.options.Debug {
		log.Print("🟢 Scheduler was not started in debug mode")
		return nil
	}

	// Start a task that will run every X seconds in order
	// to check if the game is still active, that the players
	// are still connected, game rules etc.
	_, err := scheduler.Every(60).Second().Do(func() {
		// Rule: When all the players are frozen,
		// the game should automatically move to the next round
		var frozenPlayerCount int
		for _, player := range b.players {
			if player.Player.IsFrozen {
				frozenPlayerCount++
			}

			if frozenPlayerCount == len(b.players) {
				// All player are frozen, get to the next round
			}
		}

		// Rule: When a player reaches a certain score, the game should
		// automatically end and declare the player as the winner
		for _, player := range b.players {
			// TODO: The max total score should be configurable in the game settings
			if player.Player.GetTotalScore() >= 200 {
				// Player has reached the winning score, end the game
			}
		}
	})

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
