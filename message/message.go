// DEN
// Copyright (C) 2018 Andreas T Jonsson

package message

import (
	"encoding/gob"
)

type ClientConnect struct {
	Version [3]byte
	Name    string
}

const (
	NoMove byte = iota
	MoveUp
	MoveDown
	MoveLeft
	MoveRight
)

const (
	ActionA byte = 1 << iota
	ActionB
	ActionC
	ActionD
)

type ClientInput struct {
	Movement,
	Action byte
}

type ClientClose struct{}

const (
	EmptyTile byte = iota
	WallTile
	FloorTile
)

type ServerConnected struct {
	Result string
}

type ServerSetup struct {
	Id    uint64
	Level []byte
}

type ServerClose struct {
	Message string
}

type ServerMessage struct {
	Message string
}

type ServerCharacters struct {
	Characters []struct {
		Id       uint64
		Level    uint16
		Position [2]uint16
	}
}

func init() {
	gob.Register(ClientInput{})
	gob.Register(ClientClose{})

	gob.Register(ServerConnected{})
	gob.Register(ServerClose{})
	gob.Register(ServerMessage{})
	gob.Register(ServerCharacters{})
}
