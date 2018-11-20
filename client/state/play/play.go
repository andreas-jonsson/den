// DEN
// Copyright (C) 2018 Andreas T Jonsson

package play

import (
	"github.com/nsf/termbox-go"
	"gitlab.com/phix/den/client/state/exit"
	"gitlab.com/phix/den/state"
)

const Name = "play"

type Play struct {
	posX, posY int
	m          state.Switcher
}

func New(m state.Switcher) *Play {
	return &Play{m: m}
}

func (s *Play) Name() string {
	return Name
}

func (s *Play) Enter(m state.Switcher, from string, data ...interface{}) {}

func (s *Play) Leave(to string) {}

func (s *Play) Update() error {
events:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				s.m.Switch(exit.Name)
			case termbox.KeyArrowUp:
				s.posY--
			case termbox.KeyArrowDown:
				s.posY++
			case termbox.KeyArrowLeft:
				s.posX--
			case termbox.KeyArrowRight:
				s.posX++
			}
		case termbox.EventError:
			return ev.Err
		case termbox.EventInterrupt:
			break events
		}
	}

	s.renderLevel()

	w, h := termbox.Size()
	termbox.SetCell(w/2, h/2, '@', termbox.ColorDefault, termbox.ColorDefault)
	return nil
}

func (s *Play) renderLevel() {

}
