package backend

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type GameHistory struct {
	mu sync.Mutex
}

// The GameRegistry is responsible for managing the state of all active games.
// It maintains a mapping of game IDs to their corresponding game states, which are
// stored in Redis. The registry provides methods for creating new games, retrieving existing
// games, and updating game states as players take actions.
type GameRegistry struct {
	Uuid      string                      `json:"uuid"`
	Clients   map[string]*WebsocketClient `json:"clients"`
	IsStarted bool                        `json:"is_started"`
	StartedAt time.Time                   `json:"started_at"`
	mu        sync.Mutex                  `json:"-"`
}

// Have a player join the game. This function adds the player to the game's client list and
// initializes any necessary game state for the new player.
func (registry *GameRegistry) JoinTable(client *WebsocketClient) {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	if registry.Clients == nil {
		registry.Clients = make(map[string]*WebsocketClient)
	}

	registry.Clients[client.Uuid] = client
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
	// A mapping all connected clients to their corresponding websocket connections.
	// This can be used to send messages to specific clients or broadcast messages to all clients.
	clients map[string]*WebsocketClient
	// The Redis client used to store game states and other shared data. This allows the
	// server to persist game states and share data across different instances of the server if needed.
	redisClient *redis.Client
	// A mutex to protect access to the clients map and other shared resources in the registry.
	mu sync.Mutex
}

// Add a new client to the server registry and return its unique identifier.
func (s *ServerRegistry) AddClient(username string, conn *websocket.Conn) (uuid string, client *WebsocketClient) {
	s.mu.Lock()
	defer s.mu.Unlock()

	newClient := NewWebsocketClient(username, conn)
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
func (s *ServerRegistry) CreateGame(playerUuid string) (*GameRegistry, bool) {
	client, exists := s.GetClient(playerUuid)

	if !exists {
		return nil, false
	}

	newGame := NewGameRegistry()
	s.Tables[newGame.Uuid] = newGame
	newGame.JoinTable(client)
	client.Initiator = true
	return newGame, true
}

// Indicates that the game with the given UUID has started.
func (s *ServerRegistry) StartGame(gameUuid string) (state bool, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	game, exists := s.Tables[gameUuid]
	if exists {
		game.StartedAt = time.Now()
		game.IsStarted = true
		return true, nil
	} else {
		return false, errors.New("Cannot start non-existent game")
	}
}

func NewServerRegistry(redisClient *redis.Client) *ServerRegistry {
	return &ServerRegistry{
		Tables:      make(map[string]*GameRegistry),
		clients:     make(map[string]*WebsocketClient),
		redisClient: redisClient,
	}
}

func NewGameRegistry() *GameRegistry {
	return &GameRegistry{
		Uuid:      uuid.NewString(),
		IsStarted: false,
		StartedAt: time.Now(),
	}
}
