package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"os"
	"sort"
)

func writeOutputFile() {
	file, err := os.Create("./output/output.txt")
	check(err)
	defer file.Close()

	writeDupes(file)

	writeStatisticalData(file)

	writeKeywordCount(file)
}

func writeDupes(file *os.File) {
	writeLine(file, "num_dupes", float64(dupes))
}

func writeStatisticalData(file *os.File) {
	writeMedianAndStandardDeviation("med_length", uniqueLineRuneLength, file)
	writeMedianAndStandardDeviation("med_tokens", uniqueLineTokenLength, file)
}

func writeMedianAndStandardDeviation(medLabel string, data []float64, file *os.File) {
	writeLine(file, medLabel, getMedian(data))
	writeLine(file, "std_length", getStandardDeviation(data))
}

func getMedian(list []float64) float64 {
	median, err := stats.Median(list)
	check(err)
	return median
}

func writeLine(file *os.File, key string, value float64) {
	data := fmt.Sprintf("%s\t%f\n", key, value)
	_, err := file.WriteString(data)
	check(err)
	fmt.Printf(data)
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
