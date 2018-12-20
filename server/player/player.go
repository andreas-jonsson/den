// DEN
// Copyright (C) 2018 Andreas T Jonsson

package player

import (
	"time"

	"gitlab.com/phix/den/message"
)

const (
	maxStamina = 50
	maxKeys    = 10
)

type Player struct {
	id    uint64
	alive bool

	respawn,
	lastMove time.Time

	lvl,
	keys,
	stamina,
	x, y int
}

func NewPlayer(id uint64) *Player {
	return &Player{
		id:      id,
		lvl:     1,
		keys:    maxKeys,
		alive:   true,
		stamina: maxStamina,
	}
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

func (p *Player) SetKeys(keys int) {
	p.keys = keys
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

func (p *Player) MoveTo(t byte, x, y int) bool {
	if p.stamina <= 0 || t == message.EmptyTile || t == message.WallTile {
		return false
	}

	if t == message.VDoorTile || t == message.HDoorTile {
		if p.keys > 0 {
			p.keys--
		} else {
			return false
		}
	}

	p.SetPosition(x, y)
	p.lastMove = time.Now()
	p.stamina--
	return true
}

func (p *Player) Stamina() int {
	return p.stamina
}

func (p *Player) Update() {
	if p.RespawnTime() == 0 {
		p.alive = true
	}
	if p.stamina < maxStamina && time.Since(p.lastMove) > time.Second/2 {
		p.lastMove = time.Now()
		p.stamina++
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
	p.keys = maxKeys
	if p.lvl > 1 {
		p.lvl /= 2
	}
}
