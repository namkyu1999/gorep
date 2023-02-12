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

func search(filename, pattern string, printPrefix bool) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf(cannotReadFile, err)
	}

	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	lineNumber := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if index := strings.Index(line, pattern); index > -1 {
			if printPrefix {
				fmt.Printf(searchResultTemplate, file.Name(), lineNumber, line)
			} else {
				fmt.Println(line)
			}
		}
		lineNumber++
	}
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
