package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	sorting "sort" // to match the python code we have a variable called sort
	"strings"
)

// equivalent of Python's dictionary items
// from https://github.com/tvraman/go-learn/blob/master/pairlist/pairlist.go
type Pair struct {
	Key   string
	Value int
}

// this adds the ability to sort pairs
type PairList []Pair
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// equivalent to Python's string.ascii_lowercase
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

//Takes a path to a file and returns the entire
//contents of the file as a string
func readFile(pathToFile string) string {
	f, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	data := string(f)
	return data
}

//Takes a string and returns a copy with all nonalphanumeric
//chars replaced by white space
func filterCharsAndNormalize(strData string) string {
	pattern := regexp.MustCompile(`[\W_]+`)
	return strings.ToLower(pattern.ReplaceAllString(strData, ` `))
}

//Takes a string and scans for words, returning
//a list of words.
func scan(strData string) []string {
	return strings.Fields(strData)
}


//Takes a list of words and returns a copy with all stop
//words removed
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

//Takes a list of words and returns a dictionary associating
//words with frequencies of occurrence
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

//Takes a dictionary of words and their frequencies
//and returns a list of pairs where the entries are
//sorted by frequency
func sort(wordFreq map[string]int) PairList {
	pairList := make(PairList, len(wordFreq))
	i := 0
	for word, frequency := range wordFreq {
		pairList[i] = Pair{word, frequency}
		i++
	}
	sorting.Sort(sorting.Reverse(pairList))
	return pairList
}

//Takes a list of pairs where the entries are sorted by frequency and print them recursively.
func printAll(wordFreqs PairList, limit int) {
	if len(wordFreqs) > 0 && limit > 0 {
		fmt.Println(wordFreqs[0].Key, " - ", wordFreqs[0].Value)
		printAll(wordFreqs[1:], limit-1)
	}
}

// The main function
func main() {
	printAll(sort(frequencies(removeStopWords(scan(filterCharsAndNormalize(readFile(os.Args[1])))))), 25)
}
