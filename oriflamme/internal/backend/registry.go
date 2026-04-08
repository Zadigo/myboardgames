package backend

import (
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
// TODO: Rename to GameRoom
type GameRegistry struct {
	Uuid      string                      `json:"uuid"`
	Clients   map[string]*WebsocketClient `json:"clients"`
	IsStarted bool                        `json:"is_started"`
	StartedAt time.Time                   `json:"started_at"`
	mu        sync.RWMutex                `json:"-"`
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
	defer r.mu.Unlock()

	if r.Clients == nil {
		r.Clients = make(map[string]*WebsocketClient)
	}

	r.Clients[client.Uuid] = client
	r.register <- client
}

// Starts the main loop for the game room, which listens
// for client registration, unregistration, and broadcast messages.
func (r *GameRegistry) StartRoom() {
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
		}
	}
}

// Indicates that the game with the given UUID has started.
func (r *GameRegistry) StartGame(gameUuid string) {
	r.broadcast <- WebsocketMessage{
		Action: "start_game",
	}
}

func (r *GameRegistry) EndGame(gameUuid string) {
	r.broadcast <- WebsocketMessage{
		Action: "end_game",
	}
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

	// Start the game room in a new goroutine
	// to handle client registration and broadcasting
	go newGame.StartRoom()

	client.Initiator = true
	newGame.JoinTable(client)

	return newGame, true
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
		Uuid:       uuid.NewString(),
		broadcast:  make(chan WebsocketMessage),
		register:   make(chan *WebsocketClient),
		unregister: make(chan *WebsocketClient),
		IsStarted:  false,
		StartedAt:  time.Now(),
	}
}
