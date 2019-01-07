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

	if includeClient && hostLocal {
		flag.Set("host", "localhost:5000")
		flag.Parse()
	}

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
