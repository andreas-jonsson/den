// DEN
// Copyright (C) 2018 Andreas T Jonsson

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
	ver := flag.String("variable", "DEN_VERSION", "Environment variable containing the version number.")
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

var content = `// DEN
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
