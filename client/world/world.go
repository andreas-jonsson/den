// DEN
// Copyright (C) 2018 Andreas T Jonsson

package world

import (
	"math"

	"gitlab.com/phix/den/message"
)

type World struct {
	size  int
	level []byte

	characters []message.ServerCharacter
}

func NewWorld(level []byte) *World {
	size := int(math.Sqrt(float64(len(level))))
	return &World{
		size:  size,
		level: level,
	}
}

func (w *World) Level() []byte {
	return w.level
}

func (w *World) Characters() []message.ServerCharacter {
	return w.characters
}

func (w *World) UpdateCharacters(characters []message.ServerCharacter) {
	w.characters = characters
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
