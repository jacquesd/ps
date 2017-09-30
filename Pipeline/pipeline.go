package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	sorted "sort"
	"strings"
)

type PairList []Pair
type Pair struct {
	Key   string
	Value int
}

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

var ascii_lowercase = "abcdefghijklmnopqrstuvwxyz"

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func readFile(pathToFile string) string {
	f, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	data := string(f)
	return data
}

func filterCharsAndNormalize(strData string) string {
	pattern := regexp.MustCompile(`[\W_]+`)
	return strings.ToLower(pattern.ReplaceAllString(strData, ` `))
}

func scan(strData string) []string {
	return strings.Fields(strData)
}

func removeStopWords(wordList []string) []string {
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
	return filteredWordList
}

func frequencies(wordList []string) map[string]int {
	wordFreqs := map[string]int{}
	for _, word := range wordList {
		if _, present := wordFreqs[word]; present {
			wordFreqs[word] += 1
		} else {
			wordFreqs[word] = 1
		}
	}
	return wordFreqs
}

func sort(wordFrequ map[string]int) PairList {
	pairList := make(PairList, len(wordFrequ))
	i := 0
	for word, frequency := range wordFrequ {
		pairList[i] = Pair{word, frequency}
		i++
	}
	sorted.Sort(sorted.Reverse(pairList))
	return pairList
}

func printAll(wordFreqs PairList, limit int) {
	if len(wordFreqs) > 0 && limit > 0 {
		fmt.Println(wordFreqs[0].Key, " - ", wordFreqs[0].Value)
		printAll(wordFreqs[1:], limit-1)
	}
}

func main() {
	printAll(sort(frequencies(removeStopWords(scan(filterCharsAndNormalize(readFile(os.Args[1])))))), 25)
}
