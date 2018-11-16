// DEN
// Copyright (C) 2018 Andreas T Jonsson

package main

import (
	"flag"
)

var startServer bool

func init() {
	flag.BoolVar(&startServer, "server", false, "Start server instance")
}

func main() {
	flag.Parse()
	if startServer {
		server.Start()
	}else {
		client.Start()
	}
}
