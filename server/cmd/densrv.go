// DEN
// Copyright (C) 2018 Andreas T Jonsson

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"gitlab.com/phix/den/server"
	"gitlab.com/phix/den/version"
)

var printVersion bool

func init() {
	flag.BoolVar(&printVersion, "version", false, "Show version")
}

func Start() {
	flag.Parse()
	if printVersion {
		fmt.Println(version.String)
		return
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	server.InterruptChan = signalChan

	server.Start()
}
