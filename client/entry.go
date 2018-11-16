// DEN
// Copyright (C) 2018 Andreas T Jonsson

package client

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
	"gitlab.com/phix/terminology/beep"
	"gitlab.com/phix/terminology/page437"
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
	termbox.SetInputMode(termbox.InputEsc)

	go func() {
		time.Sleep(5 * time.Second)

		beep.Beep(2000, 1*time.Second)

		time.Sleep(5 * time.Second)

		termbox.Interrupt()

		// This should never run - the Interrupt(), above, should cause the event
		// loop below to exit, which then exits the process.  If something goes
		// wrong, this panic will trigger and show what happened.
		time.Sleep(1 * time.Second)
		panic("this should never run")
	}()

	var count int
	draw(count)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == '+' {
				count++
			} else if ev.Ch == '-' {
				count--
			}

		case termbox.EventError:
			panic(ev.Err)

		case termbox.EventInterrupt:
			break mainloop
		}

		draw(count)
	}
	termbox.Close()
}
