// DEN
// Copyright (C) 2018-2019 Andreas T Jonsson

package message

import (
	"encoding/gob"
)

func Wrap(v interface{}) *Any {
	return &Any{I: v}
}

type Any struct {
	I interface{}
}

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

type ClientInput struct {
	Movement byte
}

type ClientClose struct{}

const (
	EmptyTile byte = iota
	WallTile
	FloorTile
	VDoorTile
	HDoorTile
)

type ServerConnected struct {
	Result string
}

type ServerSetup struct {
	ID    uint64
	Level []byte
}

type ServerClose struct {
	Message string
}

type ServerMessage struct {
	Message string
}

type ServerCharacter struct {
	ID uint64
	PosX,
	PosY int16
	Level,
	Keys,
	Stamina,
	Respawn byte
}

const (
	KeyItem byte = iota
	PotionItem
	LevelupItem
)

type ServerItem struct {
	ID   uint64
	Type byte
	PosX,
	PosY int16
}

func init() {
	gob.Register(ClientInput{})
	gob.Register(ClientClose{})

	gob.Register(ServerConnected{})
	gob.Register(ServerClose{})
	gob.Register(ServerMessage{})
	gob.Register([]ServerCharacter{})
	gob.Register([]ServerItem{})
}
