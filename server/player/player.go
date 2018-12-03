// DEN
// Copyright (C) 2018 Andreas T Jonsson

package player

type Player struct {
	id   uint64
	lvl  int
	x, y int
}

func NewPlayer(id uint64) *Player {
	return &Player{id: id}
}

func (p *Player) ID() uint64 {
	return p.id
}

func (p *Player) Level() int {
	return p.lvl
}

func (p *Player) Position() (int, int) {
	return p.x, p.y
}

func (p *Player) SetPosition(x, y int) {
	p.x = x
	p.y = y
}

func (p *Player) Update() {
}
