package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

type genericArgs []interface{}
type continuation func(args genericArgs, next continuation)

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

func readFile(args genericArgs, next continuation) {
	pathToFile := args[0].(string)
	stopWordsPath := args[1].(string)
	continuations := args[2].([]continuation)

	fd, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	data := string(fd)

	next(genericArgs{data, stopWordsPath, continuations[1:]}, continuations[0])
}

func filterChars(args genericArgs, next continuation) {
	strData := args[0].(string)
	stopWordsPath := args[1].(string)
	continuations := args[2].([]continuation)

	pattern := regexp.MustCompile(`[\W_]+`)

	next(genericArgs{pattern.ReplaceAllString(strData, ` `), stopWordsPath, continuations[1:]}, continuations[0])
}

func normalize(args genericArgs, next continuation) {
	strData := args[0].(string)
	stopWordsPath := args[1].(string)
	continuations := args[2].([]continuation)

	next(genericArgs{strings.ToLower(strData), stopWordsPath, continuations[1:]}, continuations[0])
}

func scan(args genericArgs, next continuation) {
	strData := args[0].(string)
	stopWordsPath := args[1].(string)
	continuations := args[2].([]continuation)

	next(genericArgs{strings.Fields(strData), stopWordsPath, continuations[1:]}, continuations[0])
}

func getStopWords(args genericArgs, next continuation) {
	wordList := args[0].([]string)
	stopWordsPath := args[1].(string)
	continuations := args[2].([]continuation)

	f, err := ioutil.ReadFile(stopWordsPath)
	if err != nil {
		panic(err)
	}
	stopWords := strings.Split(string(f), `,`)
	stopWords = append(stopWords, strings.Split(asciiLowercase, "")...)

	next(genericArgs{wordList, stopWords, continuations[1:]}, continuations[0])
}

func removeStopWords(args genericArgs, next continuation) {
	wordList := args[0].([]string)
	stopWords := args[1].([]string)
	continuations := args[2].([]continuation)

	filteredWordList := []string{}
	for _, word := range wordList {
		if !stringInSlice(word, stopWords) {
			filteredWordList = append(filteredWordList, word)
		}
	}

	next(genericArgs{filteredWordList, continuations[1:]}, continuations[0])
}

func frequencies(args genericArgs, next continuation) {
	wordList := args[0].([]string)
	continuations := args[1].([]continuation)

	wordFreqs := map[string]int{}
	for _, word := range wordList {
		if _, present := wordFreqs[word]; present {
			wordFreqs[word] += 1
		} else {
			wordFreqs[word] = 1
		}
	}

	next(genericArgs{wordFreqs, continuations[1:]}, continuations[0])
}

func buildFrequenciesPairs(args genericArgs, next continuation) {
	wordFreqs := args[0].(map[string]int)
	continuations := args[1].([]continuation)

	pairList := make(PairList, len(wordFreqs))
	i := 0
	for word, frequency := range wordFreqs {
		pairList[i] = Pair{word, frequency}
		i++
	}

	next(genericArgs{pairList, continuations[1:]}, continuations[0])

}

func sortFrequencies(args genericArgs, next continuation) {
	wordFreqs := args[0].(PairList)
	continuations := args[1].([]continuation)

	sort.Sort(sort.Reverse(wordFreqs))

	next(genericArgs{wordFreqs, continuations[1:]}, continuations[0])
}

func take25(args genericArgs, next continuation) {
	wordFreqs := args[0].(PairList)
	continuations := args[1].([]continuation)

	if len(wordFreqs) < 25 {
		next(genericArgs{wordFreqs, continuations[1:]}, continuations[0])
	} else {
		next(genericArgs{wordFreqs[0:25], continuations[1:]}, continuations[0])
	}
}

func printText(args genericArgs, next continuation) {
	wordFreqs := args[0].(PairList)
	continuations := args[1].([]continuation)

	for _, pair := range wordFreqs {
		fmt.Println(pair.Key, " - ", pair.Value)
	}

	next(genericArgs{continuations[1:]}, continuations[0])
}

func noOp(_ genericArgs, _ continuation) {}

func main() {
	continuations := []continuation{
		filterChars,
		normalize,
		scan,
		getStopWords,
		removeStopWords,
		frequencies,
		buildFrequenciesPairs,
		sortFrequencies,
		take25,
		printText,
		noOp,
		nil,  // required to be passed as the "next" continuation to noOP
	}
	readFile(genericArgs{os.Args[1], os.Args[2], continuations[1:]}, continuations[0])
}
