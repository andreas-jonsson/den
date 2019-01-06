// DEN
// Copyright (C) 2018-2019 Andreas T Jonsson

package beep

import (
	"os"
	"syscall"
	"time"
)

const kdKIOCSOUND = 0x4B2F

var commandChan = make(chan func() time.Duration, 256)

func init() {
	go func() {
		f := <-commandChan
		for {
			t := time.NewTimer(f())
			select {
			case <-t.C:
				syscall.Syscall(syscall.SYS_IOCTL, os.Stdout.Fd(), uintptr(kdKIOCSOUND), uintptr(0))
			case f = <-commandChan:
			}
			t.Stop()
		}
	}()
}

func Beep(f int, d time.Duration) {
	commandChan <- func() time.Duration {
		var freq uintptr
		if f > 0 {
			freq = uintptr(1193180 / f)
		}
		syscall.Syscall(syscall.SYS_IOCTL, os.Stdout.Fd(), uintptr(kdKIOCSOUND), freq)
		return d
	}
}
