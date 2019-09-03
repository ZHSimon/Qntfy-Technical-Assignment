package main

import (
	"bufio"
	"log"
	"os"
)

func buildKeywordMap(keywordFile string) {
	file, err := os.Open(keywordFile)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newKeyword := scanner.Text()
		addKeyword(newKeyword)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func addKeyword(keyword string) {
	keywordCounter.Lock()
	keywordCounter.keywords[keyword] = 0
	keywordCounter.Unlock()
}

func isKeyword(word string) bool {
	keywordCounter.RLock()
	_, ok := keywordCounter.keywords[word]
	keywordCounter.RUnlock()
	return ok
}

func incrementKeyword(word string) {
	keywordCounter.Lock()
	keywordCounter.keywords[word]++
	keywordCounter.Unlock()
}

func getKeywordsInLine(splitLine []string) (keywordsInLine []string) {
	for _, word := range splitLine {
		if isKeyword(word) {
			keywordsInLine = append(keywordsInLine, word)
		}
	}
	return keywordsInLine
}

func saveKeywordsInLine(line string, keywordsInLine []string) {
	uniqueLines.Lock()
	uniqueLines.lineMap[line] = keywordsInLine
	uniqueLines.Unlock()
}

func incrementKeywords(keywords []string) {
	for _, word := range keywords {
		incrementKeyword(word)
	}
}
