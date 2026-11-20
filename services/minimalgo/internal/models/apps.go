package models

import "fmt"

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

type ServerAppOptions struct{}
