package tests

import (
	"testing"

	"github.com/Zadigo/flipseven2/internal/models"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestPlayersTable(t *testing.T) {
	type testCase struct {
		table                 *models.PlayersTable
		expected              string
		expectedNumberOfCards int
		error                 error
	}

	testCases := []testCase{
		{
			table:                 models.CreatePlayersTable(),
			expected:              "",
			expectedNumberOfCards: 88,
			error:                 nil,
		},
	}

	for _, tc := range testCases {
		t.Run("Table integrity", func(t *testing.T) {
			assert.Negative(t, tc.table.Layer.PrintDetails().DeckIndex)
			assert.Empty(t, tc.table.Layer.PrintDetails().Deck)
			assert.Empty(t, tc.table.Layer.PrintDetails().Players)
			assert.Negative(t, tc.table.Layer.PrintDetails().CurrentPlayerIndex)
			assert.False(t, tc.table.Layer.PrintDetails().IsStarted)
		})

		t.Run("Function PrintDetails", func(t *testing.T) {
			result := tc.table.Layer.PrintDetails()
			assert.NotEqual(t, result.Uuid, "")
			assert.Empty(t, result.Deck)
			assert.Empty(t, result.Players)
			assert.Negative(t, result.CurrentPlayerIndex)
			assert.False(t, result.IsStarted)
		})

		t.Run("Function GetUuid", func(t *testing.T) {
			result := tc.table.Layer.GetUuid()
			assert.NotEqual(t, result, "")
		})

		t.Run("Function GetCurrentPlayer", func(t *testing.T) {
			result := tc.table.Layer.GetCurrentPlayer()
			assert.Nil(t, result)
		})

		t.Run("Function GetCurrentCard", func(t *testing.T) {
			result, err := tc.table.Layer.GetCurrentCard()
			assert.Nil(t, result)
			assert.Error(t, err)
		})

		t.Run("Function FlipCard", func(t *testing.T) {
			result, card, err := tc.table.Layer.FlipCard(nil)
			assert.Equal(t, "", result)
			assert.Nil(t, card)
			assert.Error(t, err)
			assert.ErrorContains(t, err, "Game has not started yet")
		})

		t.Run("Function StartGame", func(t *testing.T) {
			tc.table.Layer.StartGame()
			assert.True(t, tc.table.Layer.PrintDetails().IsStarted)
			assert.Negative(t, tc.table.Layer.PrintDetails().DeckIndex)
			assert.GreaterOrEqual(t, tc.table.Layer.PrintDetails().CurrentPlayerIndex, 0)
		})

		t.Run("Function FreezePlayer", func(t *testing.T) {})

		t.Run("Function ResetPlayers", func(t *testing.T) {})

		t.Run("Function GetDeck", func(t *testing.T) {
			newTable := models.CreatePlayersTable()
			deck := newTable.Layer.GetDeck()
			assert.NotEmpty(t, deck)
			assert.Len(t, deck, tc.expectedNumberOfCards)
			assert.NotEmpty(t, newTable.Layer.PrintDetails().Deck)
		})

		t.Run("Function NumberOfCards", func(t *testing.T) {
			tc.table.Layer.GetDeck()
			result := tc.table.Layer.NumberOfCards()
			assert.Equal(t, tc.expectedNumberOfCards, result)
		})

		t.Run("Function AddPlayer", func(t *testing.T) {
			conn := &websocket.Conn{}
			result := tc.table.Layer.AddPlayer(conn, "pauline", false)
			assert.NotNil(t, result)
			assert.Equal(t, "pauline", result.Username)
			assert.Equal(t, conn, result.Conn)
		})

		t.Run("Function AddPlayer & start game", func(t *testing.T) {
			conn := &websocket.Conn{}
			tc.table.Layer.AddPlayer(conn, "pauline", false)
			tc.table.Layer.AddPlayer(conn, "pauline92", false)
			tc.table.Layer.AddPlayer(conn, "pauline88", false)
			tc.table.Layer.StartGame()
			assert.True(t, tc.table.Layer.PrintDetails().IsStarted)
			assert.GreaterOrEqual(t, tc.table.Layer.PrintDetails().CurrentPlayerIndex, 0)
			assert.Less(t, tc.table.Layer.PrintDetails().CurrentPlayerIndex, tc.table.Layer.GetNumberOfPlayers())
		})

		t.Run("Function HasPlayer", func(t *testing.T) {
			conn := &websocket.Conn{}
			tc.table.Layer.AddPlayer(conn, "pauline", false)
			assert.True(t, tc.table.Layer.HasPlayer(conn))
		})

		t.Run("Function GetPlayer", func(t *testing.T) {
			conn := &websocket.Conn{}
			tc.table.Layer.AddPlayer(conn, "pauline", false)
			player := tc.table.Layer.GetPlayer(conn)
			assert.NotNil(t, player)
			assert.Equal(t, "pauline", player.Username)
			assert.Equal(t, conn, player.Conn)
		})

		t.Run("Function GetPlayerByUuid", func(t *testing.T) {
			conn := &websocket.Conn{}
			player := tc.table.Layer.AddPlayer(conn, "pauline", false)
			result := tc.table.Layer.GetPlayerByUuid(player.Uuid)
			assert.NotNil(t, result)
			assert.Equal(t, "pauline", result.Username)
			assert.Equal(t, conn, result.Conn)
		})

		t.Run("Function NextPlayer", func(t *testing.T) {
			conn1 := &websocket.Conn{}
			conn2 := &websocket.Conn{}
			conn3 := &websocket.Conn{}

			tc.table.Layer.AddPlayer(conn1, "pauline", false)
			tc.table.Layer.AddPlayer(conn2, "john", false)
			tc.table.Layer.AddPlayer(conn3, "sarah", false)
			tc.table.Layer.StartGame()

			currentPlayerIndex := tc.table.Layer.PrintDetails().CurrentPlayerIndex
			tc.table.Layer.NextPlayer(nil)

			newCurrentPlayerIndex := tc.table.Layer.PrintDetails().CurrentPlayerIndex
			assert.Equal(t, (currentPlayerIndex+1)%3, newCurrentPlayerIndex)
		})

		t.Run("Function GetNumberOfPlayers", func(t *testing.T) {
			newTable := models.CreatePlayersTable()
			conn1 := &websocket.Conn{}
			conn2 := &websocket.Conn{}
			newTable.Layer.AddPlayer(conn1, "pauline", false)
			newTable.Layer.AddPlayer(conn2, "john", false)
			result := newTable.Layer.GetNumberOfPlayers()
			assert.Equal(t, result, 2)
		})
	}
}
