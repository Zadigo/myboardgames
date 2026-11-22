package utils

import (
	"io"
	"myboardgames/minimalgo/internal/http"
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/internal/server"
	httpnet "net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func HelperJoinTableConnection(t *testing.T, path string) (*websocket.Conn, *httptest.Server, *httptest.ResponseRecorder, error) {
	t.Helper()

	httpApp := &http.HttpApp{}
	serverApp, err := server.NewServerApp(t.Context(), &models.ServerAppOptions{})
	if err != nil {
		return nil, nil, nil, err
	}

	httpApp.SetServer(serverApp)

	h := &http.BaseHandlers{}
	h.SetApp(httpApp)

	testingServer := httptest.NewServer(httpnet.HandlerFunc(h.JoinWebsite))

	recorder := httptest.NewRecorder()
	wsURL := "ws" + strings.TrimPrefix(testingServer.URL, "http") + path
	conn, response, err := websocket.DefaultDialer.Dial(wsURL, httpnet.Header{"Origin": []string{"GoTest"}})

	if err != nil {
		if response != nil {
			body, _ := io.ReadAll(response.Body)
			response.Body.Close()
			t.Fatalf("websocket dial failed (status %d): %v | body: %s", response.StatusCode, err, string(body))
		}
		t.Fatalf("websocket dial failed: %v", err)
	}

	return conn, testingServer, recorder, err
}
