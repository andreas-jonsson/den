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

var (
	printVersion,
	printAbout,
	hostLocal bool
)

func init() {
	flag.BoolVar(&printVersion, "version", false, "Show version")
	flag.BoolVar(&printAbout, "about", false, "Show information about the game")
	flag.BoolVar(&hostLocal, "local", false, "Host local game")
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

	interruptChan := make(chan os.Signal)
	if hostLocal {
		flag.Set("host", "localhost:5000")
		flag.Parse()

		server.InterruptChan = interruptChan
		go func() {
			<-client.LoggerInitializedChan
			server.Start()
		}()
	}

	client.Start()

	if hostLocal {
		interruptChan <- os.Interrupt
	}
}
