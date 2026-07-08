package game

import (
	"context"
	"fmt"
	"time"

	"github.com/Zadigo/basegame/internal/game/cards"
	"github.com/Zadigo/basegame/internal/game/games"
	"github.com/Zadigo/basegame/internal/models"
)

type GameApp struct {
	ctx               context.Context
	Game              games.GameInterface               `json:"game"`
	CurrentRound      int                               `json:"currentRound"`
	CardPile          *cards.CardPile                   `json:"cardPile"`
	Players           map[string]*games.WebsocketClient `json:"players"`
	CurrentPlayerUuid string                            `json:"currentPlayerIndex"`
	IsRunning         bool                              `json:"isRunning"`
	// An action or an event must be resolved before the game can continue.
	// This is useful for special cards that may have more complex effects
	MustResolve bool `json:"mustResolve"`
	// Options for resolving an event or an action that must be resolved
	// before the game can continue
	EventResolutionOptions EventResolutionOptions `json:"-"`
	// Reference to the main server application that
	// manages the lifecycle of the game server and
	// other services
	serverApp models.ServerAppInterface
	// ConnectionCounter int
	Scores map[string]int `json:"scores"`
}

// Build initializes and starts the game server application.
// It checks if the game instance is properly initialized, prepares the game state,
// and then starts the game logic.
func (app *GameApp) Start() error {
	if app.Game == nil {
		return fmt.Errorf("❌ Game server is not initialized")
	}

	app.Game.Prepare()
	app.Game.Start()

	app.IsRunning = true
	return nil
}

func (app *GameApp) Stop() error {
	if app.Game == nil {
		return fmt.Errorf("❌ Game server is not initialized")
	}

	timeoutCtx, cancel := context.WithTimeout(app.ctx, 5000*time.Millisecond)
	defer cancel()

	app.IsRunning = false
	timeoutCtx.Done()

	return nil
}

func (app *GameApp) GetCurrentPlayer() (client *games.WebsocketClient, err error) {
	if app.CurrentPlayerUuid == "" {
		return nil, fmt.Errorf("❌ Current player UUID is not set")
	}

	player, exists := app.Players[app.CurrentPlayerUuid]
	if !exists {
		return nil, fmt.Errorf("❌ Player not found")
	}
	return player, nil
}

func (app *GameApp) SetCurrentPlayer(clientUuid string) error {
	if clientUuid == "" {
		return fmt.Errorf("❌ Player UUID is empty")
	}

	player, exists := app.Players[clientUuid]
	if !exists {
		return fmt.Errorf("❌ Player not found")
	}

	app.CurrentPlayerUuid = player.GetUuid()
	return nil
}

func (app *GameApp) AddPlayer(client *games.WebsocketClient) error {
	if client == nil {
		return fmt.Errorf("❌ Player is nil")
	}

	if client.GetConn() == nil {
		return fmt.Errorf("❌ Websocket connection is nil")
	}

	if _, exists := app.Players[client.GetUuid()]; exists {
		return fmt.Errorf("❌ Player already exists")
	}

	app.Players[client.GetUuid()] = client

	return nil
}

