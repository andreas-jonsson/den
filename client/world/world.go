// DEN
// Copyright (C) 2018 Andreas T Jonsson

package world

import (
	"math"

	"gitlab.com/phix/den/message"
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

func (w *World) Level() []byte {
	return w.level
}

func (w *World) Index(x, y int) byte {
	if x >= w.size || y >= w.size || x < 0 || y < 0 {
		return message.EmptyTile
	}
	return w.level[y*w.size+x]
}

func (w *World) Rune(x, y int) rune {
	switch w.Index(x, y) {
	case message.EmptyTile:
		return ' '
	case message.WallTile:
		return '#'
	case message.FloorTile:
		return '.'
	default:
		return ' '
	}
}
