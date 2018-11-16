// DEN
// Copyright (C) 2018 Andreas T Jonsson

package world

import (
	"log"
	"math"

	"gitlab.com/phix/den/level"
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
		w.level = append(w.level, byte(r))
	}
	return w
}

func (w *World) Level() level.Level {
	return w.orgLevel
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
