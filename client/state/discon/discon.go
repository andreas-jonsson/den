// DEN
// Copyright (C) 2018 Andreas T Jonsson

package discon

import (
	"github.com/nsf/termbox-go"
	"gitlab.com/phix/den/state"
)

const (
	LostConnectionMsg  = "Lost connection with server!"
	CouldNotConnectMsg = "Could not connect to server!"
)

const Name = "discon"

type Discon struct {
	m   state.Switcher
	msg string
}

func New(m state.Switcher) *Discon {
	return &Discon{m: m}
}

func (s *Discon) Name() string {
	return Name
}

func (s *Discon) Enter(m state.Switcher, from string, data ...interface{}) {
	if len(data) > 0 {
		if str, ok := data[0].(string); ok {
			s.msg = str
			return
		}
	}
	s.msg = LostConnectionMsg
}

func (s *Discon) Leave(to string) {}

func (s *Discon) Update() error {
events:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			s.m.Switch("exit")
			return nil
		case termbox.EventError:
			return ev.Err
		case termbox.EventInterrupt:
			break events
		}
	}

	w, h := termbox.Size()
	for i, r := range s.msg {
		termbox.SetCell(w/2-len(s.msg)/2+i, h/2, r, termbox.ColorDefault, termbox.ColorDefault)
	}
	return nil
}
