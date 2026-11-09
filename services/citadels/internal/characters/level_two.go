package characters

import (
	"fmt"

	"github.com/Zadigo/citadels/internal/models"
)

// Call the name of another character whom you wish to rob.
// When a player reveals that character to take their turn, you immediately
// take all of their gold. You cannot rob the rank 1 character (Assassin, Witch, or Magistrate),
// the killed character, or the bewitched character.
type Thief struct {
	Character
}

func (t *Thief) Resolve(options models.CharacterOptions) error {
	if options.Target.GetPlayer() == nil {
		return fmt.Errorf("Target has no player")
	}

	if t.Player == options.Target.GetPlayer() {
		return CannotAttackSelfError
	}

	// if options.CurrentGame == nil {
	// 	return fmt.Errorf("No game")
	// }

	if options.Target.GetLevel() == 1 {
		return CannotStealLevelOneError
	}

	if options.Target.IsKilled() {
		return CannotStealKilledError
	}

	if options.Target.IsBewitched() {
		return CannotStealBewitchedError
	}

	if t.IsBewitched() {
		// Do something
	}

	_targettedPlayer := options.Target.GetPlayer()
	t.Player.IncreaseGold(_targettedPlayer.GoldCount())
	_targettedPlayer.DecreaseGold(_targettedPlayer.GoldCount())

	return nil
}
