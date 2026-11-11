package characters

import (
	"fmt"
	"uuid"

	"github.com/Zadigo/citadels/internal/models"
	"github.com/Zadigo/citadels/internal/utils"
)

const (
	AssassinCharacter     string = "assassin"
	WitchCharacter        string = "witch"
	MagistrateCharacter   string = "magistrate"
	ThiefCharacter        string = "thief"
	SpyCharacter          string = "spy"
	MagicianCharacter     string = "magician"
	BlackmailerCharacter  string = "blackmailer"
	WizardCharacter       string = "wizard"
	SeerCharacter         string = "seer"
	KingCharacter         string = "king"
	PatricianCharacter    string = "patrician"
	EmperorCharacter      string = "emperor"
	BishopCharacter       string = "bishop"
	AbbotCharacter        string = "abbot"
	AlchemistCharacter    string = "alchemist"
	TraderCharacter       string = "trader"
	ArchitectCharacter    string = "architect"
	CardinalCharacter     string = "cardinal"
	NavigatorCharacter    string = "navigator"
	MerchantCharacter     string = "merchant"
	ScholarCharacter      string = "scholar"
	WarlordCharacter      string = "warlord"
	QueenCharacter        string = "queen"
	DiplomatCharacter     string = "diplomat"
	ArtistCharacter       string = "artist"
	TaxCollectorCharacter string = "tax_collector"
	MarshalCharacter      string = "marshal"
)

func checkOptions(options models.CharacterOptions) error {
	if options.Target.GetPlayer() == nil {
		return fmt.Errorf("Target has no player")
	}

	if options.CurrentGame == nil {
		return fmt.Errorf("No game")
	}

	return nil
}

type Character struct {
	models.BaseCharacter
}

func (c *Character) Resolve(options models.CharacterOptions) error {
	return nil
}

func (c *Character) Kill() {
	c.Killed = true
}

// NewCharacter creates a new Citadels character:
//
//	character := NewCharacter()
//	character.SetPlayer(player)
//
// A character can then be resolved by calling the Resolve method
//
//	err := character.Resolve(models.CharacterOptions{})
func NewCharacter(name string) (models.CharacterInterface, error) {
	characterId := uuid.New().String()
	baseCharacter := models.BaseCharacter{
		ID:                       characterId,
		Level:                    0,
		Name:                     "",
		Image:                    "",
		Killed:                   false,
		Bewitched:                nil,
		Player:                   nil,
		CanBuildDistricts:        true,
		CanUseCharacterAbilities: true,
	}

	switch name {
	case AssassinCharacter:
		baseCharacter.Level = 1
		baseCharacter.Name = AssassinCharacter
		return &Assassin{Character: Character{baseCharacter}}, nil
	case ThiefCharacter:
		baseCharacter.Level = 2
		baseCharacter.Name = ThiefCharacter
		return &Thief{Character: Character{baseCharacter}}, nil
	case MagicianCharacter:
		baseCharacter.Level = 3
		baseCharacter.Name = MagicianCharacter
		return &Magician{Character: Character{baseCharacter}}, nil
	case KingCharacter:
		baseCharacter.Level = 4
		baseCharacter.Name = KingCharacter
		return &King{Character: Character{baseCharacter}}, nil
	case BishopCharacter:
		baseCharacter.Level = 5
		baseCharacter.Name = BishopCharacter
		return &Bishop{Character: Character{baseCharacter}}, nil
	case MerchantCharacter:
		baseCharacter.Level = 6
		baseCharacter.Name = MerchantCharacter
		return &Merchant{Character: Character{baseCharacter}}, nil
	case WarlordCharacter:
		baseCharacter.Level = 8
		baseCharacter.Name = WarlordCharacter
		return &Warlord{Character: Character{baseCharacter}}, nil
	case ArchitectCharacter:
		baseCharacter.Level = 8
		baseCharacter.Name = ArchitectCharacter
		return &Architect{Character: Character{baseCharacter}}, nil
	case QueenCharacter:
		baseCharacter.Level = 7
		baseCharacter.Name = QueenCharacter
		return &Queen{Character: Character{baseCharacter}}, nil
	case TaxCollectorCharacter:
		baseCharacter.Level = 7
		baseCharacter.Name = TaxCollectorCharacter
		return &TaxCollector{Character: Character{baseCharacter}}, nil
	}
	return nil, fmt.Errorf("unknown character: %s", name)
}

func NewCharacters(names ...string) (characters []models.CharacterInterface, err error) {
	characters = make([]models.CharacterInterface, len(names))

	if utils.HasDuplicates(names) {
		return nil, fmt.Errorf("duplicates found")
	}

	for i, n := range names {
		c, err := NewCharacter(n)
		if err != nil {
			return nil, err
		}
		characters[i] = c
	}
	return
}
