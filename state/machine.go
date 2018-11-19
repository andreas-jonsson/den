// DEN
// Copyright (C) 2018 Andreas T Jonsson

package state

import "gitlab.com/phix/den/logger"

type Switcher interface {
	Switch(to string, data ...interface{})
}

type State interface {
	Name() string
	Enter(m Switcher, from string, data ...interface{})
	Leave(to string)
	Update() error
}

type Machine struct {
	states       map[string]State
	currentState State
}

func NewMachine() *Machine {
	return &Machine{states: map[string]State{}}
}

func (m *Machine) SetState(name string) {
	m.currentState = m.states[name]
}

func (m *Machine) AddState(s State) {
	m.states[s.Name()] = s
}

func (m *Machine) CurrentState() State {
	return m.currentState
}

func (m *Machine) States() map[string]State {
	return m.states
}

func (m *Machine) Update() error {
	return m.currentState.Update()
}

func (m *Machine) Switch(to string, data ...interface{}) {
	//logger.Println(m.currentState.Name(), "->", to)

	old := m.currentState.Name()
	m.currentState.Leave(to)
	m.currentState = m.states[to]
	if m.currentState == nil {
		logger.Fatalln("Invalid state:", to)
	}
	m.currentState.Enter(m, old, data...)
}
