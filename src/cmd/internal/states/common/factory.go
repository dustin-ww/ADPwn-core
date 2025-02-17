package common

type StateFactory interface {
	CreateMainMenuState(project interface{}) State
	CreateAddHostRangeState(project interface{}) State
}
