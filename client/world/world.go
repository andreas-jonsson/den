// DEN
// Copyright (C) 2018-2019 Andreas T Jonsson

package world

import (
	"math"

	"gitlab.com/phix/den/message"
)

const (
	Visited byte = 1 << iota
	Visible
)

type World struct {
	size int
	level,
	flags []byte

	characters []message.ServerCharacter
	items      []message.ServerItem
}

func NewWorld(level []byte) *World {
	size := int(math.Sqrt(float64(len(level))))
	return &World{
		size:  size,
		level: level,
		flags: make([]byte, len(level)),
	}
}

func (w *World) ClearFlags() {
	for i := range w.flags {
		w.flags[i] = 0
	}
}

func (w *World) SetFlag(x, y int, f byte) {
	if x < w.size && y < w.size && x >= 0 && y >= 0 {
		w.flags[y*w.size+x] = f
	}
}

func (w *World) Flag(x, y int) byte {
	if x >= w.size || y >= w.size || x < 0 || y < 0 {
		return 0
	}
	return w.flags[y*w.size+x]
}

func (w *World) Characters() []message.ServerCharacter {
	return w.characters
}

func (w *World) Items() []message.ServerItem {
	return w.items
}

func (w *World) UpdateCharacters(characters []message.ServerCharacter) {
	w.characters = characters
}

func (w *World) UpdateItems(items []message.ServerItem) {
	w.items = items
}

func (w *World) Index(x, y int) byte {
	if x >= w.size || y >= w.size || x < 0 || y < 0 {
		return message.EmptyTile
	}
	return w.level[y*w.size+x]
}

func (w *World) Rune(x, y int) rune {
	return TileToRune(w.Index(x, y))
}

func TileToRune(t byte) rune {
	switch t {
	case message.EmptyTile:
		return ' '
	case message.WallTile:
		return '#'
	case message.FloorTile:
		return '.'
	case message.VDoorTile:
		return '|'
	case message.HDoorTile:
		return '='
	default:
		return ' '
	}
}
