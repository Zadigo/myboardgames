package models

type GameErrors string

func (e GameErrors) Error() string {
	return string(e)
}

const (
	PlayerChoiceError GameErrors = "You must either gather two golds or draw two districts, choose one and discard the other"
	BuildDistrictError GameErrors = "Cannot build district"
	ChoiceDoesNotExistsError GameErrors = "Choice does not exist"
	NoDistrictToBuild GameErrors = "No district to build"
)
