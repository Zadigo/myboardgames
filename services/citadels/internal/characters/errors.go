package characters

type CharacterError string

func (e CharacterError) Error() string {
	return string(e)
}

const (
	CannotAttackSelfError     CharacterError = "Cannot attack self"
	CannotStealLevelOneError  CharacterError = "Cannot steal level 1 character"
	CannotStealKilledError    CharacterError = "Cannot steal killed character"
	CannotStealBewitchedError CharacterError = "Cannot steal bewitched character"

	TargetHasNoCharacterError CharacterError = "Target has no character"
	AmbiguousTargetError      CharacterError = "Ambiguous target. Either set Target or ReplacedTarget"
	MissingTargetError        CharacterError = "Target cannot be nil for this character"

	InvalidMagistrateTargetError CharacterError = "Invalid magistrate target. Set UnsignedWarrant, SignedWarrant and WarrantMarkerBack"

	InvalidMagicianChoiceError CharacterError = "Magician can either discard or exchange cards"
)
