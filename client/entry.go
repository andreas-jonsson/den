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

package client

import (
	"flag"
	"time"

	termbox "github.com/nsf/termbox-go"

	"gitlab.com/phix/den/client/state/connect"
	"gitlab.com/phix/den/client/state/discon"
	"gitlab.com/phix/den/client/state/exit"
	"gitlab.com/phix/den/client/state/intro"
	"gitlab.com/phix/den/client/state/play"
	"gitlab.com/phix/den/logger"
	"gitlab.com/phix/den/state"
)

var (
	LoggerInitializedChan = make(chan struct{}, 1)
	GameExitedChan        = make(chan struct{}, 1)
)

var (
	logPort,
	hostAddr string
)

func init() {
	flag.StringVar(&logPort, "tcplog", "", "Port for TCP logger")
	flag.StringVar(&hostAddr, "host", "den-pub.andreasjonsson.se:5000", "Connect to server")
}

func Start() {
	defer func() {
		GameExitedChan <- struct{}{}
	}()

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	logger.Initialize(logPort)
	LoggerInitializedChan <- struct{}{}
	defer logger.Shutdown()

	termbox.SetInputMode(termbox.InputEsc)

	m := state.NewMachine()
	m.AddState(exit.New())

	m.AddState(intro.New(m, hostAddr))
	m.SetState(intro.Name)

	m.AddState(play.New(m, hostAddr))
	m.AddState(discon.New(m, hostAddr))
	m.AddState(connect.New(m, hostAddr))

	ticker := time.NewTicker(time.Second / 15)
	go func() {
		for range ticker.C {
			termbox.Interrupt()
		}
	}()

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		if err := m.Update(); err != nil {
			logger.Fatalln(err)
		}
		if m.CurrentState().Name() == exit.Name {
			ticker.Stop()
			return
		}
		termbox.Flush()
	}
}
