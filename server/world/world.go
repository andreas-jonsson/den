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
	"log"
	"math"

	"gitlab.com/phix/den/level"
	"gitlab.com/phix/den/logger"
)

type Unit interface {
	ID() uint64
	Position() (int, int)
	SetPosition(x, y int)
	Update()
}

type World struct {
	size     int
	orgLevel level.Level
	level    []byte
	jobs     chan func(*World)
	units    map[uint64]Unit
}

func NewWorld(l level.Level) *World {
	size := int(math.Sqrt(float64(len(l))))
	w := &World{
		size:     size,
		orgLevel: l,
		jobs:     make(chan func(w *World), 128),
		units:    make(map[uint64]Unit),
	}

	for _, r := range l {
		w.level = append(w.level, r)
	}
	return w
}

func (w *World) Size() int {
	return w.size
}

func (w *World) Index(x, y int) byte {
	return w.level[y*w.size+x]
}

func (w *World) Level() []byte {
	return w.level
}

func (w *World) Unit(id uint64) Unit {
	u, ok := w.units[id]
	if !ok {
		logger.Fatalln("Could not find unit:", id)
	}
	return u
}

func (w *World) Units() map[uint64]Unit {
	return w.units
}

func (w *World) Send(f func(*World)) {
	w.jobs <- f
}

func (w *World) StartUpdate() {
	for f := range w.jobs {
		f(w)
	}
}

func (w *World) Spawn(u Unit) {
	id := u.ID()
	if _, ok := w.units[id]; ok {
		log.Fatalln("Unit is already spawned:", id)
	}
	w.units[id] = u
}

func (w *World) Unspawn(u Unit) {
	id := u.ID()
	if _, ok := w.units[id]; !ok {
		log.Fatalln("Unit is not spawned:", id)
	}
	delete(w.units, id)
}

func (w *World) Update() {
	for _, u := range w.units {
		u.Update()
	}
}
