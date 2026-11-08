package characters

import "github.com/Zadigo/citadels/internal/models"

type CharacterOptions struct {
	Target      CharacterInterface
	CurrentGame models.GameInterface
}
