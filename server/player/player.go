// DEN
// Copyright (C) 2018 Andreas T Jonsson

package player

import "time"

type Player struct {
	id        uint64
	lvl, keys int
	x, y      int
	alive     bool
	respawn   time.Time
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
	p.x, p.y = x, y
}

func (p *Player) Update() {
	if p.RespawnTime() == 0 {
		p.alive = true
	}
}

func (p *Player) Alive() bool {
	return p.alive
}

func (p *Player) RespawnTime() int {
	if p.alive {
		return 0
	}
	t := p.lvl*2 - int(time.Since(p.respawn)/time.Second)
	if t <= 0 {
		return 0
	}
	return t
}

func (p *Player) Die() {
	p.alive = false
	p.respawn = time.Now()
	if p.lvl > 1 {
		p.lvl /= 2
	}

	p.x, p.y = 1, 1
}
