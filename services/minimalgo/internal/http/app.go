package http

import (
	"context"
	"fmt"
	"myboardgames/minimalgo/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HttpApp struct {
	models.BaseHttpApp
	// Keeps track of all the active clients
	// that joined the website
	clients []models.WebsocketClientInterface
	server  models.ServerAppInterface
	router  *chi.Mux
}

func (h *HttpApp) AddClient(client models.WebsocketClientInterface) {
	h.clients = append(h.clients, client)
	h.server.UpdateClients(h.clients)
}

func (h *HttpApp) RemoveClient(client models.WebsocketClientInterface) {
	for i, c := range h.clients {
		if c == client {
			h.clients = append(h.clients[:i], h.clients[i+1:]...)
			break
		}
	}
	h.server.UpdateClients(h.clients)
}

func (h *HttpApp) Start(ctx context.Context) {
	fmt.Println("Initializing HTTP server")
	server := &http.Server{Addr: ":9000", Handler: h.router}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func NewHttpApp(ctx context.Context, server models.ServerAppInterface) models.AppInterface {
	app := &HttpApp{server: server}
	app.loadRoutes()
	return app
}
