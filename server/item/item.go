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

package item

type anyItem struct {
	id   uint64
	x, y int
}

func (it *anyItem) ID() uint64 {
	return it.id
}

func (it *anyItem) Position() (int, int) {
	return it.x, it.y
}

func (it *anyItem) SetPosition(x, y int) {
	it.x, it.y = x, y
}

func (it *anyItem) Update() {
}

type Key struct {
	anyItem
}

func NewKey(id uint64) *Key {
	return &Key{anyItem{id: id}}
}

type Potion struct {
	anyItem
}

func NewPotion(id uint64) *Potion {
	return &Potion{anyItem{id: id}}
}

type Levelup struct {
	anyItem
}

func NewLevelup(id uint64) *Levelup {
	return &Levelup{anyItem{id: id}}
}
