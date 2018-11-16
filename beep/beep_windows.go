// DEN
// Copyright (C) 2018 Andreas T Jonsson

package beep

import (
	"syscall"
	"time"
)

var beepProc = syscall.MustLoadDLL("kernel32.dll").MustFindProc("Beep")

func Beep(f int, d time.Duration) {
	beepProc.Call(uintptr(f), uintptr(d/time.Millisecond))
}
