package game

type AuthenticationActionsInterface[T any] interface {
	ParseMessage(message T) error
}
