package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func ReadFilesFromDirectory() {
	var fileWaitGroup sync.WaitGroup
	files, err := ioutil.ReadDir("./files")
	check(err)
	for _, file := range files {
		if file.Name() != ".idea" {
			fileWaitGroup.Add(1)
			go readFile(&fileWaitGroup, file.Name())
		}
	}
	fileWaitGroup.Wait()
}

func readFile(fileWaitGroup *sync.WaitGroup, fileName string) {
	var lineWaitGroup sync.WaitGroup
	const BufferSize = 1048576
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	buffer := make([]byte, BufferSize)

	for {
		bytesRead, err := file.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		readableString := string(buffer[:bytesRead])
		reader := strings.NewReader(readableString)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			lineWaitGroup.Add(1)
			go processLine(&lineWaitGroup, scanner.Text())
		}
	}
	lineWaitGroup.Wait()
	fileWaitGroup.Done()
}

func processLine(lineWaitGroup *sync.WaitGroup, line string) {
	if isDuplicateLine(line) {
		dupes += 1
	} else {
		splitLine := saveOriginalLine(line)
		checkLineForKeywords(splitLine)
	}
	lineWaitGroup.Done()
}

func isDuplicateLine(line string) bool {
	uniqueLines.RLock()
	_, ok := uniqueLines.lineMap[line]
	uniqueLines.RUnlock()
	return ok
}

func saveOriginalLine(line string) []string {
	splitLine := splitLine(line)
	uniqueLines.Lock()
	uniqueLines.lineMap[line] = splitLine
	uniqueLines.Unlock()
	return splitLine
}

func splitLine(line string) []string {
	split := strings.Fields(line)
	return split
}

func checkLineForKeywords(splitLine []string) {
	var keywordWaitGroup sync.WaitGroup
	for _, word := range splitLine {
		keywordWaitGroup.Add(1)
		go checkIfKeyword(&keywordWaitGroup, word)
	}
	keywordWaitGroup.Wait()
}
