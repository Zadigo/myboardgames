package logic

import (
	"sync"

	"github.com/Zadigo/gosplendor/internal/models"
)

type PlayingTableInterface interface {
	AddClient(client *models.WebsocketClient) error
	RemoveClient(client *models.WebsocketClient) error
	GetClient(clientUuid string) (*models.WebsocketClient, bool)
	GetUuid() string
}

type PlayingTable struct {
	Uuid      string                             `json:"uuid"`
	Clients   map[string]*models.WebsocketClient `json:"clients"`
	broadcast chan models.WebsocketMessage
	mu        sync.Mutex
}

func (pt *PlayingTable) GetUuid() string {
	return pt.Uuid
}
