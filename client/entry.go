// DEN
// Copyright (C) 2018 Andreas T Jonsson

package client

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"

	"gitlab.com/phix/den/beep"
	"gitlab.com/phix/den/client/logger"
	"gitlab.com/phix/den/page437"
)

func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func draw(i int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	defer termbox.Flush()

	w, h := termbox.Size()
	s := fmt.Sprintf("count = %d", i)

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			termbox.SetCell(j, i, page437.ToUnicode(byte(i*16+j)), termbox.ColorRed, termbox.ColorDefault)
		}
	}

	tbPrint((w/2)-(len(s)/2), h/2, termbox.ColorRed, termbox.ColorDefault, s)
}

func Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer logger.Dump()
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)

	go func() {
		time.Sleep(5 * time.Second)

		logger.Fatalln("aaasas")

		beep.Beep(2000, 1*time.Second)

		time.Sleep(5 * time.Second)

		termbox.Interrupt()
	}()

	var count int
	draw(count)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == '+' {
				count++
			} else if ev.Ch == '-' {
				count--
			}

		case termbox.EventError:
			logger.Fatalln(ev.Err)

		case termbox.EventInterrupt:
			return
		}

		draw(count)
	}
}
