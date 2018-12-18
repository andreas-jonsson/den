// DEN
// Copyright (C) 2018 Andreas T Jonsson

package play

import (
	"fmt"
	"math"

	termbox "github.com/nsf/termbox-go"
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

	playerLevel,
	respawn,
	keys int
	alive bool

	wld  *world.World
	conn *connection.Connection
	m    state.Switcher
}

func New(m state.Switcher) *Play {
	return &Play{
		m:           m,
		posX:        1,
		posY:        1,
		playerLevel: 1,
		alive:       true,
	}
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

			if s.alive && s.wld.Index(posX, posY) == message.FloorTile {
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
	s.renderUI(w, h)

	return nil
}

func (s *Play) renderLevel(w, h int) {
	cornerX := s.posX - w/2
	cornerY := s.posY - h/2

	for y := 0; y < w; y++ {
		wY := cornerY + y
		for x := 0; x < w; x++ {
			wX := cornerX + x
			t := s.wld.Index(wX, wY)
			if t == message.EmptyTile {
				continue
			}

			flags := s.wld.Flag(wX, wY)
			vx := float64(s.posX - wX)
			vy := float64(s.posY - wY)
			l := math.Sqrt((vx * vx) + (vy * vy))

			const fogDist = 24
			if l > fogDist {
				flags &= ^world.Visited
			}

			if (t == message.WallTile && flags&world.Visited != 0) || s.calculateFov(wX, wY) {
				s.wld.SetFlag(wX, wY, flags|world.Visited|world.Visible)
				termbox.SetCell(x, y, world.TileToRune(t), termbox.ColorDefault, termbox.ColorDefault)
			} else {
				// Remove visible flag.
				s.wld.SetFlag(wX, wY, flags&^world.Visible)
			}
		}
	}
}

func (s *Play) renderCharacters(w, h int) {
	cornerX := s.posX - w/2
	cornerY := s.posY - h/2

	for _, c := range s.wld.Characters() {
		alive := c.Respawn == 0
		if c.ID == s.id {
			// TODO: Sync position if we get to much out of sync.
			//s.posX = int(c.PosX)
			//s.posY = int(c.PosY)

			// This is a hack!
			if !s.alive && alive {
				s.posX, s.posY = 1, 1
			}

			s.alive = alive
			s.respawn = int(c.Respawn)
			s.keys = int(c.Keys)
			s.playerLevel = int(c.Level)
		} else if alive {
			viewX := int(c.PosX) - cornerX
			viewY := int(c.PosY) - cornerY

			if s.wld.Flag(int(c.PosX), int(c.PosY))&world.Visible == 0 {
				continue
			}

			r := '0'
			switch {
			case int(c.Level) > s.playerLevel:
				r = 'O'
			case int(c.Level) < s.playerLevel:
				r = 'o'
			}
			termbox.SetCell(viewX, viewY, r, termbox.ColorDefault, termbox.ColorDefault)
		}
	}

	if s.alive {
		termbox.SetCell(w/2, h/2, '@', termbox.ColorDefault, termbox.ColorDefault)
	}
}

func (s *Play) sendPosition(move byte) error {
	return s.conn.Encode(&message.ClientInput{Movement: move})
}

func (s *Play) calculateFov(x, y int) bool {
	vx := float64(s.posX - x)
	vy := float64(s.posY - y)

	l := math.Sqrt((vx * vx) + (vy * vy))

	const minViewDist = 1
	const maxViewDist = 12

	if l > maxViewDist {
		return false
	}

	vx /= l
	vy /= l

	ox := float64(x) + vx + 0.5
	oy := float64(y) + vy + 0.5

	for i := 0; i < int(l)-minViewDist; i++ {
		t := s.wld.Index(int(ox), int(oy))
		if t == message.WallTile {
			return false
		}
		ox += vx
		oy += vy
	}
	return true
}

func (s *Play) renderUI(w, h int) {
	print(0, 0, true, fmt.Sprintf("Level: %d", s.playerLevel))
	print(16, 0, true, fmt.Sprintf("Keys: %d", s.keys))
	print(32, 0, true, fmt.Sprintf("Players: %d", len(s.wld.Characters())))

	if !s.alive {
		msg := "   You are dead!   "
		print(w/2-len(msg)/2, h/2-1, true, msg)
		msg = fmt.Sprintf("Respawn in %d...", s.respawn)
		print(w/2-len(msg)/2, h/2+1, true, msg)
	}
}

func print(x, y int, inv bool, msg string) {
	attrib := termbox.ColorDefault
	if inv {
		attrib |= termbox.AttrReverse
	}
	for i, r := range msg {
		termbox.SetCell(x+i, y, r, attrib, attrib)
	}
}
