package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
)

const cannotReadFile = "unable to read file: %v"
const errorWhileReadFile = "Error while reading file: %s"

func search(filename, pattern string, resultsChannel chan *Result, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf(cannotReadFile, err)
	}

	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	lineNumber := 0
	result := &Result{
		fileName: file.Name(),
		matches:  make([]Match, 0),
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if index := strings.Index(line, pattern); index > -1 {
			result.matches = append(result.matches, Match{lineNumber: lineNumber, line: line})
		}
		lineNumber++
	}

	resultsChannel <- result

	if err := fileScanner.Err(); err != nil {
		log.Fatalf(errorWhileReadFile, err)
	}
}
