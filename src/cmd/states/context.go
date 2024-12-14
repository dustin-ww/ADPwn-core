package states

type State interface {
	Execute(context *Context)
}

type Context struct {
	CurrentState State
}

func (c *Context) SetState(state State) {
	c.CurrentState = state
	if c.CurrentState != nil {
		c.CurrentState.Execute(c)
	}
}

func (c *Context) Execute() {
	if c.CurrentState != nil {
		c.CurrentState.Execute(c)
	}
}
