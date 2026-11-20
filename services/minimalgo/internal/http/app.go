package http

import (
	"context"
	"fmt"
	"log/slog"
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

func (h *HttpApp) SetServer(s models.ServerAppInterface) {
	h.server = s
}

func (h *HttpApp) AddClient(client models.WebsocketClientInterface) error {
	h.clients = append(h.clients, client)
	if h.server == nil {
		return fmt.Errorf("No server was set on the handler")
	}
	h.server.UpdateClients(h.clients)
	return nil
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
	slog.Info("Initializing HTTP server")

	server := &http.Server{Addr: ":9000", Handler: h.router}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func NewHttpApp(ctx context.Context, server models.ServerAppInterface) models.AppInterface {
	app := &HttpApp{
		clients: []models.WebsocketClientInterface{},
		server:  server,
	}
	app.loadRoutes()
	return app
}
