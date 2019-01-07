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

package discon

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
	"gitlab.com/phix/den/state"
)

const (
	LostConnectionMsg  = "Lost connection with server!"
	CouldNotConnectMsg = "Could not connect to:"
)

const Name = "discon"

type Discon struct {
	m         state.Switcher
	msg, host string
}

func New(m state.Switcher, host string) *Discon {
	return &Discon{m: m, host: host}
}

func (s *Discon) Name() string {
	return Name
}

func (s *Discon) Enter(m state.Switcher, from string, data ...interface{}) {
	if len(data) > 0 {
		if str, ok := data[0].(string); ok {
			if str == CouldNotConnectMsg {
				s.msg = fmt.Sprintf("%s %s", CouldNotConnectMsg, s.host)
				return
			}
			s.msg = str
			return
		}
	}
	s.msg = LostConnectionMsg
}

func (s *Discon) Leave(to string) {}

func (s *Discon) Update() error {
events:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			s.m.Switch("exit")
			return nil
		case termbox.EventError:
			return ev.Err
		case termbox.EventInterrupt:
			break events
		}
	}

	w, h := termbox.Size()
	for i, r := range s.msg {
		termbox.SetCell(w/2-len(s.msg)/2+i, h/2, r, termbox.ColorDefault|termbox.AttrReverse, termbox.ColorDefault|termbox.AttrReverse)
	}
	return nil
}
