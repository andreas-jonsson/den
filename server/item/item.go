// DEN
// Copyright (C) 2018-2019 Andreas T Jonsson

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
