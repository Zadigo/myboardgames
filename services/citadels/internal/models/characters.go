package models

// MagicianOptions are specific options
// for the Magician carachter
type MagicianOptions struct {
	Choice          string `json:"choice"`
	DiscardQuantity int    `json:"discard_quantity"`
}

// CharacterOptions holds info for resolving a character's action.
//   - Target: the character being targeted
//   - CurrentGame: the game context
type CharacterOptions struct {
	Target      CharacterInterface `json:"target"`
	CurrentGame GameInterface      `json:"current_game"`
	Magician    MagicianOptions    `json:"magician"`
}

type CharacterInterface interface {
	// Resolve calls the character's action with the given options.
	//
	// 	options := models.CharacterOptions{Target: characters.Architect{}}
	// 	character.Resolve(options) // is called to perform the action.
	Resolve(options CharacterOptions) error
	ResolveSecondary() error
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
	ID        string          `json:"id"`
	Level     int             `json:"level"`
	Name      string          `json:"name"`
	Image     string          `json:"image"`
	Killed    bool            `json:"killed"`
	Bewitched bool            `json:"bewitched"`
	Player    PlayerInterface `json:"player"`
}

func (b *BaseCharacter) ResolveSecondary() error {
	return nil
}
