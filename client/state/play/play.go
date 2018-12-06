// DEN
// Copyright (C) 2018 Andreas T Jonsson

package play

import (
	"github.com/nsf/termbox-go"
	"gitlab.com/phix/den/client/connection"
	"gitlab.com/phix/den/client/state/discon"
	"gitlab.com/phix/den/client/state/exit"
	"gitlab.com/phix/den/client/world"
	"gitlab.com/phix/den/logger"
	"gitlab.com/phix/den/message"
	"gitlab.com/phix/den/state"
)

const Name = "play"

type Play struct {
	id         uint64
	posX, posY int
	wld        *world.World
	conn       *connection.Connection
	m          state.Switcher
}

func New(m state.Switcher) *Play {
	return &Play{m: m, posX: 1, posY: 1}
}

func (s *Play) Name() string {
	return Name
}

func (s *Play) Enter(m state.Switcher, from string, data ...interface{}) {
	s.conn = connection.Current
	s.id = s.conn.Setup().ID
	s.wld = s.conn.World()
}

func (s *Play) Leave(to string) {}

func (s *Play) Update() error {
events:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			posX := s.posX
			posY := s.posY
			move := message.NoMove

			switch ev.Key {
			case termbox.KeyEsc:
				s.m.Switch(exit.Name)
				return nil
			case termbox.KeyArrowUp:
				posY--
				move = message.MoveUp
			case termbox.KeyArrowDown:
				posY++
				move = message.MoveDown
			case termbox.KeyArrowLeft:
				posX--
				move = message.MoveLeft
			case termbox.KeyArrowRight:
				posX++
				move = message.MoveRight
			}

			if s.wld.Index(posX, posY) == message.FloorTile {
				s.posX = posX
				s.posY = posY

				if err := s.sendPosition(move); err != nil {
					s.m.Switch(discon.Name)
					return nil
				}
			}
		case termbox.EventError:
			return ev.Err
		case termbox.EventInterrupt:
			break events
		}
	}

	msg, err := s.conn.Decode()
	if err != nil {
		s.m.Switch(discon.Name)
		return nil
	}

	switch t := msg.(type) {
	case nil:
	case []message.ServerCharacter:
		s.wld.UpdateCharacters(t)
	default:
		logger.Fatalf("Invalid message: %T", msg)
	}

	w, h := termbox.Size()
	s.renderLevel(w, h)
	s.renderCharacters(w, h)

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

func (s *Play) renderCharacters(w, h int) {
	const playerLevel = 1

	cornerX := s.posX - w/2
	cornerY := s.posY - h/2

	for _, c := range s.wld.Characters() {
		if c.ID == s.id {
			termbox.SetCell(w/2, h/2, '@', termbox.ColorDefault, termbox.ColorDefault)
		} else {
			r := '0'
			switch {
			case c.Level > playerLevel:
				r = 'O'
			case c.Level < playerLevel:
				r = 'o'
			}
			termbox.SetCell(cornerX+int(c.PosX), cornerY+int(c.PosY), r, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func (s *Play) sendPosition(move byte) error {
	return s.conn.Encode(&message.ClientInput{move, 0})
}
