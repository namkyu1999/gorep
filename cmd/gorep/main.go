package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	config, err := setup()
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(2)
	}

	if config.dereferenceRecursive {
		paths := make([]string, 0)
		err := filepath.Walk(config.fileName[0],
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					paths = append(paths, path)
				}
				return nil
			})
		if err != nil {
			fmt.Errorf(err.Error())
			os.Exit(2)
		}
		config.fileName = paths
	}

	if printPrefix := len(config.fileName) > 1; config.count {
		for _, fileName := range config.fileName {
			searchCount(fileName, config.pattern, printPrefix)
		}
	} else {
		for _, fileName := range config.fileName {
			search(fileName, config.pattern, printPrefix)
		}
	}
}
