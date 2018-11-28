// DEN
// Copyright (C) 2018 Andreas T Jonsson

package world

import (
	"log"
	"math"

	"gitlab.com/phix/den/level"
	"gitlab.com/phix/den/logger"
	"gitlab.com/phix/den/message"
)

type Unit interface {
	Id() uint64
	Position() (int, int)
	Update()
}

type World struct {
	size     int
	orgLevel level.Level
	level    []byte
	units    map[uint64]Unit
}

func NewWorld(l level.Level) *World {
	size := int(math.Sqrt(float64(len(l))))
	w := &World{size: size, orgLevel: l}
	for _, r := range l {
		w.level = append(w.level, runeToTile(r))
	}
	return w
}

func (w *World) Level() []byte {
	return w.level
}

func (w *World) Spawn(u Unit) {
	id := u.Id()
	if _, ok := w.units[id]; ok {
		log.Fatalln("Unit is already spawned: ", id)
	}
	w.units[id] = u
}

func (w *World) Update() {
	for _, u := range w.units {
		u.Update()
	}
}

func runeToTile(r rune) byte {
	switch r {
	case ' ':
		return message.EmptyTile
	case '#':
		return message.WallTile
	case '.':
		return message.FloorTile
	default:
		logger.Fatalln("Invalid tile:", r)
		return 0
	}
}
