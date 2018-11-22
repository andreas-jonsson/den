// DEN
// Copyright (C) 2018 Andreas T Jonsson

//go:generate go run tools/version/version.go -file version/version.go

package main

import (
	"flag"

	"gitlab.com/phix/den/client"
	"gitlab.com/phix/den/server"
)

var startServer bool

func init() {
	flag.BoolVar(&startServer, "server", false, "Start server instance")
}

func main() {
	flag.Parse()
	if startServer {
		server.Start()
	} else {
		client.Start()
	}
}
