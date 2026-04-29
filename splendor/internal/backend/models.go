package backend

import "github.com/google/uuid"

// CardResources represents the resources
// required to buy a card in the Splendor game.
type CardResources struct {
	Emerald  int `json:"emerald"`
	Diamond  int `json:"diamond"`
	Sapphire int `json:"sapphire"`
	Onyx     int `json:"onyx"`
	Ruby     int `json:"ruby"`
}

// MarvelCardResources represents the specific
// resources required for the Splendor Marvel cards.
type MarvelCardResources struct {
	Mind    int `json:"mind"`
	Space   int `json:"space"`
	Soul    int `json:"soul"`
	Power   int `json:"power"`
	Reality int `json:"reality"`
	Time    int `json:"time"`
	Shield  int `json:"shield"`
}

// CardDetails contains the common details
// for both Normal and Marvel cards.
type CardDetails struct {
	Uuid   string  `json:"uuid"`
	Name   string  `json:"name"`
	Level  int     `json:"level"`
	Points int     `json:"points"`
	Owner  *Player `json:"owner"`
}

// NormalCard represents a standard
// card in the Splendor game.
type NormalCard struct {
	CardResources
	CardDetails
}

// MarvelCard represents a card from the
// Marvel expansion of the Splendor game.
type MarvelCard struct {
	MarvelCardResources
	CardDetails
}

// CardInterface defines the methods that
// both NormalCard and MarvelCard must implement.
type CardInterface interface {
	// The player buys the card,
	// paying the required resources.
	Buy()
	// The player reserves the card,
	// setting it aside for future purchase.
	Reserve()
}

func CreateCard[T MarvelCard | NormalCard](cardType T, level int, characters []map[string]any) []CardInterface {
	cards := []CardInterface{}

	for _, character := range characters {
		switch any(cardType).(type) {
		case NormalCard:
			card := NormalCard{
				CardResources: CardResources{
					Emerald:  character["emerald"].(int),
					Diamond:  character["diamond"].(int),
					Sapphire: character["sapphire"].(int),
					Onyx:     character["onyx"].(int),
					Ruby:     character["ruby"].(int),
				},
				CardDetails: CardDetails{
					Uuid:   uuid.New().String(),
					Name:   character["name"].(string),
					Level:  level,
					Points: character["points"].(int),
				},
			}
			cards = append(cards, &card)
		case MarvelCard:
			card := MarvelCard{
				MarvelCardResources: MarvelCardResources{
					Mind:    character["mind"].(int),
					Space:   character["space"].(int),
					Soul:    character["soul"].(int),
					Power:   character["power"].(int),
					Reality: character["reality"].(int),
					Time:    character["time"].(int),
					Shield:  character["shield"].(int),
				},
				CardDetails: CardDetails{
					Uuid:   uuid.New().String(),
					Name:   character["name"].(string),
					Level:  level,
					Points: character["points"].(int),
				},
			}
			cards = append(cards, &card)
		}
	}

	return cards
}
