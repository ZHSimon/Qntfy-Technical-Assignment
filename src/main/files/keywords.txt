package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func BuildKeywordMap() {
	file, err := os.Open("./keywords.txt")
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


func checkIfKeyword(keywordWaitGroup *sync.WaitGroup, word string) {
	if isKeyword(word) {
		incrementKeyword(word)
	}
	keywordWaitGroup.Done()
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