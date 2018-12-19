// DEN
// Copyright (C) 2018 Andreas T Jonsson

package intro

import (
	"time"

	termbox "github.com/nsf/termbox-go"
	"gitlab.com/phix/den/client/state/connect"
	"gitlab.com/phix/den/client/state/exit"
	"gitlab.com/phix/den/state"
)

const displayTime = 3 * time.Second

const Name = "intro"

type Intro struct {
	m    state.Switcher
	t    *time.Timer
	host string
}

func New(m state.Switcher, host string) *Intro {
	return &Intro{m,
		time.NewTimer(displayTime),
		host,
	}
}

func (s *Intro) Name() string {
	return Name
}

func (s *Intro) Enter(m state.Switcher, from string, data ...interface{}) {
	s.t.Reset(displayTime)
}

func (s *Intro) Leave(to string) {}

func (s *Intro) Update() error {
events:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				s.m.Switch(exit.Name)
				return nil
			}
			s.m.Switch(connect.Name)
			return nil
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
			termbox.SetCell(j, i+y, r, termbox.ColorDefault, termbox.ColorDefault)
		}
	}

	for i, r := range s.host {
		termbox.SetCell(i, h-1, r, termbox.ColorDefault|termbox.AttrReverse, termbox.ColorDefault|termbox.AttrReverse)
	}

	select {
	case <-s.t.C:
		s.m.Switch(connect.Name)
	default:
	}
	return nil
}
