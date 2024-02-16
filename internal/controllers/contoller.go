package controllers

import (
	"github.com/IlyaZh/feedsgram/internal/components/global_state"
)

type Controller struct {
	globalState *globalstate.GlobalState
}

func NewPublicApi() *Controller {
	return &Controller{
		globalState: globalstate.NewGlobalState(),
	}
}
