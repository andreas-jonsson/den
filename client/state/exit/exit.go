// DEN
// Copyright (C) 2018 Andreas T Jonsson

package exit

import "gitlab.com/phix/den/state"

const Name = "exit"

type Exit struct {
}

func New() *Exit {
	return &Exit{}
}

func (s *Exit) Name() string {
	return Name
}

func (s *Exit) Enter(m state.Switcher, from string, data ...interface{}) {}

func (s *Exit) Leave(to string) {}

func (s *Exit) Update() error {
	return nil
}
