package characters

import (
	"github.com/Zadigo/citadels/internal/models"
)

type Assassin struct {
	Character
}

func (a *Assassin) Resolve(options models.CharacterOptions) error {
	if options.Target == nil && options.ReplacedTarget == nil {
		return MissingTargetError
	}

	if options.Target != nil && options.ReplacedTarget != nil {
		return AmbiguousTargetError
	}

	_target := options.Target
	if options.ReplacedTarget != nil {
		_target = options.ReplacedTarget
	}

	_target.Kill()
	return nil
}

type Witch struct {
	Character
	// The target carachter that witch affects
	target models.CharacterInterface
	// The new target after witching (e.g. if the witch
	// attacks the assassin, it can set a new target to eliminate
	// other than the card that was witched)
	newTarget models.CharacterInterface
}

func (w *Witch) Resolve(options models.CharacterOptions) error {
	w.target = options.Target

	options.Target.BeWitch(w)
	w.target = options.Target

	options.ForPlayer = w.GetPlayer()
	// Resolve the target with new options
	options.Target = nil
	if err := w.target.Resolve(options); err != nil {
		return err
	}

	return nil
}

type Magistrate struct {
	Character
}

func (m *Magistrate) Resolve(options models.CharacterOptions) error {
	if options.Magistrate.UnsignedWarrant == nil && options.Magistrate.SignedWarrant == nil && options.Magistrate.WarrantMarkerBack == nil {
		return InvalidMagistrateTargetError
	}

	return nil
}
