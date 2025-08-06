package core

import (
	"sync"
)

type Starter interface {
	Init(state *State)
	Destroy(state *State)
}

func Boot(Starters ...Starter) *State {
	state := newState()
	state.SetState("@Starters", Starters)

	// 启动boot阶段
	state.startBoot(len(Starters))

	var g sync.WaitGroup
	for _, starter := range Starters {
		g.Add(1)
		go func() {
			defer g.Done()
			starter.Init(state)
		}()
	}

	g.Wait()

	// 结束boot阶段
	state.endBoot()

	return state
}

func (s *State) Destroy() {
	_Starters, ok := s.GetState("@Starters")
	if !ok {
		return
	}
	Starters, ok := _Starters.([]Starter)
	if !ok {
		return
	}
	for _, Starter := range Starters {
		Starter.Destroy(s)
	}
}
