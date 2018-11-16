// DEN
// Copyright (C) 2018 Andreas T Jonsson

package message

const (
	MessageClose byte = 1 << iota

	// Client to server
	MessageConnect
	MessageInput

	// Server to client
	MessageSetup
	MessageUpdate
)

type Header byte

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

type Input struct {
	Header
	Movement,
	Action byte
}

type Connect struct {
	Header
	Name [128]byte
}

func (c *Connect) GetName() string {
	for i := 0; i < len(c.Name); i++ {
		if c.Name[i] == 0 {
			return string(c.Name[:i])
		}
	}
	return string(c.Name[:])
}
