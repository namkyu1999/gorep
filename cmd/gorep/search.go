package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const cannotReadFile = "unable to read file: %v"
const errorWhileReadFile = "Error while reading file: %s"
const searchResultTemplate = "%s:%d|%s\n"
const countResultTemplate = "%s:%d\n"

// TODO: 하나의 파일 안에서도 청크 단위로 리팩토링
func search(filename, pattern string, resultsChannel chan *Result) {
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

func searchCount(filename, pattern string, printPrefix bool) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf(cannotReadFile, err)
	}

	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	count := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		count += strings.Count(line, pattern)
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf(errorWhileReadFile, err)
	}

	if count != 0 {
		if printPrefix {
			fmt.Printf(countResultTemplate, file.Name(), count)
		} else {
			fmt.Println(count)
		}
	}
}