// DrawCardFromPile draws a card from the card pile and adds it to the current player's hand.
// It also checks for special cards and handles the game logic accordingly.
func (app *GameApp) DrawCardFromPile() error {
	newCard := app.CardPile.DrawCard()
	if newCard == nil {
		return fmt.Errorf("❌ No cards left in the pile")
	}

	player, err := app.GetCurrentPlayer()
	if err != nil {
		return err
	}

	// Freez and Flip 3 cards must be resolved immediately. In other words, the
	// user has to decide what he wants to do with the card since it can affect scores
	if newCard.IsSpecial() && (newCard.IsCardOperation(cards.OPERATION_FREEZE) || newCard.IsCardOperation(cards.OPERATION_FLIP3)) {
		app.MustResolve = true
		app.EventResolutionOptions = EventResolutionOptions{
			Card:   newCard,
			Player: player,
		}
		return nil
	}

	// Add the drawn card to the player's hand
	player.Player.Cards = append(player.Player.Cards, newCard)

	// Additional logic to handler
	// once the function is complete
	defer func() {
		if app.CardPile.RemainingCards() == 0 {
			// Reshuffle the card pile when all cards have been drawn
			return
		}

		// Recalculate the points for each player
		// after a card has been drawn
		for _, player := range app.Players {
			totalScore := player.Player.GetTotalScore()
			app.Scores[player.GetUuid()] = totalScore
		}
	}()

	// Map the other types of special cards that can be resolved
	// later on in the game (e.g. second chance)
	otherSpecialCards := map[string]bool{}
	for _, cardInHand := range player.Player.GetCards() {
		if cardInHand.IsSpecial() {
			otherSpecialCards[cardInHand.GetOperator()] = true
		}
	}

	for range player.Player.GetCards() {
		if player.Player.GetNumberOfCards() > 1 {
			// A player cannot draw a card if they already have a card
			// of the same value in their hand. Their turn stops and
			// they loose all the points of all the cards that they
			// currently have in their hand
			values := make(map[int]bool)
			for _, card := range player.Player.GetCards() {
				if values[card.GetValue()] {
					if otherSpecialCards[cards.OPERATION_SECOND_CHANCE] {
						// Remove the second chance card from the player's hand
						// and allow them to continue their turn -; also remove
						// the last card that was drawn from the player's hand since
						// it has the same value as another card in their hand
						for i, c := range player.Player.GetCards() {
							if c.GetOperator() == cards.OPERATION_SECOND_CHANCE {
								player.Player.Cards = append(player.Player.Cards[:i], player.Player.Cards[i+1:]...)
								break
							}
						}
						player.Player.Cards = player.Player.Cards[:len(player.Player.Cards)-1]
						return nil
					}

					// Reset the player's hand and end their turn
					player.Player.Cards = []cards.CardInterface{}
					player.Player.IsFrozen = true
					return nil
				}
				values[card.GetValue()] = true
			}
		}
	}

	return nil
}

func (app *GameApp) ResolveEvent(action EventResolutionAction) {
	if app.MustResolve {
		// Resolve a pending event
	}
}

func (app *GameApp) NextRound() {

}

// Create a new game instance based on the specified game type. This method is a placeholder
// and should be implemented to return a new game instance that adheres to the GameInterface.
func (app *GameApp) Create(options models.GameAppOptions) games.GameInterface {
	game := games.CreateGame(options)
	return game
}

func (app *GameApp) NotifyAll() {
	// TODO: This is a simple implementation but for production, create a
	// loop that sends messages to all players when a message is sent
	// to a channel. This will allow for better performance and scalability.
	for _, player := range app.Players {
		player.GetConn().WriteJSON(models.WebsocketMessage{})
	}
}

func (app *GameApp) GetGameState() bool {
	return app.IsRunning
}

// NewGameApp creates a new instance of the game server application
// based on the specified game type (standard or extension). This server
// is independent of the main chi server and can be used to manage game sessions,
// handle player interactions, and maintain game state.
func NewGameApp(ctx context.Context, options models.GameAppOptions) *GameApp {
	// FIXME: Create a map that identifies each game individually and
	// allows for multiple games to be played simultaneously. This will
	// require a unique identifier for each game session and a way to manage
	// the state of each game independently.
	game := games.CreateGame(options)
	return &GameApp{
		ctx:                    ctx,
		Game:                   game, // Rename to Games which will be a map of game instances to allow for multiple games to be played simultaneously
		CurrentRound:           0,
		CardPile:               &cards.CardPile{},
		CurrentPlayerUuid:      "",
		IsRunning:              false,
		MustResolve:            false,
		EventResolutionOptions: EventResolutionOptions{},
		serverApp:              options.ServerApp,
		Players:                map[string]*games.WebsocketClient{},
	}
}
