package tests

import (
	"testing"

	"github.com/Zadigo/oriflamme/internal/backend"
	"github.com/stretchr/testify/assert"
)

func cardConstructor() (*backend.WebsocketClient, []backend.Card) {
	alice := &backend.WebsocketClient{Username: "Alice", Tokens: 1}
	blueCards := backend.CreateBlueCards(alice)

	return alice, blueCards
}

func TestCard(t *testing.T) {
	card := backend.Card{
		IsRemoved:   false,
		IsSelected:  true,
		IsDiscarded: false,
	}

	t.Run("Card revealed", func(t *testing.T) {
		assert.False(t, card.IsRevealed)

		card.Reveal(&backend.InfluenceQueue{})
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
		error    error
	}

	testCases := []testCase{
		{"assassination", false, nil},
		{"archer", false, nil},
		{"soldier", false, nil},
		{"spy", false, nil},
	}

	for _, tc := range testCases {
		t.Run("Invalid action: "+tc.action, func(t *testing.T) {
			var result bool

			switch tc.action {
			case "assassination":
				result, _ = card.ApplyAssassination(&backend.InfluenceQueue{}, backend.PlayerChoices{AtIndex: 1})
			case "archer":
				result, _ = card.ApplyArcher(&backend.InfluenceQueue{}, backend.PlayerChoices{FirstCard: true})
			case "soldier":
				result, _ = card.ApplySoldier(&backend.InfluenceQueue{}, backend.PlayerChoices{CardBefore: false})
			case "spy":
				result, _ = card.ApplySpy(&backend.InfluenceQueue{}, backend.PlayerChoices{CardBefore: true})
			}

			assert.NoError(t, tc.error)
			assert.Equal(t, tc.expected, result, "Expected %s to return %v", tc.action, tc.expected)
		})
	}
}

func TestArcherCard(t *testing.T) {
	alice := &backend.WebsocketClient{Username: "Alice"}
	blueCards := backend.CreateBlueCards(alice)
	simpleQueue := backend.CreateInfluenceQueue()

	type testCase struct {
		name      string
		errorText string
		firstCard bool
		queue     []*backend.Card
	}

	testCases := []testCase{
		{
			name:      "Archer in first position",
			firstCard: false,
			errorText: "Expected only the last card in the queue to be eliminated when the archer is in the first position",
			queue: []*backend.Card{
				&blueCards[0],
				{Uuid: "1", Name: "Test Card 1", IsRevealed: false},
				{Uuid: "2", Name: "Test Card 2", IsRevealed: false},
			},
		},
		{
			name:      "Archer in middle position",
			firstCard: true,
			errorText: "Expected the first or last card in the queue to be eliminated when the archer is in the middle position",
			queue: []*backend.Card{
				{Uuid: "1", Name: "Test Card 1", IsRevealed: false},
				&blueCards[0],
				{Uuid: "2", Name: "Test Card 2", IsRevealed: false},
			},
		},
		{
			name:      "Archer in last position",
			firstCard: true,
			errorText: "Expected only the first card in the queue to be eliminated when the archer is in the last position",
			queue: []*backend.Card{
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

			currentCard, _ := simpleQueue.GetCurrentCard()
			assert.Equal(t, "Archer", currentCard.Name)
			currentCard.Reveal(simpleQueue)

			result, _ := currentCard.ApplyArcher(simpleQueue, backend.PlayerChoices{FirstCard: tc.firstCard})
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
	alice := &backend.WebsocketClient{Username: "Alice"}
	blueCards := backend.CreateBlueCards(alice)

	simpleQueue := backend.CreateInfluenceQueue()

	type testCase struct {
		name            string
		resolutionIndex int
		cardBefore      bool
		errorText       string
		firstCard       bool
		queue           []*backend.Card
	}

	testCases := []testCase{
		{
			name:            "Soldier in first position",
			resolutionIndex: 0,
			cardBefore:      false,
			firstCard:       false,
			errorText:       "Expected card on the right of the soldier to be eliminated when the soldier is in the first position",
			queue: []*backend.Card{
				&blueCards[1],
				{Uuid: "1", Name: "Test Card 1"},
				{Uuid: "2", Name: "Test Card 2"},
			},
		},
		{
			name:            "Soldier in middle position 1",
			resolutionIndex: 1,
			cardBefore:      true,
			firstCard:       true,
			errorText:       "Expected the card on the left of the soldier to be eliminated when the soldier is in the middle position",
			queue: []*backend.Card{
				{Uuid: "1", Name: "Test Card 1"},
				&blueCards[1],
				{Uuid: "2", Name: "Test Card 2"},
			},
		},
		{
			name:            "Soldier in middle position 2",
			cardBefore:      false,
			resolutionIndex: 1,
			firstCard:       true,
			errorText:       "Expected the card on the right of the soldier to be eliminated when the soldier is in the middle position",
			queue: []*backend.Card{
				{Uuid: "1", Name: "Test Card 1"},
				&blueCards[1],
				{Uuid: "2", Name: "Test Card 2"},
			},
		},
		{
			name:            "Soldier in last position",
			cardBefore:      true,
			resolutionIndex: 2,
			firstCard:       true,
			errorText:       "Expected card on the left of the soldier to be eliminated when the soldier is in the last position",
			queue: []*backend.Card{
				{Uuid: "1", Name: "Test Card 1"},
				{Uuid: "2", Name: "Test Card 2"},
				&blueCards[1],
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			simpleQueue.Queue = tc.queue
			simpleQueue.ResolutionIndex = tc.resolutionIndex

			currentCard, _ := simpleQueue.GetCurrentCard()
			assert.Equal(t, "Soldier", currentCard.Name)
			currentCard.Reveal(simpleQueue)

			result, err := currentCard.ApplySoldier(simpleQueue, backend.PlayerChoices{CardBefore: tc.cardBefore})
			assert.NoError(t, err, err)
			assert.True(t, result, tc.errorText)

			if tc.cardBefore {
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

func TestSpyCard(t *testing.T) {
	alice := &backend.WebsocketClient{Username: "Alice", Tokens: 1}
	martin := &backend.WebsocketClient{Username: "Martin", Tokens: 1}
	pauline := &backend.WebsocketClient{Username: "Pauline", Tokens: 1}

	blueCards := backend.CreateBlueCards(alice)

	simpleQueue := backend.CreateInfluenceQueue()

	type testCase struct {
		name                   string
		resolutionIndex        int
		expectedResult         bool
		expectedOwnerTokens    int
		expectedAdjacentTokens int
		errorText              string
		cardBefore             bool
		queue                  []*backend.Card
	}

	testCases := []testCase{
		{
			name:                   "Spy in first position",
			resolutionIndex:        0,
			cardBefore:             false,
			expectedResult:         true,
			expectedOwnerTokens:    2,
			expectedAdjacentTokens: 0,
			errorText:              "Expected the owner of the next adjacent card to lose 1 token when the spy is in the first position",
			queue: []*backend.Card{
				&blueCards[2],
				{Uuid: "1", Name: "Test Card 1", Owner: martin},
				{Uuid: "2", Name: "Test Card 2", Owner: pauline},
			},
		},
		{
			name:                   "Spy in middle position",
			resolutionIndex:        1,
			cardBefore:             true,
			expectedResult:         true,
			expectedOwnerTokens:    2,
			expectedAdjacentTokens: 0,
			errorText:              "Expected the owner of the next adjacent card to lose 1 token when the spy is in the middle position",
			queue: []*backend.Card{
				{Uuid: "1", Name: "Test Card 1", Owner: martin},
				&blueCards[2],
				{Uuid: "2", Name: "Test Card 2", Owner: pauline},
			},
		},
		{
			name:                   "Spy in last position",
			resolutionIndex:        2,
			cardBefore:             true,
			expectedResult:         true,
			expectedOwnerTokens:    2,
			expectedAdjacentTokens: 0,
			errorText:              "Expected the owner of the next adjacent card to lose 1 token when the spy is in the last position",
			queue: []*backend.Card{
				{Uuid: "1", Name: "Test Card 1", Owner: martin},
				{Uuid: "2", Name: "Test Card 2", Owner: pauline},
				&blueCards[2],
			},
		},
		{
			name:                   "Spy in last position and no card after",
			resolutionIndex:        2,
			cardBefore:             true, // This should have no effect since there is no card after the spy to steal from
			expectedResult:         true,
			expectedOwnerTokens:    2,
			expectedAdjacentTokens: 0,
			errorText:              "Cannot steal from the next adjacent card when the spy is in the last position and there is no card after it",
			queue: []*backend.Card{
				{Uuid: "1", Name: "Test Card 1", Owner: martin},
				{Uuid: "2", Name: "Test Card 2", Owner: pauline},
				&blueCards[2],
			},
		},
		{
			name:                   "Spy on my own card",
			resolutionIndex:        2,
			cardBefore:             true,
			expectedResult:         false,
			expectedOwnerTokens:    1,
			expectedAdjacentTokens: 1,
			errorText:              "Expected the same owner of the next adjacent card to not lose 1 token",
			queue: []*backend.Card{
				{Uuid: "1", Name: "Test Card 1", Owner: martin},
				{Uuid: "2", Name: "Test Card 2", Owner: alice},
				&blueCards[2],
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			simpleQueue.Queue = tc.queue
			simpleQueue.ResolutionIndex = tc.resolutionIndex

			currentCard, _ := simpleQueue.GetCurrentCard()
			assert.Equal(t, "Spy", currentCard.Name)
			currentCard.Reveal(simpleQueue)

			result, err := currentCard.ApplySpy(simpleQueue, backend.PlayerChoices{CardBefore: tc.cardBefore})
			if tc.expectedResult {
				assert.True(t, result)
				assert.NoError(t, err, err)
			} else {
				assert.False(t, result)
				assert.Error(t, err)
			}

			resultingAdjacentTokens := 0

			if tc.cardBefore {
				resultingAdjacentTokens = simpleQueue.Queue[tc.resolutionIndex-1].Owner.Tokens
			} else {
				resultingAdjacentTokens = simpleQueue.Queue[tc.resolutionIndex+1].Owner.Tokens
			}

			assert.Equal(t, resultingAdjacentTokens, tc.expectedAdjacentTokens, tc.errorText)
			assert.Equal(t, currentCard.Owner.Tokens, tc.expectedOwnerTokens, "Expected player to gain tokens stealing from the adjacent card")

			t.Cleanup(func() {
				currentCard.Owner.Tokens = 1

				for _, card := range simpleQueue.Queue {
					card.IsRemoved = false
				}
			})
		})
	}
}

func TestConspiracyCard(t *testing.T) {
	_, blueCards := cardConstructor()

	simpleQueue := backend.InfluenceQueue{
		ResolutionIndex: -1,
		Queue: []*backend.Card{
			&blueCards[8],
		},
	}

	t.Run("Conspiracy card revealed", func(t *testing.T) {
		simpleQueue.ResolutionIndex = 0
		currentCard, _ := simpleQueue.GetCurrentCard()

		currentCard.Tokens = 2
		currentCard.Reveal(&simpleQueue)
		state, err := currentCard.ApplyConspiracy(&simpleQueue)

		assert.True(t, state, "Expected effect to be applied successfully")
		assert.NoError(t, err, "Expected no error when applying conspiracy effect")
		assert.Equal(t, 5, currentCard.Owner.Tokens, "Expected player to gain 4 tokens for a total of 5")
	})
}

func TestAmbush(t *testing.T) {
	blueCards := backend.CreateBlueCards(&backend.WebsocketClient{Username: "Alice", Tokens: 1})
	redCards := backend.CreateRedCards(&backend.WebsocketClient{Username: "Pauline", Tokens: 1})

	type testCase struct {
		name                   string
		resolutionIndex        int
		revealed               bool
		expectedOwnerTokens    int
		expectedAdjacentTokens int
		errorText              string
		queue                  []*backend.Card
	}

	testCases := []testCase{
		{
			name:                   "Soldier eliminates Ambush card",
			resolutionIndex:        0,
			revealed:               false,
			expectedOwnerTokens:    2,
			expectedAdjacentTokens: 5,
			errorText:              "Expected owner of the ambush card to gain 4 tokens when the soldier eliminates it",
			queue: []*backend.Card{
				&redCards[1],  // Solider card
				&blueCards[9], // Ambush card
			},
		},
		{
			name:                "Player reveals Ambush card and it is not eliminated",
			resolutionIndex:     0,
			revealed:            true,
			expectedOwnerTokens: 2,
			errorText:           "Expected owner of the ambush card to gain 1 token",
			queue: []*backend.Card{
				&redCards[9],
			},
		},
	}

	simpleQueue := backend.CreateInfluenceQueue()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			simpleQueue.Queue = tc.queue
			simpleQueue.ResolutionIndex = tc.resolutionIndex

			currentCard, err := simpleQueue.GetCurrentCard()
			currentCard.Reveal(simpleQueue)

			assert.NoError(t, err, err)

			if tc.revealed {
				assert.Equal(t, "Ambush", currentCard.Name)
				assert.Equal(t, tc.expectedOwnerTokens, currentCard.Owner.Tokens, tc.errorText)
			} else {
				assert.Equal(t, "Soldier", currentCard.Name)
				// Simulate the soldier card effect by attacking the
				// card on the right (or the Ambush card)
				_, err := currentCard.ApplySoldier(simpleQueue, backend.PlayerChoices{CardBefore: false})
				assert.NoError(t, err, err)
				assert.Equal(t, tc.expectedOwnerTokens, currentCard.Owner.Tokens)
				assert.Equal(t, tc.expectedAdjacentTokens, simpleQueue.Queue[1].Owner.Tokens, tc.errorText)
			}

			t.Cleanup(func() {
				currentCard.Owner.Tokens = 1

				for _, card := range simpleQueue.Queue {
					card.IsRemoved = false
				}
			})
		})
	}
}

func TestShapeShifterCard(t *testing.T) {
	alice := &backend.WebsocketClient{Username: "Alice", Tokens: 1}
	pauline := &backend.WebsocketClient{Username: "Pauline", Tokens: 1}
	martin := &backend.WebsocketClient{Username: "Martin", Tokens: 1}

	blueCards := backend.CreateBlueCards(alice)
	redCards := backend.CreateRedCards(pauline)
	yellowCards := backend.CreateYellowCards(martin)

	simpleQueue := backend.CreateInfluenceQueue()

	type testCase struct {
		name                string
		resolutionIndex     int
		expectedResult      bool
		expectedOwnerTokens int
		errorText           string
		cardBefore          bool
		adjacentRevealed    bool
		queue               []*backend.Card
	}

	testCases := []testCase{
		{
			name:                "Shapeshifter copies revealed card",
			resolutionIndex:     0,
			cardBefore:          false,
			expectedResult:      false,
			expectedOwnerTokens: 2,
			adjacentRevealed:    true,
			errorText:           "Expected the owner of the next adjacent card to lose 1 token when the shape shifter is in the first position",
			queue: []*backend.Card{
				&blueCards[4],
				&redCards[1],    // Solider card
				&yellowCards[3], // Heir
			},
		},
		// {
		// 	name:                   "Shapeshifter copies none revealed card",
		// 	resolutionIndex:        0,
		// 	cardBefore:             false,
		// 	expectedResult:         true,
		// 	expectedOwnerTokens:    2,
		// 	expectedAdjacentTokens: 0,
		// 	adjacentRevealed:       false,
		// 	errorText:              "Expected the owner of the next adjacent card to lose 1 token when the shape shifter is in the first position",
		// 	queue: []*backend.Card{
		// 		&blueCards[4],
		// 		&redCards[1],    // Solider card
		// 		&yellowCards[3], // Heir
		// 	},
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			simpleQueue.Queue = tc.queue
			simpleQueue.ResolutionIndex = tc.resolutionIndex

			currentCard, _ := simpleQueue.GetCurrentCard()
			assert.Equal(t, "Shapeshifter", currentCard.Name)
			currentCard.Reveal(simpleQueue)

			if tc.adjacentRevealed {
				simpleQueue.Queue[1].Reveal(simpleQueue)
			}

			result, err := currentCard.ApplyShapeshifter(simpleQueue, backend.PlayerChoices{
				CardBefore:             tc.cardBefore,
				ShapeShifterCardBefore: false,
			})
			assert.NoError(t, err, err)

			if tc.adjacentRevealed {
				assert.True(t, result, tc.errorText)
			} else {
				assert.False(t, result, tc.errorText)
			}

			assert.Equal(t, tc.expectedOwnerTokens, currentCard.Owner.Tokens)

			t.Cleanup(func() {
				currentCard.Owner.Tokens = 1

				for _, card := range simpleQueue.Queue {
					card.IsRemoved = false
				}

				simpleQueue.Queue[1].IsRevealed = false
			})
		})
	}
}

func TestHeir(t *testing.T) {
	alice := &backend.WebsocketClient{Username: "Alice", Tokens: 1}
	pauline := &backend.WebsocketClient{Username: "Pauline", Tokens: 1}

	blueCards := backend.CreateBlueCards(alice)
	redCards := backend.CreateRedCards(pauline)

	simpleQueue := backend.CreateInfluenceQueue()

	type testCase struct {
		name            string
		resolutionIndex int
		expectedTokens  int
		queue           []*backend.Card
	}

	testCases := []testCase{
		{
			name:            "Heir is in the only one in the queue",
			resolutionIndex: 0,
			expectedTokens:  3,
			queue: []*backend.Card{
				&blueCards[3],
			},
		},
		{
			name:            "Heir is not the only one in the queue",
			resolutionIndex: 0,
			expectedTokens:  1,
			queue: []*backend.Card{
				&blueCards[3],
				&redCards[3],
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			simpleQueue.Queue = tc.queue
			simpleQueue.ResolutionIndex = tc.resolutionIndex

			currentCard, _ := simpleQueue.GetCurrentCard()
			assert.Equal(t, "Heir", currentCard.Name)
			currentCard.Reveal(simpleQueue)

			_, err := currentCard.ApplyHeir(simpleQueue)
			assert.NoError(t, err, err)
			assert.Equal(t, tc.expectedTokens, currentCard.Owner.Tokens)

			t.Cleanup(func() {
				for _, card := range simpleQueue.Queue {
					card.Owner.Tokens = 1
				}
			})
		})
	}
}
