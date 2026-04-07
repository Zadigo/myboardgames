package tests

import (
	"testing"

	"github.com/Zadigo/oriflamme/internal/models"
	"github.com/stretchr/testify/assert"
)

func cardConstructor() (*models.Player, []models.Card) {
	alice := &models.Player{Username: "Alice", Tokens: 1}
	blueCards := models.CreateBlueCards(alice)

	return alice, blueCards
}

func TestCard(t *testing.T) {
	card := models.Card{
		IsRemoved:   false,
		IsSelected:  true,
		IsDiscarded: false,
	}

	t.Run("Card revealed", func(t *testing.T) {
		assert.False(t, card.IsRevealed)

		card.Reveal()
		assert.True(t, card.IsRevealed)
	})

	t.Run("Card discarded", func(t *testing.T) {
		assert.False(t, card.IsDiscarded)

		card.Discard()
		assert.True(t, card.IsDiscarded)
	})

	t.Run("Add and remove tokens", func(t *testing.T) {
		assert.Equal(t, card.Tokens, 0)

		card.AddToken(3)
		assert.Equal(t, card.Tokens, 3)

		card.RemoveToken(2)
		assert.Equal(t, card.Tokens, 1)
	})

	type testCase struct {
		action   string
		expected bool
	}

	testCases := []testCase{
		{"assassination", false},
		{"archer", false},
		{"soldier", false},
		{"spy", false},
	}

	for _, tc := range testCases {
		t.Run("Invalid action: "+tc.action, func(t *testing.T) {
			var result bool

			switch tc.action {
			case "assassination":
				result = card.ApplyAssassination(&models.InfluenceQueue{}, 1)
			case "archer":
				result = card.ApplyArcher(&models.InfluenceQueue{}, true)
			case "soldier":
				result = card.ApplySoldier(&models.InfluenceQueue{}, "after")
			case "spy":
				result = card.ApplySpy(&models.InfluenceQueue{}, true)
			}

			assert.Equal(t, tc.expected, result, "Expected %s to return %v", tc.action, tc.expected)
		})
	}
}

func TestArcherCard(t *testing.T) {
	alice := &models.Player{Username: "Alice"}
	blueCards := models.CreateBlueCards(alice)
	simpleQueue := models.CreateInfluenceQueue()

	type testCase struct {
		name      string
		errorText string
		firstCard bool
		queue     []*models.Card
	}

	testCases := []testCase{
		{
			name:      "Archer in first position",
			firstCard: false,
			errorText: "Expected only the last card in the queue to be eliminated when the archer is in the first position",
			queue: []*models.Card{
				&blueCards[0],
				{Uuid: "1", Name: "Test Card 1", IsRevealed: false},
				{Uuid: "2", Name: "Test Card 2", IsRevealed: false},
			},
		},
		{
			name:      "Archer in middle position",
			firstCard: true,
			errorText: "Expected the first or last card in the queue to be eliminated when the archer is in the middle position",
			queue: []*models.Card{
				{Uuid: "1", Name: "Test Card 1", IsRevealed: false},
				&blueCards[0],
				{Uuid: "2", Name: "Test Card 2", IsRevealed: false},
			},
		},
		{
			name:      "Archer in last position",
			firstCard: true,
			errorText: "Expected only the first card in the queue to be eliminated when the archer is in the last position",
			queue: []*models.Card{
				{Uuid: "1", Name: "Test Card 1", IsRevealed: false},
				{Uuid: "2", Name: "Test Card 2", IsRevealed: false},
				&blueCards[0],
			},
		},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			simpleQueue.Queue = tc.queue
			simpleQueue.ResolutionIndex = i

			currentCard := simpleQueue.GetCurrentCard()
			assert.Equal(t, "Archer", currentCard.Name)
			currentCard.Reveal()

			result := currentCard.ApplyArcher(simpleQueue, tc.firstCard)
			assert.True(t, result, tc.errorText)

			if tc.firstCard {
				assert.True(t, simpleQueue.Queue[0].IsRemoved, tc.errorText)
			} else {
				assert.True(t, simpleQueue.Queue[len(simpleQueue.Queue)-1].IsRemoved, tc.errorText)
			}

			assert.Equal(t, currentCard.Owner.Tokens, 1, "Expected player to gain tokens for eliminated cards")

			t.Cleanup(func() {
				currentCard.Owner.Tokens = 0

				for _, card := range simpleQueue.Queue {
					card.IsRemoved = false
				}
			})
		})
	}
}

