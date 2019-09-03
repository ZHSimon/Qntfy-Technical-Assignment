package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

func readFilesFromDirectory(fileDirectory string) {
	var fileWaitGroup sync.WaitGroup
	files, err := ioutil.ReadDir(fileDirectory)
	check(err)
	for _, file := range files {
		if !file.IsDir() {
			fileWaitGroup.Add(1)
			go readFile(&fileWaitGroup, fileDirectory+file.Name())
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

	splitLine := splitLine(line)
	if keywordsInLine, ok := isDuplicateLine(line); ok {
		handleDuplicateLine(line, splitLine, keywordsInLine)
	} else {
		handleUniqueLine(line, splitLine)
	}
	lineWaitGroup.Done()
}

func splitLine(line string) []string {
	split := strings.Fields(line)
	return split
}

func isDuplicateLine(line string) ([]string, bool) {
	uniqueLines.RLock()
	value, ok := uniqueLines.lineMap[line]
	uniqueLines.RUnlock()
	return value, ok
}

func handleDuplicateLine(line string, splitLine []string, keywordsInLine []string) {
	if countDuplicateLines {
		recordLineStats(line, splitLine, keywordsInLine)
	}
	dupes += 1
}

func recordLineStats(line string, splitLine []string, keywordsInLine []string) {
	getLineStatistics(line, splitLine)
	incrementKeywords(keywordsInLine)
}

func getLineStatistics(line string, split []string) {
	uniqueLineRuneLength = append(uniqueLineRuneLength, float64(utf8.RuneCountInString(line)))
	uniqueLineTokenLength = append(uniqueLineTokenLength, float64(len(split)))
}

func handleUniqueLine(line string, splitLine []string) {
	keywordsInLine := getKeywordsInLine(splitLine)
	saveKeywordsInLine(line, keywordsInLine)
	recordLineStats(line, splitLine, keywordsInLine)
}
