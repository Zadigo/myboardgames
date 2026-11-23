package tests

import (
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/internal/wsclient"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DummyClient struct {
	models.ServerAppInterface
}

type DummyWsClient struct {
	models.WebsocketClientInterface
	t *testing.T
}

func (d *DummyWsClient) GetID() string {
	return "dummy"
}

func (d *DummyWsClient) WriteJSON(message models.BaseWebsocketMessage) error {
	d.t.Log(message)
	assert.Equal(d.t, message.Action, models.PingRpc)
	return nil
}

func TestMessageBroker(t *testing.T) {
	t.Run("should create new message broker without action", func(t *testing.T) {
		broker := wsclient.NewMessageBroker(&DummyClient{})
		broker.Execute(t.Context(), &DummyWsClient{}, &models.BaseWebsocketMessage{})
	})

	t.Run("should create new message ping rpc", func(t *testing.T) {
		broker := wsclient.NewMessageBroker(&DummyClient{})
		broker.Execute(t.Context(), &DummyWsClient{t: t}, &models.BaseWebsocketMessage{Action: models.PingRpc})
	})
}
