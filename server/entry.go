// DEN
// Copyright (C) 2018 Andreas T Jonsson

package server

import (
	"encoding/gob"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"gitlab.com/phix/den/level"
	"gitlab.com/phix/den/logger"
	"gitlab.com/phix/den/message"
	"gitlab.com/phix/den/server/player"
	"gitlab.com/phix/den/server/world"
	"gitlab.com/phix/den/version"
)

var (
	SocketOpenChan   = make(chan struct{}, 1)
	ServerExitedChan = make(chan struct{}, 1)
	InterruptChan    = make(chan os.Signal, 1)
)

var listenPort uint

var wld *world.World

func init() {
	flag.UintVar(&listenPort, "port", 5000, "Listen for connections on specified port")
}

func Start() {
	defer func() {
		ServerExitedChan <- struct{}{}
	}()

	if !logger.IsInitialized() {
		logger.Initialize(os.Stdout)
	}

	lsock, err := net.Listen("tcp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		return
	}
	defer lsock.Close()
	SocketOpenChan <- struct{}{}

	var wg sync.WaitGroup
	closeChan := make(chan struct{})

	defer wg.Wait()
	defer close(closeChan)

	wld = world.NewWorld(level.Level1)
	go wld.StartUpdate()

	logger.Println("Server started!")

	var playerID uint64
	for {
		lsock.(*net.TCPListener).SetDeadline(time.Now().Add(time.Second))

		conn, err := lsock.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Timeout() {
				select {
				case <-InterruptChan:
					return
				default:
					continue
				}
			}

			logger.Println("Error in TCP listen:", err)
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

	logger.Println("Player connected:", id)

	dec := gob.NewDecoder(conn)
	enc := gob.NewEncoder(conn)

	conn.SetDeadline(time.Now().Add(time.Second))

	var msg message.ClientConnect
	if err := dec.Decode(&msg); err != nil {
		logger.Println("Handshake failed", err)
		return
	}

	conn.SetDeadline(time.Now().Add(time.Second))
	if err := sendSetupData(enc, id, msg); err != nil {
		logger.Println("Player initialization failed:", err)
		return
	}

	wld.Send(func(w *world.World) {
		p := player.NewPlayer(id)
		p.SetPosition(1, 1)
		w.Spawn(p)
	})

	defer wld.Send(func(w *world.World) {
		w.Unspawn(w.Unit(id))
	})

	charactertTimer := time.NewTicker(time.Second / 10)
	defer charactertTimer.Stop()

	messageQueue := make(chan func() error, 128)

	for {
		select {
		case _, ok := <-closeChan:
			if ok {
				logger.Fatalln("We should never receive on this channel")
			}
			return
		case f := <-messageQueue:
			if err := f(); err != nil {
				logger.Println(err)
				return
			}
		case <-charactertTimer.C:
			wld.Send(func(w *world.World) {
				var characters []message.ServerCharacter
				for uid, u := range w.Units() {
					c, ok := u.(world.Character)
					if !ok {
						continue
					}

					x, y := c.Position()
					cmsg := message.ServerCharacter{
						ID:    uid,
						Level: int16(c.Level()),
						PosX:  int16(x),
						PosY:  int16(y),
						Alive: c.Alive(),
					}

					if uid == id {
						cmsg.Keys = int16(c.Keys())
					}
					characters = append(characters, cmsg)
				}

				messageQueue <- func() error {
					conn.SetWriteDeadline(time.Now().Add(time.Second))
					if err := enc.Encode(message.Wrap(characters)); err != nil {
						return err
					}
					return nil
				}
			})
		default:
		}

		conn.SetReadDeadline(time.Now().Add(time.Millisecond))
		var msg message.Any
		if err := dec.Decode(&msg); err != nil {
			if ne, ok := err.(net.Error); ok && ne.Timeout() {
				continue
			}
			logger.Println(err)
			return
		}

		switch t := msg.I.(type) {
		case message.ClientInput:
			wld.Send(func(w *world.World) {
				c := w.Unit(id).(world.Character)
				x, y := c.Position()
				nx, ny := 0, 0

				switch t.Movement {
				case message.MoveUp:
					nx, ny = x, y-1
				case message.MoveDown:
					nx, ny = x, y+1
				case message.MoveLeft:
					nx, ny = x-1, y
				case message.MoveRight:
					nx, ny = x+1, y
				default:
					return
				}

				if nx >= 0 && ny >= 0 && w.Index(nx, ny) == message.FloorTile {
					c.SetPosition(nx, ny)

					for id, otherUnit := range w.Units() {
						if id == c.ID() {
							continue
						}

						otherCharacter, ok := otherUnit.(world.Character)
						if !ok || !otherCharacter.Alive() {
							continue
						}

						x, y := otherCharacter.Position()
						if x == nx && y == ny {
							playerLevel := c.Level()
							otherLevel := otherCharacter.Level()

							switch {
							case playerLevel > otherLevel:
								otherCharacter.Die()
								c.SetLevel(playerLevel + 1)
							case playerLevel < otherLevel:
								c.Die()
							default:
								if time.Now().UnixNano()%2 == 0 {
									otherCharacter.Die()
								} else {
									c.Die()
								}
							}
						}
					}
				}
			})
		default:
			logger.Printf("Invalid message type: %T", msg.I)
		}
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

	if err := enc.Encode(&resp); err != nil {
		return err
	}

	var setup message.ServerSetup
	setup.ID = id
	setup.Level = wld.Level()
	return enc.Encode(&setup)
}
