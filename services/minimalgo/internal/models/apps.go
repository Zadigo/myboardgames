package models

type BaseHttpApp struct {
}

type BaseServerApp struct {
	// Top variable used to track the clients that were
	// regisered in the managed HTTP server
	_clients []WebsocketClientInterface
}

func (s *BaseServerApp) UpdateClients(clients []WebsocketClientInterface) {
	s._clients = append(s._clients, clients...)
}

type ServerAppOptions struct{}
