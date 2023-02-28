package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	config, err := setup()
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(2)
	}

	resultsChannel := make(chan *Result)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go resultHandler(resultsChannel, &wg)

	//TODO: 다 끝나면 resultsChannel 닫아 줘야 한다.
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
			search(fileName, config.pattern, resultsChannel)
		}
	}
	close(resultsChannel)
	wg.Wait()
}
