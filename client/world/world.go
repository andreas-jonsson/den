// DEN
// Copyright (C) 2018 Andreas T Jonsson

package world

import (
	"math"

	"gitlab.com/phix/den/logger"
	"gitlab.com/phix/den/message"
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

func (w *World) Level() []byte {
	return w.level
}

func (w *World) Rune(x, y int) rune {
	if x > w.size || y > w.size || x < 0 || y < 0 {
		return ' '
	}

	tile := w.level[y*w.size+x]
	switch tile {
	case message.EmptyTile:
		return ' '
	case message.WallTile:
		return '#'
	case message.FloorTile:
		return '.'
	default:
		logger.Fatalln("Invalid tile")
		return 0
	}
}
