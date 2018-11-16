// DEN
// Copyright (C) 2018 Andreas T Jonsson

package server

import (
	"flag"
	"fmt"
	"net"
)

var listenPort uint

func init() {
	flag.UintVar(&listenPort, "port", 5000, "Listen for connections on specified port")
}

func Start() {
	lsock, err := net.Listen("tcp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		return
	}

	closeChan := make(chan struct{})
	for {
		conn, err := lsock.Accept()
		if err != nil {
			return
		}
		go serveConnection(conn, closeChan)
	}
}

func serveConnection(conn net.Conn, closeChan <-chan struct{}) {
	defer conn.Close()

	for {
		select {
		case _, ok := <-closeChan:
			if !ok {
				return
			}
		default:
		}
	}
}
