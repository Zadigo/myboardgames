package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Zadigo/flipseven3/internal/backend/logic"
	"github.com/Zadigo/flipseven3/internal/backend/redisclient"
	"github.com/Zadigo/flipseven3/internal/handlers"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func CreateConnection(t *testing.T) (*websocket.Conn, *httptest.Server) {
	t.Helper()

	redisClient := redisclient.NewRedisClient("redis://@localhost:6379/0")
	sr := logic.NewServerRegistry(redisClient)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Origin", "http://localhost:3000")
		handlers.LiveGameHandler(w, r, sr)
	})

	server := httptest.NewServer(handler)

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/oriflamme/live"
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

	assert.NotNil(t, conn)
	assert.NoError(t, err)

	time.Sleep(50 * time.Millisecond)

	return conn, server
}

func CreateTestCards(t *testing.T) []*logic.Card {
	t.Helper()

	cards := []*logic.Card{}

	cards = append(cards, logic.GetNumberCards()...)
	return cards
}

func CreateCardQueue(t *testing.T) *logic.CardQueue {
	t.Helper()

	queue := logic.NewCardQueue()
	queue.Queue = CreateTestCards(t)
	return queue
}
