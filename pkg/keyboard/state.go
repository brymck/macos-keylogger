package keyboard

type State uint8

const (
	InvalidState State = iota
	Down
	Up
)
