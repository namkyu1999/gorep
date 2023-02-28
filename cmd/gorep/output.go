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

type ResultHandler interface {
	handle(results chan *Result, wg *sync.WaitGroup)
}

type NormalResultHandler struct{}

func (n NormalResultHandler) handle(results chan *Result, wg *sync.WaitGroup) {
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

type CountResultHandler struct{}

func (c CountResultHandler) handle(results chan *Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range results {
		count := len(result.matches)
		fmt.Printf("%s|%d\n", result.fileName, count)
	}
}

func NewResultHandler(isCount bool) ResultHandler {
	if isCount {
		return CountResultHandler{}
	} else {
		return NormalResultHandler{}
	}
}
