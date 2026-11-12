package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/characters"
	"github.com/Zadigo/citadels/internal/models"
)

func BenchmarkAssasin(b *testing.B) {
	b.ResetTimer()

	x := &characters.Assassin{}
	y := &characters.Bishop{}

	for b.Loop() {
		x.Resolve(models.CharacterOptions{Target: y})
	}
}

func BenchmarkThief(b *testing.B) {
	b.ResetTimer()

	x := &characters.Thief{Player: models.NewPlayer(nil)}
	y := &characters.Bishop{Player: models.NewPlayer(nil)}

	for b.Loop() {
		x.Player.ResetGold()
		y.Player.ResetGold()
		x.Resolve(models.CharacterOptions{Target: y})
	}
}

func BenchmarkMerchant(b *testing.B) {
	b.ResetTimer()

	x := &characters.Merchant{Player: models.NewPlayer(nil)}

	for b.Loop() {
		x.Resolve(models.CharacterOptions{})
	}
}
