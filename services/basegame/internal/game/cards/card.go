package cards

import "slices"

const (
	STANDARD_CARD         = "standard"
	STANDARD_SPECIAL_CARD = "standard_special"
)

const (
	OPERATION_ADD           = "+"
	OPERATION_SUBTRACT      = "-"
	OPERATION_MULTIPLY      = "*"
	OPERATION_DIVIDE        = "/"
	OPERATION_FREEZE        = "freeze"
	OPERATION_FLIP3         = "flip3"
	OPERATION_SECOND_CHANCE = "second_chance"
)

// Additional options that may be needed
// for resolving card effects in the future
// useful for special cards that may have
// more complex effects
type ResolveOptions struct {
	// Uuid of the player to whom the
	// card effect is being applied, if applicable
	ToPlayer string
}

type BaseCard struct {
	Uuid         string `json:"uuid"`
	CardType     string `json:"cardType"`
	CardValue    int    `json:"cardValue"`
	CardOperator string `json:"cardOperator"`
	Discarded    bool   `json:"discarded"`
}

func (b *BaseCard) GetUuid() string {
	return b.Uuid
}

func (b *BaseCard) Resolve(currentScore int, options ResolveOptions) int {
	newScore := currentScore

	switch b.CardOperator {
	case OPERATION_ADD:
		newScore += b.CardValue
	case OPERATION_SUBTRACT:
		newScore -= b.CardValue
	case OPERATION_MULTIPLY:
		newScore *= b.CardValue
	case OPERATION_DIVIDE:
		if b.CardValue != 0 {
			newScore /= b.CardValue
		}
	}
	return newScore
}

func (b *BaseCard) SetPlayer(playerUuid string) {
	b.Uuid = playerUuid
}

func (b *BaseCard) GetPlayer() string {
	return b.Uuid
}

func (b *BaseCard) GetValue() int {
	return b.CardValue
}

func (b *BaseCard) IsSpecial() bool {
	values := []string{OPERATION_FREEZE, OPERATION_FLIP3, OPERATION_SECOND_CHANCE}
	return slices.Contains(values, b.CardOperator)
}

// MakeCards is a factory function that returns the appropriate CardFactory
// in order to create cards the corresponding cards.
func MakeCards(cardType string) CardFactory {
	switch cardType {
	case STANDARD_CARD:
		return &StandardCardFactory{}
	case STANDARD_SPECIAL_CARD:
		return &StandardSpecialCardsFactory{}
	default:
		return nil
	}
}
