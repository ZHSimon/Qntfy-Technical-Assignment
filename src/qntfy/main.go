package qntfy

import (
	"os"
	"sync"
)

var dupes int
var countDuplicateLines = false

var uniqueLines = struct {
	sync.RWMutex
	lineMap map[string][]string
}{lineMap: make(map[string][]string)}

var uniqueLineRuneLength []float64
var uniqueLineTokenLength []float64

var keywordCounter = struct {
	sync.RWMutex
	keywords map[string]int
}{keywords: make(map[string]int)}

func main() {
	keywordFileName, fileDirectory := getArgs()
	buildKeywordMap(keywordFileName)
	readFilesFromDirectory(fileDirectory)
	writeOutputFile()
}

func getArgs() (string, string) {
	keywordFileName := "./keywords.txt"
	fileDirectory := "./files/"
	if len(os.Args) > 1 {
		keywordFileName = os.Args[1]
	}
	if len(os.Args) > 2 {
		fileDirectory = os.Args[2]
	}
	return keywordFileName, fileDirectory
}
