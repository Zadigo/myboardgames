package tests

// func TestGameRedis(t *testing.T) {
// 	type testCase struct {
// 		text string
// 	}

// 	testCases := []testCase{
// 		{
// 			text: "Save",
// 		},
// 		{
// 			text: "Get",
// 		},
// 	}

// 	cards := CreateTestCards()
// 	redisHandler := redis.NewGameRedis(t.Context(), nil)

// 	for _, tc := range testCases {
// 		t.Run(tc.text, func(t *testing.T) {
// 			if tc.text == "Save" {
// 				redisHandler.SaveCards(cards)
// 				assert.NotNil(t, cards)
// 				assert.Len(t, cards, 1)
// 			}
// 		})
// 	}
// }
