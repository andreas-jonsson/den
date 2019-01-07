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

package logger

import (
	"io"
	"io/ioutil"
	"log"
	"net"

	termbox "github.com/nsf/termbox-go"
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

func IsInitialized() bool {
	return internalLogger != nil
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
