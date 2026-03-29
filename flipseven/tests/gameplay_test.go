package tests

import (
	"log"
	"testing"

	"github.com/Zadigo/flipseven/internal/cards"
	"github.com/Zadigo/flipseven/internal/logic"
)

func TestCheckPlayerCardsNoDuplicates(*testing.T) {
	noDuplicateNumbers := []cards.Card{}
	noDuplicateNumbers = append(noDuplicateNumbers, cards.Card{
		Value:        10,
		Owner:        0,
		Category:     "",
		IsMultiplier: false,
		IsNumber:     false,
		IsBonus:      false,
		IsSpecial:    false,
	})

	player := logic.Player{
		Username:      "test",
		Uuid:          "test-uuid",
		TableUuid:     "test-table-uuid",
		NumberOfCards: 0,
		IsFreezed:     false,
		Cards:         noDuplicateNumbers,
	}

	logic.CheckPlayerCards(&player)

	if player.IsFreezed {
		log.Panic("Player should not be freezed")
	}
}

func TestCheckPlayerCardsWithDuplicates(*testing.T) {
	withDuplicateNubers := []cards.Card{}

	withDuplicateNubers = append(withDuplicateNubers, cards.Card{
		Value:        10,
		Owner:        0,
		Category:     "",
		IsMultiplier: false,
		IsNumber:     false,
		IsBonus:      false,
		IsSpecial:    false,
	})

	withDuplicateNubers = append(withDuplicateNubers, cards.Card{
		Value:        10,
		Owner:        0,
		Category:     "",
		IsMultiplier: false,
		IsNumber:     false,
		IsBonus:      false,
		IsSpecial:    false,
	})

	player := logic.Player{
		Username:      "test",
		Uuid:          "test-uuid",
		TableUuid:     "test-table-uuid",
		NumberOfCards: 0,
		IsFreezed:     false,
		Cards:         withDuplicateNubers,
	}

	logic.CheckPlayerCards(&player)

	if !player.IsFreezed {
		log.Panic("Player should be freezed")
	}
}

func TestAttributeCardToPlayer(*testing.T) {

}

func TestFlipCard(*testing.T) {

}
