// DEN
// Copyright (C) 2018 Andreas T Jonsson

package logger

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/nsf/termbox-go"
)

var (
	buffer         bytes.Buffer
	internalLogger = log.New(&buffer, "", 0)
)

func Print(a ...interface{}) {
	internalLogger.Print(a...)
}

func Println(a ...interface{}) {
	internalLogger.Println(a...)
}

func Printf(format string, v ...interface{}) {
	internalLogger.Printf(format, v...)
}

func Fatal(a ...interface{}) {
	termbox.Close()
	internalLogger.Print(a...)
	Dump()
	os.Exit(1)
}

func Fatalln(a ...interface{}) {
	termbox.Close()
	internalLogger.Println(a...)
	Dump()
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	termbox.Close()
	internalLogger.Printf(format, v...)
	Dump()
	os.Exit(1)
}

func Dump() {
	if str := buffer.String(); len(str) > 0 {
		fmt.Print(str)
	}
}
