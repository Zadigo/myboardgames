package tests

import (
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableCreation(t *testing.T) {
	t.Run("should connect to websocket and receive setup", func(t *testing.T) {
		conn, s, rec, err := utils.HelperJoinTableConnection(t, "/join")
		assert.NoError(t, err)

		// Connect to the websocket and
		// receive the routing table
		data := &models.Router{}
		err = conn.ReadJSON(data)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		// Select a game (by selecting a rpc route)
		route := data.Routes[0]

		err = conn.WriteJSON(map[string]string{"action": "ping_rpc", "rpc": route.Hash, "client_id": "1234"})
		assert.NoError(t, err)

		t.Cleanup(func() {
			rec.Flush()
			conn.Close()
			s.Close()
		})
	})

	t.Run("should be able to operate on created table", func(t *testing.T) {
		// t.Cleanup(func() {
		// 	rec.Flush()
		// 	conn.Close()
		// 	s.Close()
		// })
	})
}
