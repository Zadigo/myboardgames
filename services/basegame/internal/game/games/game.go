package games

const (
	STANDARD  = "standard"
	EXTENSION = "extension"
)

// CreateGame is a factory function that creates a new game instance based on the specified game type.
// It returns an instance of GameInterface, which can be either a StandardGame or an ExtensionGame,
// depending on the provided gameType. If the gameType is not recognized, it returns nil.
func CreateGame(gameType string) GameInterface {
	switch gameType {
	case STANDARD:
		return &StandardGame{}
	case EXTENSION:
		return &ExtensionGame{}
	default:
		return nil
	}
}
