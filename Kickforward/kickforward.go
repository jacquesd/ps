package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	sorted "sort"
	"strings"
)

// equivalent of Python's dictionary items
// from https://github.com/tvraman/go-learn/blob/master/pairlist/pairlist.go
type PairList []Pair
type Pair struct {
	Key   string
	Value int
}

// this adds the ability to sort pairs
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

var ascii_lowercase = "abcdefghijklmnopqrstuvwxyz"

// equivalent of Python's x in y
// from https://stackoverflow.com/a/15323988
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

type NoOpFn func()
type PrintFn func(PairList, NoOpFn)
type SortFn func(map[string]int, PrintFn)
type FrequenciesFn func(wordList []string, fn SortFn)
type RemoveStopWordsFn func(wordList []string, fn FrequenciesFn)
type ScanFn func(strData string, fn RemoveStopWordsFn)
type NormalizeFn func(strData string, fn ScanFn)
type FilterCharsFn func(strData string, fn NormalizeFn)


func readFile(pathToFile string, fn FilterCharsFn) {
	f, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}

	data := string(f)
	fn(data, normalize)
}

func filterChars(strData string, fn NormalizeFn) {
	pattern := regexp.MustCompile(`[\W_]+`)
	fn(pattern.ReplaceAllString(strData, ` `), scan)
}

func normalize(strData string, fn ScanFn) {
	fn(strings.ToLower(strData), removeStopWords)
}

func scan(strData string, fn RemoveStopWordsFn) {
	fn(strings.Fields(strData), frequencies)
}

func removeStopWords(wordList []string, fn FrequenciesFn) {
	f, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		panic(err)
	}
	stopWords := strings.Split(string(f), `,`)
	stopWords = append(stopWords, strings.Split(ascii_lowercase, "")...)

	filteredWordList := []string{}
	for _, word := range wordList {
		if !stringInSlice(word, stopWords) {
			filteredWordList = append(filteredWordList, word)
		}
	}
	fn(filteredWordList, sort)
}

func frequencies(wordList []string, fn SortFn) {
	wordFreqs := map[string]int{}
	for _, word := range wordList {
		if _, present := wordFreqs[word]; present {
			wordFreqs[word] += 1
		} else {
			wordFreqs[word] = 1
		}
	}
	fn(wordFreqs, printText)
}

func sort(wf map[string]int, fn PrintFn) {
	pairList := make(PairList, len(wf))
	i := 0
	for word, frequency := range wf {
		pairList[i] = Pair{word, frequency}
		i++
	}
	sorted.Sort(sorted.Reverse(pairList))
	fn(pairList, noOp)
}

func printText(wordFreqs PairList, fn NoOpFn) {
	for i, pair := range wordFreqs {
		if i >= 25 {
			return
		}
		fmt.Println(pair.Key, " - ", pair.Value)
	}
	fn()
}

func noOp() {}

func main() {
	readFile(os.Args[1], filterChars)
}
