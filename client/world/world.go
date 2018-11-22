// DEN
// Copyright (C) 2018 Andreas T Jonsson

package world

import (
	"math"

	"gitlab.com/phix/den/level"
)

const (
	Ground = iota
	Wall
)

type Character struct {
	Id       uint64
	Level    uint16
	Position [2]uint16
}

type World struct {
	size  int
	level []byte

	characters map[uint64]Character
}

func NewWorld(level []byte) *World {
	size := int(math.Sqrt(float64(len(level))))
	return &World{
		size:       size,
		level:      level,
		characters: make(map[uint64]Character),
	}
}

func (w *World) Level() level.Level {
	return w.orgLevel
}
