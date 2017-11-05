package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
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

var asciiLowercase = "abcdefghijklmnopqrstuvwxyz"

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

type generic interface{}
type callable func(args ...generic) []generic

type TheOne struct {
	values []generic
}

func (theone *TheOne) Bind(function callable) {
	theone.values = function(theone.values...)
}

func (theone TheOne) Print() {
	for _, value := range theone.values {
		fmt.Printf("%v", value)
	}
}

func readFile(args ...generic) []generic {
	pathToFile := args[0].(string)
	stopWordsPath := args[1].(string)

	fd, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	data := string(fd)

	return []generic{data, stopWordsPath}
}

func filterChars(args ...generic) []generic {
	strData := args[0].(string)
	stopWordsPath := args[1].(string)

	pattern := regexp.MustCompile(`[\W_]+`)

	return []generic{pattern.ReplaceAllString(strData, ` `), stopWordsPath}
}

func normalize(args ...generic) []generic {
	strData := args[0].(string)
	stopWordsPath := args[1].(string)

	return []generic{strings.ToLower(strData), stopWordsPath}
}

func scan(args ...generic) []generic {
	strData := args[0].(string)
	stopWordsPath := args[1].(string)

	return []generic{strings.Fields(strData), stopWordsPath}
}

func getStopWords(args ...generic) []generic {
	wordList := args[0].([]string)
	stopWordsPath := args[1].(string)

	f, err := ioutil.ReadFile(stopWordsPath)
	if err != nil {
		panic(err)
	}
	stopWords := strings.Split(string(f), `,`)
	stopWords = append(stopWords, strings.Split(asciiLowercase, "")...)

	return []generic{wordList, stopWords}
}

func removeStopWords(args ...generic) []generic {
	wordList := args[0].([]string)
	stopWords := args[1].([]string)

	filteredWordList := []string{}
	for _, word := range wordList {
		if !stringInSlice(word, stopWords) {
			filteredWordList = append(filteredWordList, word)
		}
	}

	return []generic{filteredWordList}
}

func frequencies(args ...generic) []generic {
	wordList := args[0].([]string)

	wordFreqs := map[string]int{}
	for _, word := range wordList {
		if _, present := wordFreqs[word]; present {
			wordFreqs[word] += 1
		} else {
			wordFreqs[word] = 1
		}
	}

	return []generic{wordFreqs}
}

func buildFrequenciesPairs(args ...generic) []generic {
	wordFreqs := args[0].(map[string]int)

	pairList := make(PairList, len(wordFreqs))
	i := 0
	for word, frequency := range wordFreqs {
		pairList[i] = Pair{word, frequency}
		i++
	}

	return []generic{pairList}

}

func sortFrequencies(args ...generic) []generic {
	wordFreqs := args[0].(PairList)

	sort.Sort(sort.Reverse(wordFreqs))

	return []generic{wordFreqs}
}

func take25(args ...generic) []generic {
	wordFreqs := args[0].(PairList)

	if len(wordFreqs) < 25 {
		return []generic{wordFreqs}
	} else {
		return []generic{wordFreqs[0:25]}
	}
}

func printText(args ...generic) []generic {
	wordFreqs := args[0].(PairList)
	var buffer bytes.Buffer

	for _, pair := range wordFreqs {
		buffer.WriteString(fmt.Sprintf("%s  -  %d\n", pair.Key, pair.Value))
	}

	return []generic{buffer.String()}
}

func main() {
	theOne := TheOne{[]generic{os.Args[1], os.Args[2]}}
	theOne.Bind(readFile)
	theOne.Bind(filterChars)
	theOne.Bind(normalize)
	theOne.Bind(scan)
	theOne.Bind(getStopWords)
	theOne.Bind(removeStopWords)
	theOne.Bind(frequencies)
	theOne.Bind(buildFrequenciesPairs)
	theOne.Bind(sortFrequencies)
	theOne.Bind(take25)
	theOne.Bind(printText)
	theOne.Print()
}
