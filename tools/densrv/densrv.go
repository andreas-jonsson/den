// DEN
// Copyright (C) 2018 Andreas T Jonsson

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"gitlab.com/phix/den/client"
	"gitlab.com/phix/den/server"
	"gitlab.com/phix/den/version"
)

const includeClient = true

var (
	printVersion,
	hostLocal bool
)

func init() {
	flag.BoolVar(&printVersion, "version", false, "Show version")

	if includeClient {
		flag.BoolVar(&hostLocal, "local", false, "Host local game")
	}
}

func main() {
	flag.Parse()
	if printVersion {
		fmt.Println(version.String)
		return
	}

	signalChan := make(chan os.Signal, 1)
	server.InterruptChan = signalChan

	if hostLocal {
		go func() {
			<-client.GameExitedChan
			signalChan <- os.Interrupt
		}()
		go client.Start()
		<-client.LoggerInitializedChan
	} else {
		signal.Notify(signalChan, os.Interrupt)
	}
	server.Start()
}
