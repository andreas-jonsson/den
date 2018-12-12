// DEN
// Copyright (C) 2018 Andreas T Jonsson

package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var host string

func init() {
	flag.StringVar(&host, "host", "localhost:5001", "Connect to game instance")
}

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", host)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to", host)
	io.Copy(os.Stdout, conn)
	fmt.Println("Connection closed!")
}
