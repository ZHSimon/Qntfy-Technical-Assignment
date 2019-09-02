package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"os"
	"sort"
	"unicode/utf8"
)

func WriteOutputFile() {
	f, err := os.Create("output.txt")
	check(err)
	defer f.Close()

	writeDupes(f)

	writeStatisticalData(f)

	writeKeywordCount(f)
}

func writeDupes(file *os.File) {
	writeLine(file, "num_dupes", float64(dupes))
}

func writeStatisticalData(file *os.File) {
	keys, tokens := getCountOfKeysAndTokens()
	writeMedianAndStandardDeviation(keys, file)
	writeMedianAndStandardDeviation(tokens, file)
}

func getCountOfKeysAndTokens() (keys []float64, tokens []float64) {
	i := 0 //I read that this was up to 20% faster than using append() to achieve a cleaner-looking result
	uniqueLines.RLock()
	keys = make([]float64, len(uniqueLines.lineMap))
	tokens = make([]float64, len(uniqueLines.lineMap))
	for key, value := range uniqueLines.lineMap {
		keys[i] = float64(utf8.RuneCountInString(key))
		tokens[i] = float64(len(value))
		i++
	}
	uniqueLines.RUnlock()
	return
}

func writeMedianAndStandardDeviation(data []float64, file *os.File) {
	writeLine(file, "med_length", getMedian(data))
	writeLine(file, "std_length", getStandardDeviation(data))
}

func getMedian(list []float64) float64 {
	median, err := stats.Median(list)
	check(err)
	return median
}

func writeLine(f *os.File, key string, value float64) {
	data := fmt.Sprintf("%s\t%f\n", key, value)
	_, err := f.WriteString(data)
	check(err)
}

func getStandardDeviation(list []float64) float64 {
	standardDeviation, err := stats.StandardDeviation(list)
	check(err)
	return standardDeviation
}

func writeKeywordCount(file *os.File) {
	keywords := getSortedKeywords()
	for _, keyword := range keywords {
		keywordCounter.RLock()
		writeLine(file, keyword, float64(keywordCounter.keywords[keyword]))
		keywordCounter.RUnlock()
	}
}

func getSortedKeywords() []string {
	i := 0
	keywordCounter.RLock()
	keywords := make([]string, len(keywordCounter.keywords))
	for keyword := range keywordCounter.keywords {
		keywords[i] = keyword
		i++
	}
	keywordCounter.RUnlock()
	sort.Strings(keywords)
	return keywords
}
