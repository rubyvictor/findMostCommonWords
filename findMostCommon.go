package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

func filterPairs(allPairs PairList) PairList {
	return allPairs[0:10]
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(pl)
	return filterPairs(pl)
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func displayPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprintf(w, "Hello & Welcome To My Golang Server to find top ten words")
}

func text(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		render(w, "form.gtpl", nil)
	} else {
		r.ParseForm()
		msg := &Message{
			Content: r.FormValue("text"),
		}
		if msg.Validate() == false {
			render(w, "form.gtpl", msg)
		} else {
			for key, value := range r.Form {
				fmt.Println("key", key)
				fmt.Println("val:", s.Join(value, ""))
				fmt.Fprintf(w, "Top 10 Words:%v", rankByWordCount(countWords(getWordsFrom(msg.Content))))
			}
		}
	}
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", displayPage)
	http.HandleFunc("/text", text)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
