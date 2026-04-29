package backend

import (
	"fmt"
	"sync"
)

type ServerRegistry struct {
	// Clients holds the active websocket clients connected to the server.
	Clients    map[string]WebsocketClientInterface[WebsocketMessage]
	Games      map[string]*PlayingTable
	broadcast  chan WebsocketMessage
	register   chan WebsocketClientInterface[WebsocketMessage]
	unregister chan string
	Mu         *sync.RWMutex
}

type ServerRegistryInterface interface {
	AddClient(client WebsocketClientInterface[WebsocketMessage]) error
	RemoveClient(client WebsocketClientInterface[WebsocketMessage]) error
	GetClient(clientUuid string) (WebsocketClientInterface[WebsocketMessage], bool)
	AddPlayingTable(table *PlayingTable) error
}

func (sr *ServerRegistry) AddClient(client WebsocketClientInterface[WebsocketMessage]) error {
	sr.Mu.Lock()
	defer sr.Mu.Unlock()

	player := client.GetPlayer()
	_, exists := sr.Clients[player.Uuid]

	if exists {
		return fmt.Errorf("Client with UUID %s already exists", player.Uuid)
	}

	sr.Clients[player.Uuid] = client
	return nil
}

func (sr *ServerRegistry) RemoveClient(client WebsocketClientInterface[WebsocketMessage]) error {
	sr.Mu.Lock()
	defer sr.Mu.Unlock()

	player := client.GetPlayer()
	_, exists := sr.Clients[player.Uuid]
	if !exists {
		return fmt.Errorf("Client with UUID %s does not exist", player.Uuid)
	}

	delete(sr.Clients, player.Uuid)
	return nil
}

func (sr *ServerRegistry) GetClient(clientUuid string) (WebsocketClientInterface[WebsocketMessage], bool) {
	sr.Mu.RLock()
	defer sr.Mu.RUnlock()

	client, exists := sr.Clients[clientUuid]
	return client, exists
}

func (sr *ServerRegistry) AddPlayingTable(table *PlayingTable) error {
	sr.Mu.Lock()
	defer sr.Mu.Unlock()

	_, exists := sr.Games[table.Uuid]
	if exists {
		return fmt.Errorf("Table with UUID %s already exists", table.Uuid)
	}

	sr.Games[table.Uuid] = table
	return nil
}

func NewServerRegistry() ServerRegistryInterface {
	return &ServerRegistry{
		Clients:    make(map[string]WebsocketClientInterface[WebsocketMessage]),
		Games:      make(map[string]*PlayingTable),
		broadcast:  make(chan WebsocketMessage),
		register:   make(chan WebsocketClientInterface[WebsocketMessage]),
		unregister: make(chan string),
		Mu:         &sync.RWMutex{},
	}
}
