package tests

import (
	"myboardgames/minimalgame/internal/backend"
	"myboardgames/minimalgame/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadRules(t *testing.T) {
	t.Run("should load rules", func(t *testing.T) {
		rules := &utils.Rule{}
		err := utils.LoadRules("..", rules)

		assert.NoError(t, err)
		assert.IsType(t, rules, &utils.Rule{})
		assert.Equal(t, rules.Rules.StartsFirst, "random")

		t.Run("should be able to save rules", func(t *testing.T) {
			client, _ := backend.CreateRedisConnection(t.Context())
			rules.SetRedis(client)
			err := rules.Save(t.Context(), rules)
			assert.NoError(t, err)

			t.Run("should be able to load rules", func(t *testing.T) {
				b := &utils.Rule{}
				err := rules.Load(t.Context(), b)
				assert.NoError(t, err)
			})
		})
	})
}

func TestCreateRdisConnection(t *testing.T) {
	t.Run("should open connection", func(t *testing.T) {
		client, err := backend.CreateRedisConnection(t.Context())

		defer func() {
			client.Close()
		}()

		assert.NoError(t, err)
		assert.NotNil(t, client)

	})
}
