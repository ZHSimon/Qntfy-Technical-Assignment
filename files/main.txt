package main

import "sync"

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
	BuildKeywordMap()
	ReadFilesFromDirectory()
	WriteOutputFile()
}
