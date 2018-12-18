// DEN
// Copyright (C) 2018 Andreas T Jonsson

package player

type Player struct {
	id        uint64
	lvl, keys int
	x, y      int
	alive     bool
}

func NewPlayer(id uint64) *Player {
	return &Player{id: id, lvl: 1, alive: true}
}

func (p *Player) ID() uint64 {
	return p.id
}

func (p *Player) Level() int {
	return p.lvl
}

func (p *Player) Keys() int {
	return p.keys
}

func (p *Player) SetLevel(lvl int) {
	p.lvl = lvl
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

func (p *Player) Alive() bool {
	return p.alive
}

func (p *Player) Die() {
	p.alive = false
	if p.lvl > 1 {
		p.lvl /= 2
	}
}
