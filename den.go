/*
Copyright (C) 2018-2019 Andreas T Jonsson

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
