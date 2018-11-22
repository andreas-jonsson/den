// DEN
// Copyright (C) 2018 Andreas T Jonsson

package server

import (
	"encoding/gob"
	"errors"
	"flag"
	"fmt"
	"net"
	"sync"
	"time"

	"gitlab.com/phix/den/logger"
	"gitlab.com/phix/den/message"
	"gitlab.com/phix/den/server/world"
	"gitlab.com/phix/den/version"
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

	var playerID uint64
	for {
		conn, err := lsock.Accept()
		if err != nil {
			return
		}

		playerID++
		wg.Add(1)
		go serveConnection(conn, &wg, closeChan, playerID)
	}
}

func serveConnection(conn net.Conn, wg *sync.WaitGroup, closeChan <-chan struct{}, id uint64) {
	defer wg.Done()
	defer conn.Close()

	dec := gob.NewDecoder(conn)
	enc := gob.NewEncoder(conn)

	conn.SetDeadline(time.Now().Add(time.Second))

	var msg message.ClientConnect
	if err := dec.Decode(&msg); err != nil {
		logger.Println(err)
	}

	conn.SetDeadline(time.Now().Add(time.Second))
	if err := sendSetupData(enc, id, msg); err != nil {
		return
	}

	for {
		select {
		case _, ok := <-closeChan:
			if ok {
				logger.Fatalln("We should never receive on this channel")
			}
			return
		default:
		}

		conn.SetDeadline(time.Now().Add(time.Second))
	}
}

func sendSetupData(enc *gob.Encoder, id uint64, msg message.ClientConnect) error {
	var resp message.ServerConnected
	if msg.Version[0] != version.Major || msg.Version[1] != version.Minor {
		err := errors.New("Invalid version. Server is running: " + version.String)
		resp.Result = err.Error()
		enc.Encode(&resp)
		return err
	}

	resp.Id = id
	//resp.Level
	return enc.Encode(&resp)
}
