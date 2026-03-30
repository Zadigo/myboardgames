package tests

import (
	"log"
	"testing"

	"github.com/Zadigo/flipseven/internal"
	"github.com/Zadigo/flipseven/internal/cards"
)

func TestCheckPlayerCardsNoDuplicates(*testing.T) {
	noDuplicateNumbers := []internal.Card{}
	noDuplicateNumbers = append(noDuplicateNumbers, internal.Card{
		Value:        10,
		Owner:        0,
		Category:     "",
		IsMultiplier: false,
		IsNumber:     false,
		IsBonus:      false,
		IsSpecial:    false,
	})

	player := internal.Player{
		Username:      "test",
		Uuid:          "test-uuid",
		TableUuid:     "test-table-uuid",
		NumberOfCards: 0,
		IsFreezed:     false,
		Cards:         noDuplicateNumbers,
	}

	cards.CheckPlayerCards(&player)

	if player.IsFreezed {
		log.Panic("Player should not be freezed")
	}
}

func TestCheckPlayerCardsWithDuplicates(*testing.T) {
	withDuplicateNubers := []internal.Card{}

	withDuplicateNubers = append(withDuplicateNubers, internal.Card{
		Value:        10,
		Owner:        0,
		Category:     "",
		IsMultiplier: false,
		IsNumber:     false,
		IsBonus:      false,
		IsSpecial:    false,
	})

	withDuplicateNubers = append(withDuplicateNubers, internal.Card{
		Value:        10,
		Owner:        0,
		Category:     "",
		IsMultiplier: false,
		IsNumber:     false,
		IsBonus:      false,
		IsSpecial:    false,
	})

	player := internal.Player{
		Username:      "test",
		Uuid:          "test-uuid",
		TableUuid:     "test-table-uuid",
		NumberOfCards: 0,
		IsFreezed:     false,
		Cards:         withDuplicateNubers,
	}

	cards.CheckPlayerCards(&player)

	if !player.IsFreezed {
		log.Panic("Player should be freezed")
	}
}

func TestAttributeCardToPlayer(*testing.T) {

}

func TestFlipCard(*testing.T) {

}
