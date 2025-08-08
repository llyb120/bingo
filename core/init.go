package core

import (
	"github.com/llyb120/yoya2/y"
)

type Starter func() func()

func Boot(Starters ...Starter) {
	state := newState()
	state.SetState("@Starters", Starters)

	// 启动boot阶段
	state.startBoot(len(Starters))

	finalizers := y.Flex(Starters, func(starter Starter, _ int) any {
		ender := starter()
		return ender
	}, y.UseAsync, y.UsePanic, y.NotNil)

	state.SetState("@finalizers", finalizers)

	// 结束boot阶段
	state.endBoot()

}

func Shutdown() {
	_finalizers, ok := globalState.GetState("@finalizers")
	if !ok {
		return
	}
	finalizers, ok := _finalizers.([]func())
	if !ok {
		return
	}
	for _, ender := range finalizers {
		ender()
	}
}
