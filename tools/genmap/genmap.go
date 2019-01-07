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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/phix/den/client/world"
	"gitlab.com/phix/den/logger"
	"gitlab.com/phix/den/message"
)

var pkg string

func init() {
	flag.StringVar(&pkg, "pkg", "./level", "Location of the level package")
}

func main() {
	flag.Parse()

	files, err := filepath.Glob(filepath.Join(pkg, "*.ascii"))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println("Generating:", file)

		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		var width int
		level := [][]byte{[]byte{}}

		var x, y int
		for _, b := range data {
			if b == '\n' {
				if x > width {
					width = x
				}

				level = append(level, []byte{})
				x = 0
				y++
				continue
			}

			t := runeToTile(rune(b))
			level[y] = append(level[y], t)
			x++
		}

		if len(level) != width {
			panic("level is not square")
		}

		var rawLevel []byte
		for y, line := range level {
			if ln := len(line); ln < width {
				line = append(level[y], []byte(strings.Repeat(" ", width-ln))...)
				level[y] = line
			}
			rawLevel = append(rawLevel, line...)
		}

		name := strings.TrimSuffix(file, ".ascii")
		fp, err := os.Create(name + ".go")
		if err != nil {
			panic(err)
		}

		varName := "L" + filepath.Base(name)[1:]

		fmt.Fprint(fp, "package level\n\n")
		fmt.Fprintf(fp, "var %s = %#v\n", varName, rawLevel)
		fp.Close()

		for _, line := range level {
			for _, t := range line {
				fmt.Printf("%c", world.TileToRune(t))
			}
			fmt.Println("")
		}
	}
}

func runeToTile(r rune) byte {
	switch r {
	case ' ':
		return message.EmptyTile
	case '#':
		return message.WallTile
	case '.':
		return message.FloorTile
	case '|':
		return message.VDoorTile
	case '=':
		return message.HDoorTile
	default:
		logger.Fatalln("Invalid tile:", r)
		return 0
	}
}
