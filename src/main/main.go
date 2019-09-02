package main

import (
	"os"
	"sync"
)

var dupes int

var uniqueLines = struct {
	sync.RWMutex
	lineMap map[string][]string
}{lineMap: make(map[string][]string)}

var keywordCounter = struct {
	sync.RWMutex
	keywords map[string]int
}{keywords: make(map[string]int)}

func main() {
	keywordFileName := os.Args[1]
	fileDirectory := os.Args[2]
	buildKeywordMap(keywordFileName)
	readFilesFromDirectory(fileDirectory)
	writeOutputFile()
}
