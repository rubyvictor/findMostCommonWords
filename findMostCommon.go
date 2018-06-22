package main

import (
	"fmt"
	"regexp"
	"sort"
	s "strings"
)

func getWordsFrom(text string) []string {
	words := regexp.MustCompile("\\w+")
	return words.FindAllString(s.ToLower(text), -1)
}

func countWords(words []string) map[string]int {
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}
	return wordCounts
}

func consoleOut(orderedWordCounts PairList) {
	for _, wordCount := range orderedWordCounts {
		fmt.Printf("%v\n", wordCount)
	}
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	text := "To be, I am learning Go! go is a nice, nice, nice, nice, nice, language to learn so, so, so, so, so, so, so, as, as, as, as, as, to become best programmer. I got this from a book and it is my input to this test"
	consoleOut(rankByWordCount(countWords(getWordsFrom(text))))

}
