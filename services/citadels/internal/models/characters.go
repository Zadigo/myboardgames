package models

// MagicianOptions are specific options
// for the Magician carachter
type MagicianOptions struct {
	Choice          string
	DiscardQuantity int
}

// CharacterOptions holds info for resolving a character's action.
//   - Target: the character being targeted
//   - CurrentGame: the game context
type CharacterOptions struct {
	Target      CharacterInterface
	CurrentGame GameInterface
	Magician    MagicianOptions
}

type CharacterInterface interface {
	// Resolve calls the character's action with the given options.
	//
	// 	options := models.CharacterOptions{Target: characters.Architect{}}
	// 	character.Resolve(options) // is called to perform the action.
	Resolve(options CharacterOptions) error
	GetId() string
	GetName() string
	IsKilled() bool
	IsBewitched() bool
	GetLevel() int
	Kill()
	// SetPlayer assigns the player controlling this character.
	SetPlayer(player PlayerInterface)
	// GetPlayer returns the player controlling this character.
	GetPlayer() PlayerInterface
}

type BaseCharacter struct {
	ID        string
	Level     int
	Name      string
	Image     string
	Killed    bool
	Bewitched bool
	Player    PlayerInterface
}
