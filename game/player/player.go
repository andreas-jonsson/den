// DEN
// Copyright (C) 2018 Andreas T Jonsson

package player

type Driver interface {
	ReadInput()
}

type Player struct {
	id     uint64
	lvl    int
	driver Driver

	x, y int
}

func NewPlayer(d Driver) *Player {
	return &Player{driver: d}
}

func (p *Player) Id() uint64 {
	return p.id
}

func (p *Player) Level() int {
	return p.lvl
}

func (p *Player) Position() (int, int) {
	return p.x, p.y
}

func (p *Player) Updata() {
	p.driver.ReadInput()
}
