// DEN
// Copyright (C) 2018 Andreas T Jonsson

package logger

import (
	"io"
	"io/ioutil"
	"log"
	"net"

	"github.com/nsf/termbox-go"
)

var (
	internalLogger *log.Logger
	connection     net.Conn
)

func Initialize(a interface{}) {
	switch ty := a.(type) {
	case io.Writer:
		internalLogger = log.New(ty, "", 0)
	case string:
		if ty == "" {
			goto handle_err
		}

		lsock, err := net.Listen("tcp", ":"+ty)
		if err != nil {
			goto handle_err
		}
		defer lsock.Close()

		connection, err = lsock.Accept()
		if err != nil {
			goto handle_err
		}
		internalLogger = log.New(connection, "", 0)
	default:
		goto handle_err
	}
	return

handle_err:
	internalLogger = log.New(ioutil.Discard, "", 0)
}

func Shutdown() {
	if connection != nil {
		connection.Close()
	}
}

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
	internalLogger.Fatal(a...)
}

func Fatalln(a ...interface{}) {
	termbox.Close()
	internalLogger.Fatalln(a...)
}

func Fatalf(format string, v ...interface{}) {
	termbox.Close()
	internalLogger.Fatalf(format, v...)
}
