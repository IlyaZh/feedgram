package globalstate

import "sync/atomic"

type GlobalState struct {
	isFailed atomic.Bool
}

var globalState *GlobalState

func NewGlobalState() *GlobalState {
	if globalState != nil {
		return globalState
	}
	globalState = &GlobalState{}
	globalState.isFailed.Store(false)
	return globalState
}

func (c *GlobalState) SetFailed(state bool) {
	c.isFailed.Store(state)
}

func (c *GlobalState) IsFailed() bool {
	return c.isFailed.Load()
}
