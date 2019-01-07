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
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
	"time"
)

func main() {
	file := flag.String("file", "-", "Save the generated output to file.")
	pkg := flag.String("package", "version", "Package name of the generated output.")
	ver := flag.String("variable", "FULL_VERSION", "Environment variable containing the version number.")
	flag.Parse()

	cmd := exec.Command("git", "rev-parse", "HEAD")
	res, err := cmd.Output()
	if err != nil {
		log.Panicln(err)
	}

	version := os.Getenv(*ver)
	if version == "" {
		version = "0.0.0.0"
		log.Printf("%s is not set. Defaulting to 0.0.0-0\n", *ver)
	}

	parts := strings.SplitN(version, ".", 4)
	if len(parts) != 4 {
		log.Panicf("invalid version format: %s\n", version)
	}

	const (
		startYear    = 2018
		copyrightFmt = "Copyright (C) %v Andreas T Jonsson"
	)

	copyrightString := fmt.Sprintf(copyrightFmt, startYear)
	if year := time.Now().Year(); year != startYear {
		copyrightString = fmt.Sprintf(copyrightFmt, fmt.Sprintf("%d-%d", startYear, year))
	}

	values := map[string]interface{}{
		"hash":  strings.TrimSpace(string(res)),
		"major": parts[0],
		"minor": parts[1],
		"patch": parts[2],
		"build": parts[3],
		"copy":  copyrightString,
		"pkg":   *pkg,
	}

	tmpl := template.New("version")
	tmpl = template.Must(tmpl.Parse(content))
	os.MkdirAll(path.Dir(*file), 0777)

	fp := os.Stdout
	if *file != "-" {
		fp, err = os.Create(*file)
		if err != nil {
			log.Panicln(err)
		}
		defer fp.Close()
	}

	if err := tmpl.Execute(fp, values); err != nil {
		log.Panicln(err)
	}
}

var content = `// DEN is released under GPL3
// {{.copy}}

package version

const (
	Major = {{.major}}
	Minor = {{.minor}}
	Patch = {{.patch}}
)

var (
	String = "{{.major}}.{{.minor}}.{{.patch}}"
	Full = "{{.major}}.{{.minor}}.{{.patch}}-{{.build}} ({{.hash}})"
)

var (
	Build = {{.build}}
	Hash = "{{.hash}}"
	Copyright = "{{.copy}}"
)
`
