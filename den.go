// DEN
// Copyright (C) 2018 Andreas T Jonsson

//go:generate go run tools/version/version.go -file version/version.go

package main

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/phix/den/client"
	"gitlab.com/phix/den/server"
	"gitlab.com/phix/den/version"
)

// This sould only be enabled during development.
const includeServer = false

var (
	printVersion,
	printAbout,
	hostLocal bool
)

func init() {
	flag.BoolVar(&printVersion, "version", false, "Show version")
	flag.BoolVar(&printAbout, "about", false, "Show information about the game")

	if includeServer {
		flag.BoolVar(&hostLocal, "local", false, "Host local game")
	}
}

func main() {
	flag.Parse()

	if printAbout {
		fmt.Println("-=D=E=N=-")
		fmt.Println("\n", version.Copyright)
		fmt.Println("Contact: mail@andreasjonsson.se")
		fmt.Println("Version:", version.Full)
		return
	}

	if printVersion {
		fmt.Println(version.String)
		return
	}

	if hostLocal {
		flag.Set("host", "localhost:5000")
		flag.Parse()

		go func() {
			<-client.LoggerInitializedChan
			server.Start()
		}()

		defer func() {
			server.InterruptChan <- os.Interrupt
			<-server.ServerExitedChan
		}()
	}

	client.Start()
}
