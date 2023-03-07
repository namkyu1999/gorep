package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	config, err := setup()
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(2)
	}

	resultsChannel := make(chan *Result)
	resultHandler := NewResultHandler(config.count)

	resultWaitGroup := sync.WaitGroup{}
	resultWaitGroup.Add(1)

	go resultHandler.handle(resultsChannel, &resultWaitGroup)

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

	fileWaitGroup := sync.WaitGroup{}
	for _, fileName := range config.fileName {
		fileWaitGroup.Add(1)
		go search(fileName, config.pattern, resultsChannel, &fileWaitGroup)
	}
	fileWaitGroup.Wait()

	close(resultsChannel)

	resultWaitGroup.Wait()
	fmt.Println("\nTime taken - ", time.Since(now))
}
