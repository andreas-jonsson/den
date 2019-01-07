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

package player

import (
	"time"

	"gitlab.com/phix/den/message"
)

const MaxStamina = 50

const initialKeys = 1

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
		keys:    initialKeys,
		alive:   true,
		stamina: MaxStamina,
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

func (p *Player) SetStamina(s int) {
	p.stamina = s
}

func (p *Player) Update() {
	if p.RespawnTime() == 0 {
		p.alive = true
	}
	if p.stamina < MaxStamina && time.Since(p.lastMove) > time.Second/2 {
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
	if p.lvl > 1 {
		p.lvl /= 2
	}
}
