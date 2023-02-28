package main

import (
	"fmt"
	"sort"
	"sync"
)

type Result struct {
	fileName string
	matches  []Match
}

type Match struct {
	lineNumber int
	line       string
}

func resultHandler(results chan *Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range results {
		sort.Slice(result.matches, func(i, j int) bool {
			return result.matches[i].lineNumber < result.matches[j].lineNumber
		})

		for _, match := range result.matches {
			fmt.Printf("%s:%d|%s\n", result.fileName, match.lineNumber+1, match.line)
		}
	}
}
