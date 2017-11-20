package main

import (
	"io/ioutil"
	"regexp"
	"strings"
	"sort"
	"os"
	"fmt"
)

type generic interface{}
type closedMap map[string]generic

var dataStorageObj closedMap
var stopWordsObj closedMap
var wordFreqsObj closedMap

func init() {
	dataStorageObj = closedMap{
		"data": []string{},
		"init": func(pathToFile string) {
			fd, err := ioutil.ReadFile(pathToFile)
			if err != nil { panic(err) }

			pattern := regexp.MustCompile(`[\W_]+`)
			data := strings.ToLower(pattern.ReplaceAllString(string(fd), ` `))
			dataStorageObj["data"] = strings.Fields(data)
		},
		"words": func() []string { return dataStorageObj["data"].([]string) },
	}


	stopWordsObj = closedMap{
		"stopWords": []string{},
		"init": func(stopWordsPath string) {
			f, err := ioutil.ReadFile(stopWordsPath)
			if err != nil { panic(err) }

			stopWords := strings.Split(string(f), `,`)
			stopWords = append(stopWords, strings.Split(asciiLowercase, "")...)
			stopWordsObj["stopWords"] = stopWords
		},
		"isStopWord": func(word string) bool { return stringInSlice(word, stopWordsObj["stopWords"].([]string)) },
	}

	wordFreqsObj = closedMap{
		"freqs": map[string]int{},
		"incrementCount": func(word string) { wordFreqsObj["freqs"].(map[string]int)[word] += 1 },
		"sorted": func() PairList {
			wordFreqs := wordFreqsObj["freqs"].(map[string]int)
			pairList := make(PairList, len(wordFreqs))
			i := 0
			for word, frequency := range wordFreqs {
				pairList[i] = Pair{word, frequency}
				i++
			}
			sort.Sort(sort.Reverse(pairList))
			return pairList
		},
		"compute": func(getWords func() []string, isStopWord func(string) bool) {
			for _, word := range getWords() {
				if !isStopWord(word) {
					wordFreqsObj["incrementCount"].(func(string))(word)
				}
			}
		},
		"print": func() {
			wordFreqs := wordFreqsObj["sorted"].(func() PairList)()
			if len(wordFreqs) > 25 { wordFreqs = wordFreqs[0:25] }

			for _, pair := range wordFreqs { fmt.Printf("%s  -  %d\n", pair.Key, pair.Value) }
		},
	}
}

func main() {
	dataStorageObj["init"].(func(string))(os.Args[1])
	stopWordsObj["init"].(func(string))(os.Args[2])
	wordFreqsObj["compute"].(func(func() []string, func(string) bool))(
		dataStorageObj["words"].(func() []string),
		stopWordsObj["isStopWord"].(func(string) bool),
	)
	wordFreqsObj["print"].(func())()
}