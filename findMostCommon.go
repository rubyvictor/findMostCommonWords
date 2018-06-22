package main

import (
    "fmt"
    "regexp"
    s "strings"
    "sort"
)

func get_words_from(text string) []string{
    words:= regexp.MustCompile("\\w+")
    return words.FindAllString(s.ToLower(text), -1)
}

func count_words (words []string) map[string]int{
    word_counts := make(map[string]int)
    for _, word :=range words{
        word_counts[word]++
    }
    return word_counts;
}

func console_out (ordered_word_counts PairList){
    for _, word_count :=range ordered_word_counts{
        fmt.Printf("%v\n", word_count)
    }
}


func rankByWordCount(wordFrequencies map[string]int) PairList{
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
  Key string
  Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }


func main() {
    text := "To be, I am learning Go! go is a nice, nice, nice, nice, nice, language to learn so, so, so, so, so, so, so, as, as, as, as, as, to become best programmer. I got this from a book and it is my input to this test"
console_out(rankByWordCount(count_words(get_words_from(text))))  

}