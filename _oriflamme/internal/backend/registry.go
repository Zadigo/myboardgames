package backend

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

// The GameRegistry is responsible for managing the state of all active games.
// It maintains a mapping of game IDs to their corresponding game states, which are
// stored in Redis. The registry provides methods for creating nesw games, retrieving existing
// games, and updating game states as players take actions.
// TODO: Rename to GameRoom
type GameRegistry struct {
	Uuid                string                      `json:"uuid"`
	Clients             map[string]*WebsocketClient `json:"clients"`
	InfluenceQueueLayer *InfluenceQueue             `json:"influenceQueueLayer"`
	IsRunning           bool                        `json:"isRunning"`
	IsStarted           bool                        `json:"isStarted"`
	StartedAt           time.Time                   `json:"startedAt"`
	// CardsInPlay represents all the cards that were selected by the players
	// to be played in the current game
	CardsInPlay []CardInterface `json:"cardsInPlay"`
	mu          sync.RWMutex    `json:"-"`
	// The broadcast channel is used to send messages to all clients in the game.
	// When a message is sent to the broadcast channel, it will be received by all clients
	// in the game and can be used to update their game state accordingly.
	broadcast  chan WebsocketMessage `json:"-"`
	register   chan *WebsocketClient `json:"-"`
	unregister chan *WebsocketClient `json:"-"`
}

// Have a player join the game. This function adds the player to the game's client list and
// initializes any necessary game state for the new player.
func (r *GameRegistry) JoinTable(client *WebsocketClient) {
	r.mu.Lock()
	r.Clients[client.Uuid] = client
	r.mu.Unlock()

	r.register <- client
}

// Starts the main loop for the game room, which listens
// for client registration, unregistration, and broadcast messages.
func (r *GameRegistry) ListenToRoomChannels() {
	for {
		select {

		case client := <-r.register:
			r.Clients[client.Uuid] = client

		case client := <-r.unregister:
			delete(r.Clients, client.Uuid)

		case msg := <-r.broadcast:
			for _, client := range r.Clients {
				err := client.SendJsonMessage(msg)

				if err != nil {
					delete(r.Clients, client.Uuid)
				}
			}
		default:
			if !r.IsRunning {
				close(r.register)
				close(r.unregister)
				close(r.broadcast)
				return
			}
		}
	}
}

func (r *GameRegistry) SubscribeToRedisChannel(redisClient *redis.Client) {
	// Subscribe the client to the Redis channel
	// for the game they are in (if applicable)
	subcription := redisClient.Subscribe(context.Background(), r.RedisRoomId())
	redisChannel := subcription.Channel()

	for message := range redisChannel {
		var wsMsg WebsocketMessage
		json.Unmarshal([]byte(message.Payload), &wsMsg)

		// 👇 fan out to local clients
		r.broadcast <- wsMsg
		// for _, client := range r.Clients {
		// 	client.send <- wsMsg
		// }
	}
}

func (r *GameRegistry) RedisRoomId() string {
	return fmt.Sprintf("room:%s", r.Uuid)
}

func (r *GameRegistry) PublishToRoom(redisClient *redis.Client, message WebsocketMessage) error {
	payload, _ := json.Marshal(message)
	return redisClient.Publish(context.Background(), r.RedisRoomId(), payload).Err()
}

// Indicates that the game with the given UUID has started.
func (r *GameRegistry) StartGame(gameUuid string) error {
	colors := []string{"red", "blue", "green", "yellow", "purple"}

	if len(r.Clients) > len(colors) {
		return errors.New("There are too many players for this game")
	}

	rand.Shuffle(len(colors), func(x int, y int) {
		colors[x], colors[y] = colors[y], colors[x]
	})

	cards := CreateBaseCards()

	// For each player, create the cards of their color and add
	// them to the game registry's CardsInPlay.
	addToCards := func(color string, owner *WebsocketClient) {
		for _, card := range cards {
			card.GetBaseCard().Owner = owner
			card.GetBaseCard().Color = color
			r.CardsInPlay = append(r.CardsInPlay, card)
		}
	}

	for _, client := range r.Clients {
		color := slices.Delete(colors, 1, 1)

		switch color[0] {
		case "red":
			addToCards("red", client)
		case "blue":
			addToCards("blue", client)
		case "green":
			addToCards("green", client)
		case "yellow":
			addToCards("yellow", client)
		case "purple":
			addToCards("purple", client)
		}
	}

	// r.broadcast <- WebsocketMessage{
	// 	Action:      "start_game",
	// 	CardsInPlay: r.CardsInPlay,
	// }

	return nil
}

