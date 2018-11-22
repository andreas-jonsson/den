// DEN
// Copyright (C) 2018 Andreas T Jonsson

package connect

import (
	"encoding/gob"
	"net"

	"gitlab.com/phix/den/message"

	"github.com/nsf/termbox-go"
	"gitlab.com/phix/den/client/state/discon"
	"gitlab.com/phix/den/client/state/play"
	"gitlab.com/phix/den/logger"
	"gitlab.com/phix/den/state"
	"gitlab.com/phix/den/version"
)

const Name = "connect"

type Connect struct {
	m        state.Switcher
	host     string
	connChan chan net.Conn
}

func New(m state.Switcher, host string) *Connect {
	return &Connect{
		m:    m,
		host: host,
	}
}

func (s *Connect) Name() string {
	return Name
}

func (s *Connect) Enter(m state.Switcher, from string, data ...interface{}) {
	s.connChan = make(chan net.Conn, 1)
	go func() {
		conn, err := net.Dial("tcp", s.host)
		if err != nil {
			logger.Println("Could not connect to:", s.host)
			close(s.connChan)
			return
		}

		enc := gob.NewEncoder(conn)
		msg := message.ClientConnect{
			Name:    "noname",
			Version: [3]byte{version.Major, version.Minor, version.Patch},
		}

		if enc.Encode(&msg) != nil {
			logger.Println("Could not send connection request")
			close(s.connChan)
			return
		}
		s.connChan <- conn
	}()
}

func (s *Connect) Leave(to string) {
	//if s.connChan != nil {
	//<-s.connChan
	//}
}

func (s *Connect) Update() error {
events:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				s.m.Switch("intro")
				return nil
			}
		case termbox.EventError:
			return ev.Err
		case termbox.EventInterrupt:
			break events
		}
	}

	select {
	case conn, ok := <-s.connChan:
		if !ok {
			s.m.Switch(discon.Name, discon.CouldNotConnectMsg)
			return nil
		}
		s.connChan = nil
		s.m.Switch(play.Name, conn)
		return nil
	default:
	}

	const str = "Connecting to server..."

	w, h := termbox.Size()
	for i, r := range str {
		termbox.SetCell(w/2-len(str)/2+i, h/2, r, termbox.ColorDefault, termbox.ColorDefault)
	}
	return nil
}
