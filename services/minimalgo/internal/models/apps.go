package models

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type BaseHttpApp struct {
}

type BaseServerApp struct {
	// Top variable used to track the clients that were
	// regisered in the managed HTTP server
	_clients []WebsocketClientInterface
}

func (s *BaseServerApp) Clients() []WebsocketClientInterface {
	return s._clients
}

func (s *BaseServerApp) UpdateClients(clients []WebsocketClientInterface) {
	s._clients = clients
	fmt.Printf("%d active clients\n", len(s._clients))
}

// ServerAppOptions is used to set main global
// options the main server application
type ServerAppOptions struct {
	Name        string
	RedisClient *redis.Client
	Port        int
}
