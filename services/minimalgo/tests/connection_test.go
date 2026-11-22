package tests

import (
	"myboardgames/minimalgo/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableCreation(t *testing.T) {
	t.Run("should be able to create table", func(t *testing.T) {
		conn, s, rec, err := utils.HelperJoinTableConnection(t, "/join")
		assert.NoError(t, err)

		err = conn.WriteJSON(map[string]string{"client_id": "1234"})
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
