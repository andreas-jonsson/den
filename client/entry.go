// DEN
// Copyright (C) 2018 Andreas T Jonsson

package client

import (
	"flag"
	"time"

	"github.com/nsf/termbox-go"

	"gitlab.com/phix/den/client/state/exit"
	"gitlab.com/phix/den/client/state/intro"
	"gitlab.com/phix/den/client/state/play"
	"gitlab.com/phix/den/logger"
	"gitlab.com/phix/den/state"
)

var logPort string

func init() {
	flag.StringVar(&logPort, "tcplog", "", "Port for TCP logger")
}

func Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	logger.Initialize(logPort)
	defer logger.Shutdown()

	termbox.SetInputMode(termbox.InputEsc)

	m := state.NewMachine()
	m.AddState(exit.New())

	m.AddState(intro.New(m))
	m.SetState(intro.Name)

	m.AddState(play.New(m))

	ticker := time.NewTicker(time.Second / 30)
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
