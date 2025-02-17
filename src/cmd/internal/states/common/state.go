package common

type State interface {
	Execute(context *Context)
}
