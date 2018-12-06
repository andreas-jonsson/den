// DEN
// Copyright (C) 2018 Andreas T Jonsson

package connection

import (
	"encoding/gob"
	"errors"
	"net"
	"time"

	"gitlab.com/phix/den/client/world"
	"gitlab.com/phix/den/logger"
	"gitlab.com/phix/den/message"
)

const decodeChanSize = 128

var Current *Connection

type Connection struct {
	conn net.Conn
	enc  *gob.Encoder
	dec  *gob.Decoder
	wld  *world.World

	setup message.ServerSetup

	decodeChan chan interface{}
	closeChan  chan struct{}
}

func New(conn net.Conn, setup message.ServerSetup, wld *world.World) *Connection {
	c := &Connection{
		conn:       conn,
		enc:        gob.NewEncoder(conn),
		dec:        gob.NewDecoder(conn),
		setup:      setup,
		wld:        wld,
		decodeChan: make(chan interface{}, decodeChanSize),
		closeChan:  make(chan struct{}),
	}

	go c.startDecoder()
	return c
}

func (c *Connection) World() *world.World {
	return c.wld
}

func (c *Connection) Setup() message.ServerSetup {
	return c.setup
}

func (c *Connection) Encode(v interface{}) error {
	c.conn.SetWriteDeadline(time.Now().Add(time.Second))
	return c.enc.Encode(message.Wrap(v))
}

func (c *Connection) Decode() (interface{}, error) {
	select {
	case v, ok := <-c.decodeChan:
		if !ok {
			return nil, errors.New("connection was closed")
		}
		return v, nil
	default:
		return nil, nil
	}
}

func (c *Connection) Close() {
	c.closeChan <- struct{}{}
	c.conn.Close()
}

func (c *Connection) startDecoder() {
	var msg message.Any
	for {
		select {
		case <-c.closeChan:
			return
		default:
		}

		c.conn.SetReadDeadline(time.Now().Add(time.Second))
		if err := c.dec.Decode(&msg); err != nil {
			if ne, ok := err.(net.Error); ok && ne.Timeout() {
				continue
			}

			logger.Println(err)
			close(c.decodeChan)
			return
		}

		c.decodeChan <- msg.I
	}
}
