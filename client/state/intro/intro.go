// DEN
// Copyright (C) 2018 Andreas T Jonsson

package intro

import (
	"github.com/nsf/termbox-go"
	"gitlab.com/phix/den/client/state/exit"
	"gitlab.com/phix/den/state"
)

const Name = "intro"

type Intro struct {
	m state.Switcher
}

func New(m state.Switcher) *Intro {
	return &Intro{m}
}

func (s *Intro) Name() string {
	return Name
}

func (s *Intro) Enter(m state.Switcher, from string, data ...interface{}) {}

func (s *Intro) Leave(to string) {}

func (s *Intro) Update() error {
events:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			s.m.Switch(exit.Name)
		case termbox.EventError:
			return ev.Err
		case termbox.EventInterrupt:
			break events
		}
	}

	const (
		logoCenter = 19
		yOffset    = 4
	)

	w, h := termbox.Size()
	y := h/2 - len(logo)/2 - yOffset

	for i := 0; i < len(logo) && i+y < h; i++ {
		for j, r := range logo[i] {
			j += w/2 - logoCenter
			if j < w {
				termbox.SetCell(j, i+y, r, termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}
	return nil
}
