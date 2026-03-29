package logic

func SumValues(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func CalcualtePlayerScores(player *ConnectedPlayer) {
	for _, card := range player.Details.Cards {
		player.Details.Score += card.Value
	}
}
