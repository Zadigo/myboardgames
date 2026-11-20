package tests

import (
	"myboardgames/minimalgo/internal/models"
	"myboardgames/minimalgo/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	t.Run("should be able to join table", func(t *testing.T) {
		conn, _, rec, err := utils.HelperJoinTableConnection(t, "/join")
		assert.Equal(t, rec.Result().StatusCode, 200)
		assert.NoError(t, err)

		err = conn.WriteJSON(map[string]string{"Action": models.HelloPing})
		assert.NoError(t, err)

		_, rcvMessage, err := conn.ReadMessage()
		assert.NoError(t, err)
		assert.Equal(t, "Hello", string(rcvMessage))
	})
}
