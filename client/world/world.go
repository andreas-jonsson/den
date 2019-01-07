/*
Copyright (C) 2018-2019 Andreas T Jonsson

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
