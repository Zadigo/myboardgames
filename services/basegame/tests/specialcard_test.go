package tests

import (
	"testing"

	"github.com/Zadigo/basegame/internal/game/cards"
	"github.com/Zadigo/basegame/internal/game/games"
	"github.com/Zadigo/basegame/internal/models"
	"github.com/Zadigo/basegame/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestDrawCardFromPileFlip3(t *testing.T) {
	pile := utils.CreatePileFlipOnTop(t)

	app := games.NewGame(t.Context(), models.GameAppOptions{})
	app.CardPile.AddCards(pile...)

	// Draw the flip card
	result := app.CardPile.DrawCard()
	assert.NotNil(t, result)
	assert.True(t, app.MustResolve)

	score := 0
	score = result.Resolve(score, cards.CardResolveOptions{})

	assert.Equal(t, 0, score)
}