func TestSoldierCard(t *testing.T) {
	alice := &models.Player{Username: "Alice"}
	blueCards := models.CreateBlueCards(alice)
	simpleQueue := models.CreateInfluenceQueue()

	type testCase struct {
		name            string
		resolutionIndex int
		direction       string
		errorText       string
		firstCard       bool
		queue           []*models.Card
	}

	testCases := []testCase{
		{
			name:            "Soldier in first position",
			resolutionIndex: 0,
			direction:       "",
			firstCard:       false,
			errorText:       "Expected card on the right of the soldier to be eliminated when the soldier is in the first position",
			queue: []*models.Card{
				&blueCards[1],
				{Uuid: "1", Name: "Test Card 1", IsRevealed: false},
				{Uuid: "2", Name: "Test Card 2", IsRevealed: false},
			},
		},
		{
			name:            "Soldier in middle position 1",
			resolutionIndex: 1,
			direction:       "before",
			firstCard:       true,
			errorText:       "Expected the card on the left of the soldier to be eliminated when the soldier is in the middle position",
			queue: []*models.Card{
				{Uuid: "1", Name: "Test Card 1", IsRevealed: false},
				&blueCards[1],
				{Uuid: "2", Name: "Test Card 2", IsRevealed: false},
			},
		},
		{
			name:            "Soldier in middle position 2",
			direction:       "after",
			resolutionIndex: 1,
			firstCard:       true,
			errorText:       "Expected the card on the right of the soldier to be eliminated when the soldier is in the middle position",
			queue: []*models.Card{
				{Uuid: "1", Name: "Test Card 1", IsRevealed: false},
				&blueCards[1],
				{Uuid: "2", Name: "Test Card 2", IsRevealed: false},
			},
		},
		{
			name:            "Soldier in last position",
			direction:       "",
			resolutionIndex: 2,
			firstCard:       true,
			errorText:       "Expected card on the left of the soldier to be eliminated when the soldier is in the last position",
			queue: []*models.Card{
				{Uuid: "1", Name: "Test Card 1", IsRevealed: false},
				{Uuid: "2", Name: "Test Card 2", IsRevealed: false},
				&blueCards[1],
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			simpleQueue.Queue = tc.queue
			simpleQueue.ResolutionIndex = tc.resolutionIndex

			currentCard := simpleQueue.GetCurrentCard()
			assert.Equal(t, "Soldier", currentCard.Name)
			currentCard.Reveal()

			result := currentCard.ApplySoldier(simpleQueue, tc.direction)
			assert.True(t, result, tc.errorText)

			if tc.direction != "" {
				switch tc.direction {
				case "before":
					assert.True(t, simpleQueue.Queue[0].IsRemoved, tc.errorText)
				case "after":
					assert.True(t, simpleQueue.Queue[len(simpleQueue.Queue)-1].IsRemoved, tc.errorText)
				}
			}

			assert.Equal(t, currentCard.Owner.Tokens, 1, "Expected player to gain tokens for eliminated cards")

			t.Cleanup(func() {
				currentCard.Owner.Tokens = 0

				for _, card := range simpleQueue.Queue {
					card.IsRemoved = false
				}
			})
		})
	}
}

func TestSpyCard(t *testing.T) {
	_, blueCards := cardConstructor()

	simpleQueue := models.InfluenceQueue{
		ResolutionIndex: -1,
		Queue: []*models.Card{
			{Uuid: "1", Name: "Test Card", Tokens: 1, Owner: &models.Player{Username: "Pierre"}},
			&blueCards[2], // Spy
			{Uuid: "1", Name: "Test Card", Tokens: 1, Owner: &models.Player{Username: "Pauline"}},
		},
	}

	simpleQueue.ResolutionIndex = 1
	currentCard := simpleQueue.GetCurrentCard()

	currentCard.Reveal()

	result := currentCard.ApplySpy(&simpleQueue, true)

	if !result {
		t.Error("Expected effect to be applied successfully")
	}

	prevCard := simpleQueue.Queue[0]

	if prevCard.Owner.Tokens != 0 {
		t.Errorf("Expected previous card to lose 1 token, got %d", prevCard.Tokens)
	}
}

func TestConspiracyCard(t *testing.T) {
	_, blueCards := cardConstructor()

	simpleQueue := models.InfluenceQueue{
		ResolutionIndex: -1,
		Queue: []*models.Card{
			&blueCards[8],
		},
	}

	simpleQueue.ResolutionIndex = 0
	currentCard := simpleQueue.GetCurrentCard()
	currentCard.Tokens = 2

	currentCard.Reveal()
	result := currentCard.ApplyConspiracy(&simpleQueue)

	if !result {
		t.Error("Expected effect to be applied successfully")
	}

	if currentCard.Owner.Tokens != 5 {
		t.Errorf("Expected player to gain 4 tokens for a total of 5, got %d", currentCard.Owner.Tokens)
	}
}
