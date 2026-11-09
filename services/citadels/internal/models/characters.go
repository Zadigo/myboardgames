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
	// The character that a card wants to attack/affect
	Target CharacterInterface `json:"target"`
	// The new target after witching (e.g. if the witch
	// attacks the assassin, it can set a new target to eliminate
	// other than the card that was witched)
	ReplacedTarget CharacterInterface `json:"replaced_target"`
	CurrentGame    GameInterface      `json:"current_game"`
	Magician       MagicianOptions    `json:"magician"`
	// The player for whom this action is taken. In other words,
	// the resolve (gains) of the targetted character will go to this player
	// (the one controlling the attacker)
	ForPlayer PlayerInterface `json:"for_player"`
}

type CharacterInterface interface {
	// Resolve calls the character's action with the given options.
	//
	// 	options := models.CharacterOptions{Target: characters.Architect{}}
	// 	character.Resolve(options) // is called to perform the action.
	Resolve(options CharacterOptions) error
	// Resolve secondary actions of the card e.g. gaining points etc.
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
	BeWitch(witch CharacterInterface)
	Reset()
}

type BaseCharacter struct {
	ID     string `json:"id"`
	Level  int    `json:"level"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Killed bool   `json:"killed"`
	// The card/character that bewitched this one
	Bewitched CharacterInterface `json:"bewitched"`
	// The player controlling this character
	Player PlayerInterface `json:"player"`
	// Whether the player can build districts this turn
	CanBuildDistricts bool `json:"can_build_districts"`
	// Whether the character can use abilities this turn
	CanUseCharacterAbilities bool `json:"can_use_character_abilities"`
}

func (b *BaseCharacter) ResolveSecondary() error {
	return nil
}

func (b *BaseCharacter) BeWitch(witch CharacterInterface) {
	b.Bewitched = witch
	b.CanBuildDistricts = false
	b.CanUseCharacterAbilities = false
}

func (b *BaseCharacter) Reset() {
	b.Killed = false
	b.Bewitched = nil
	b.Player = nil
	b.CanBuildDistricts = true
	b.CanUseCharacterAbilities = true
}
