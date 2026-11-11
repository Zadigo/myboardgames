package models

import "fmt"

const (
	SignedWarrant     string = "SignedWarrant"
	UnsignedWarrant   string = "UnsignedWarrant"
	WarrantMarkerBack string = "WarrantMarkerBack"
)

// MagicianOptions are specific options
// for the Magician carachter
type MagicianOptions struct {
	Choice          string `json:"choice"`
	DiscardQuantity int    `json:"discard_quantity"`
}

type MagistrateOptions struct {
	WarrantMarkerBack CharacterInterface `json:"warrant_marker_back"`
	SignedWarrant     CharacterInterface `json:"signed_warrant"`
	UnsignedWarrant   CharacterInterface `json:"unsigned_warrant"`
}

type NavigatorOptions struct {
	GainFourGolds bool
	GainFourCards bool
}

// CharacterOptions holds info for resolving a character's action.
// The options can be used to target a character for an action
//
//	c := &characters.Character{}
//	other := &characters.Character{}
//	c.Resolve(models.CharacterOptions{Target: other})
type CharacterOptions struct {
	// The character that a card wants to attack/affect
	Target CharacterInterface `json:"target"`
	// The new target after witching (e.g. if the witch
	// attacks the assassin, it can set a new target to eliminate
	// other than the card that was witched)
	ReplacedTarget CharacterInterface `json:"replaced_target"`
	CurrentGame    GameInterface      `json:"current_game"`
	Magician       MagicianOptions    `json:"magician"`
	Magistrate     MagistrateOptions  `json:"magistrate"`
	Navigator      NavigatorOptions   `json:"navigator"`
	// The player for whom this action is taken. In other words,
	// the resolve (gains) of the targetted character will go to this player
	// (the one controlling the attacker)
	ForPlayer PlayerInterface `json:"for_player"`
}

type CharacterInterface interface {
	// Resolve calls the character's action with the given options.
	//
	// 	options := models.CharacterOptions{Target: &characters.Architect{}}
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
	BeWitch(witch CharacterInterface) error
	PlaceWarrantMarkerBack(magistrate CharacterInterface) error
	PlaceSignedWarrant(magistrate CharacterInterface) error
	PlaceUnsignedWarrant(magistrate CharacterInterface) error
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
	UnsignedWarrant          bool `json:"unsigned_warrant"`
	SignedWarrant            bool `json:"signed_warrant"`
	WarrantMarkerBack        bool `json:"warrant_marker_back"`
}

func (b *BaseCharacter) GetId() string {
	return b.ID
}

func (b *BaseCharacter) GetName() string {
	return b.Name
}

func (b *BaseCharacter) GetLevel() int {
	return b.Level
}

func (b *BaseCharacter) ResolveSecondary() error {
	return nil
}

// BeWitch is used by the other characters when the witch resolves itself.
//
//	w := characters.Character{Name: "witch"}
//	other := characters.Character{}
//	other.BeWitch(w)
//
// This is equivalent to calling `Resolve` on the witch itself which then calls
// “BeWitch“ on the other character
//
//	w.Resolve(models.CharacterOptions{Target: other})
func (b *BaseCharacter) BeWitch(witch CharacterInterface) error {
	if b.Name == "witch" {
		return fmt.Errorf("Cannot witch own self")
	}

	b.Bewitched = witch
	b.CanBuildDistricts = false
	b.CanUseCharacterAbilities = false

	return nil
}

func (b *BaseCharacter) checkWarrantState() error {
	if b.WarrantMarkerBack || b.SignedWarrant || b.UnsignedWarrant {
		return fmt.Errorf("warrant already placed")
	}
	return nil
}

func (b *BaseCharacter) PlaceSignedWarrant(magistrate CharacterInterface) error {
	if err := b.checkWarrantState(); err != nil {
		return err
	} else {
		b.SignedWarrant = true
		return nil
	}
}

func (b *BaseCharacter) PlaceUnsignedWarrant(magistrate CharacterInterface) error {
	if err := b.checkWarrantState(); err != nil {
		return err
	} else {
		b.UnsignedWarrant = true
		return nil
	}
}

func (b *BaseCharacter) PlaceWarrantMarkerBack(magistrate CharacterInterface) error {
	if err := b.checkWarrantState(); err != nil {
		return err
	} else {
		b.WarrantMarkerBack = true
		return nil
	}
}

func (b *BaseCharacter) SetPlayer(player PlayerInterface) {
	b.Player = player
}

func (b *BaseCharacter) GetPlayer() PlayerInterface {
	return b.Player
}

func (b *BaseCharacter) IsKilled() bool {
	return b.Killed
}

func (b *BaseCharacter) IsBewitched() bool {
	return b.Bewitched != nil
}

func (b *BaseCharacter) Reset() {
	b.Killed = false
	b.Bewitched = nil
	b.Player = nil
	b.CanBuildDistricts = true
	b.CanUseCharacterAbilities = true
	b.UnsignedWarrant = false
	b.SignedWarrant = false
	b.WarrantMarkerBack = false
}

// BaseCharacterIndicator is used by the computer on
// the turn phase to keep track of the carachter
// being resolved
type BaseCharacterIndicator struct {
	Name  string
	Level int
}
