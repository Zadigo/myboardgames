package backend

import (
	"fmt"
	"sync"
)

type ServerRegistry struct {
	// Clients holds the active websocket clients connected to the server.
	Clients    map[string]WebsocketClientInterface[WebsocketMessage]
	broadcast  chan WebsocketMessage
	register   chan WebsocketClientInterface[WebsocketMessage]
	unregister chan string
	Mu         *sync.RWMutex
}

type ServerRegistryInterface interface {
	AddClient(client WebsocketClientInterface[WebsocketMessage]) error
	RemoveClient(clientUuid string) error
	GetClient(clientUuid string) (WebsocketClientInterface[WebsocketMessage], bool)
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

func (sr *ServerRegistry) RemoveClient(clientUuid string) error {
	sr.Mu.Lock()
	defer sr.Mu.Unlock()

	_, exists := sr.Clients[clientUuid]
	if !exists {
		return fmt.Errorf("Client with UUID %s does not exist", clientUuid)
	}

	delete(sr.Clients, clientUuid)
	return nil
}

func (sr *ServerRegistry) GetClient(clientUuid string) (WebsocketClientInterface[WebsocketMessage], bool) {
	sr.Mu.RLock()
	defer sr.Mu.RUnlock()

	client, exists := sr.Clients[clientUuid]
	return client, exists
}

func NewServerRegistry() ServerRegistryInterface {
	return &ServerRegistry{
		Clients: make(map[string]WebsocketClientInterface[WebsocketMessage]),
	}
}
