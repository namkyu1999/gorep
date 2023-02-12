package main

import (
	"errors"
	"flag"
	"io"
	"os"
)

type configuration struct {
	dereferenceRecursive bool
	count                bool
	pattern              string
	fileName             []string
	fromPipe             bool
}

const invalidArgument = "invalid argument count"
const cannotCreatTemp = "cannot create temp file"
const cannotReadFromPipe = "cannot read from pipe"
const cannotWriteTemp = "cannot write to temp file"

func setup() (configuration, error) {
	var config configuration
	var dereferenceRecursive, count bool

	flag.BoolVar(&dereferenceRecursive, "R", false, "dereference recursive")
	flag.BoolVar(&count, "c", false, "count")
	flag.Parse()
	tail := flag.Args()

	if fi, _ := os.Stdin.Stat(); !dereferenceRecursive && (fi.Mode()&os.ModeCharDevice) == 0 {
		tmpFile, err := os.CreateTemp("", "gorep")
		if err != nil {
			return config, errors.New(cannotCreatTemp)
		}

		line, err := io.ReadAll(os.Stdin)
		if err != nil {
			return config, errors.New(cannotReadFromPipe)
		}

		err = os.WriteFile(tmpFile.Name(), line, 0644)
		if err != nil {
			return config, errors.New(cannotWriteTemp)
		}
		defer tmpFile.Close()

		return configuration{
			dereferenceRecursive: dereferenceRecursive,
			count:                count,
			pattern:              tail[0],
			fileName:             []string{tmpFile.Name()},
			fromPipe:             true,
		}, nil
	} else if len(tail) < 2 {
		return config, errors.New(invalidArgument)
	} else {
		return configuration{
			dereferenceRecursive: dereferenceRecursive,
			count:                count,
			pattern:              tail[0],
			fileName:             tail[1:],
			fromPipe:             false,
		}, nil
	}
}
