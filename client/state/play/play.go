// DEN
// Copyright (C) 2018 Andreas T Jonsson

package play

import (
	"encoding/gob"
	"net"
	"time"

	"github.com/nsf/termbox-go"
	"gitlab.com/phix/den/client/state/exit"
	"gitlab.com/phix/den/client/world"
	"gitlab.com/phix/den/message"
	"gitlab.com/phix/den/state"
)

const Name = "play"

type Play struct {
	id         uint64
	posX, posY int
	wld        *world.World
	conn       net.Conn
	enc        *gob.Encoder
	dec        *gob.Decoder
	m          state.Switcher
}

func New(m state.Switcher) *Play {
	return &Play{m: m}
}

func (s *Play) SetupConnection(conn net.Conn) error {
	s.conn = conn
	s.enc = gob.NewEncoder(conn)
	s.dec = gob.NewDecoder(conn)

	var setup message.ServerSetup

	s.conn.SetDeadline(time.Now().Add(time.Second))
	if err := s.enc.Encode(&setup); err != nil {
		return nil
	}

	s.id = setup.Id
	s.wld = world.NewWorld(setup.Level)
	return nil
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
				return nil
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

	w, h := termbox.Size()
	s.renderLevel(w, h)

	termbox.SetCell(w/2, h/2, '@', termbox.ColorDefault, termbox.ColorDefault)
	return nil
}

func (s *Play) renderLevel(w, h int) {
	cornerX := s.posX - w/2
	cornerY := s.posY - h/2

	for y := 0; y < w; y++ {
		for x := 0; x < w; x++ {
			termbox.SetCell(x, y, s.wld.Rune(cornerX+x, cornerY+y), termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}
