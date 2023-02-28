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
const countResultTemplate = "%s:%d\n"

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

//func Process(f *os.File, start time.Time, end time.Time) error {
//	linesPool := sync.Pool{New: func() interface{} {
//		lines := make([]byte, 250*1024)
//		return lines
//	}}
//
//	stringPool := sync.Pool{New: func() interface{} {
//		lines := ""
//		return lines
//	}}
//
//	r := bufio.NewReader(f)
//
//	var wg sync.WaitGroup
//
//	for {
//		buf := linesPool.Get().([]byte)
//
//		n, err := r.Read(buf)
//		buf = buf[:n]
//
//		if n == 0 {
//			if err != nil {
//				fmt.Println(err)
//				break
//			}
//			if err == io.EOF {
//				break
//			}
//			return err
//		}
//
//		nextUntillNewline, err := r.ReadBytes('\n')
//
//		if err != io.EOF {
//			buf = append(buf, nextUntillNewline...)
//		}
//
//		wg.Add(1)
//		go func() {
//			ProcessChunk(buf, &linesPool, &stringPool, start, end)
//			wg.Done()
//		}()
//
//	}
//
//	wg.Wait()
//	return nil
//}
//
//func ProcessChunk(chunk []byte, linesPool *sync.Pool, stringPool *sync.Pool, start time.Time, end time.Time) {
//
//	var wg2 sync.WaitGroup
//
//	logs := stringPool.Get().(string)
//	logs = string(chunk)
//
//	linesPool.Put(chunk)
//
//	logsSlice := strings.Split(logs, "\n")
//
//	stringPool.Put(logs)
//
//	chunkSize := 300
//	n := len(logsSlice)
//	noOfThread := n / chunkSize
//
//	if n%chunkSize != 0 {
//		noOfThread++
//	}
//
//	for i := 0; i < (noOfThread); i++ {
//
//		wg2.Add(1)
//		go func(s int, e int) {
//			defer wg2.Done() //to avaoid deadlocks
//			for i := s; i < e; i++ {
//				text := logsSlice[i]
//				if len(text) == 0 {
//					continue
//				}
//				logSlice := strings.SplitN(text, ",", 2)
//				logCreationTimeString := logSlice[0]
//
//				logCreationTime, err := time.Parse("2006-01-02T15:04:05.0000Z", logCreationTimeString)
//				if err != nil {
//					fmt.Printf("\n Could not able to parse the time :%s for log : %v", logCreationTimeString, text)
//					return
//				}
//
//				if logCreationTime.After(start) && logCreationTime.Before(end) {
//					//fmt.Println(text)
//				}
//			}
//
//
//		}(i*chunkSize, int(math.Min(float64((i+1)*chunkSize), float64(len(logsSlice)))))
//	}
//
//	wg2.Wait()
//	logsSlice = nil
//}