func (r *GameRegistry) EndGame(gameUuid string) {
	r.broadcast <- WebsocketMessage{
		Action: "end_game"}
}

// GetCardByUuid retrieves a card from the game registry's CardsInPlay based on its UUID.
func (r *GameRegistry) GetCardByUuid(cardUuid string) (CardInterface, error) {
	for _, card := range r.CardsInPlay {
		if card.GetBaseCard().Uuid == cardUuid {
			return card, nil
		}
	}
	return nil, errors.New("Card not found")
}

// The ServerRegistry is responsible for managing the state of the server,
// including the Redis client and any other global resources that need to be
// shared across games. It provides a centralized place to access these resources
// and ensures that they are properly initialized and managed throughout
// the lifecycle of the server.
type ServerRegistry struct {
	// All the games being played on the server, mapped by their unique identifiers.
	// This allows the server to manage multiple games simultaneously and route player
	// actions to the correct game state.
	Tables map[string]*GameRegistry `json:"tables"`
	// A mapping of all connected clients to their corresponding websocket connections.
	// This can be used to send messages to specific clients or broadcast messages to all clients.
	clients map[string]*WebsocketClient
	// The Redis client used to store game states and other shared data. This allows the
	// server to persist game states and share data across different instances of the server if needed.
	redisClient *redis.Client
	// A mutex to protect access to the clients map and other shared resources in the registry.
	mu sync.Mutex
}

// Add a new client to the server registry and return its unique identifier.
func (s *ServerRegistry) AddClient(conn *websocket.Conn) (uuid string, client *WebsocketClient) {
	s.mu.Lock()
	defer s.mu.Unlock()

	newClient := NewWebsocketClient(conn)
	s.clients[newClient.Uuid] = newClient

	return newClient.Uuid, newClient
}

func (s *ServerRegistry) GetClient(clientUuid string) (conn *WebsocketClient, exists bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	client, exists := s.clients[clientUuid]
	return client, exists
}

// Create a new game and add it to the registry. The playerUuid is used to identify the player who initiated the game.
// The function returns the newly created game registry and a boolean indicating whether the game was successfully created.
// The broadcast channel of the game registry is used to send messages to all players in the game.
func (s *ServerRegistry) CreateGame(client *WebsocketClient) (*GameRegistry, bool) {
	client, exists := s.GetClient(client.Uuid)

	if !exists {
		return nil, false
	}

	newGame := NewGameRegistry()

	s.mu.Lock()
	s.Tables[newGame.Uuid] = newGame
	s.mu.Unlock()

	// Create both local and Redis channels
	// Start listening to the local broadcast channel for the game
	//  and fan out messages to clients in the game
	go newGame.ListenToRoomChannels()
	go newGame.SubscribeToRedisChannel(s.redisClient)

	client.Initiator = true
	newGame.JoinTable(client)

	return newGame, true
}

func (s *ServerRegistry) GetGame(tableUuid string) (*GameRegistry, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	game, exists := s.Tables[tableUuid]
	if !exists {
		return nil, errors.New("Game/Room not found")
	}

	return game, nil
}

func NewServerRegistry(redisClient *redis.Client) *ServerRegistry {
	return &ServerRegistry{
		Tables:      make(map[string]*GameRegistry),
		clients:     make(map[string]*WebsocketClient),
		redisClient: redisClient,
	}
}

func NewGameRegistry() *GameRegistry {
	influenceQueueLayer := NewInfluenceQueue()

	return &GameRegistry{
		Uuid:                uuid.NewString(),
		broadcast:           make(chan WebsocketMessage),
		register:            make(chan *WebsocketClient),
		unregister:          make(chan *WebsocketClient),
		Clients:             make(map[string]*WebsocketClient),
		IsRunning:           true,
		IsStarted:           false,
		StartedAt:           time.Now(),
		InfluenceQueueLayer: influenceQueueLayer,
	}
}
