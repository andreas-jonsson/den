// DEN
// Copyright (C) 2018 Andreas T Jonsson

package server

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"gitlab.com/phix/den/game/world"
)

var listenPort uint

var wld *world.World

func init() {
	flag.UintVar(&listenPort, "port", 5000, "Listen for connections on specified port")
}

func Start() {
	lsock, err := net.Listen("tcp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	closeChan := make(chan struct{})

	defer wg.Wait()
	defer close(closeChan)

	for {
		conn, err := lsock.Accept()
		if err != nil {
			return
		}

		wg.Add(1)
		go serveConnection(conn, &wg, closeChan)
	}
}

func serveConnection(conn net.Conn, wg *sync.WaitGroup, closeChan <-chan struct{}) {
	defer wg.Done()
	defer conn.Close()

	//dec := gob.NewDecoder(conn)
	enc := gob.NewEncoder(conn)

	conn.SetDeadline(time.Now().Add(time.Second))
	if err := sendSetupData(enc); err != nil {
		return
	}

	for {
		select {
		case _, ok := <-closeChan:
			if ok {
				log.Fatalln("We should never receive on this channel")
			}
			return
		default:
		}

		conn.SetDeadline(time.Now().Add(time.Second))
	}
}

func sendSetupData(enc *gob.Encoder) error {
	if err := enc.Encode(wld.Level()); err != nil {
		return err
	}
	return nil
}
