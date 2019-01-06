// DEN
// Copyright (C) 2018-2019 Andreas T Jonsson

//go:generate go run tools/version/version.go -file version/version.go

package main

import (
	"flag"
	"fmt"

	"gitlab.com/phix/den/client"
	"gitlab.com/phix/den/version"
)

var (
	printVersion,
	printAbout bool
)

func init() {
	flag.BoolVar(&printVersion, "version", false, "Show version")
	flag.BoolVar(&printAbout, "about", false, "Show information about the game")
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
	client.Start()
}
