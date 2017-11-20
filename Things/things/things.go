package things

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)


type Thing interface { Info() string }

type dataStorageManager struct { data string }

func (dsm dataStorageManager) Words() []string { return strings.Fields(dsm.data) }

func NewDataStorageManager(pathToFile string) *dataStorageManager {
	fd, err := ioutil.ReadFile(pathToFile)
	if err != nil { panic(err) }

	pattern := regexp.MustCompile(`[\W_]+`)
	data := strings.ToLower(pattern.ReplaceAllString(string(fd), ` `))
	return &dataStorageManager{data}
}

type stopWordManager struct { stopWords []string }

func (swm stopWordManager) IsStopWord(word string) bool { return stringInSlice(word, swm.stopWords) }

func NewStopWordManager(stopWordsPath string) *stopWordManager {
	f, err := ioutil.ReadFile(stopWordsPath)
	if err != nil { panic(err) }

	stopWords := strings.Split(string(f), `,`)
	stopWords = append(stopWords, strings.Split(asciiLowercase, "")...)
	return &stopWordManager{stopWords}
}

type wordsFrequencyManager struct { wordFreqs map[string]int }


func (wfm wordsFrequencyManager) IncreaseCount(word string) { wfm.wordFreqs[word] += 1 }

func (wfm wordsFrequencyManager) Sorted() PairList {
	pairList := make(PairList, len(wfm.wordFreqs))
	i := 0
	for word, frequency := range wfm.wordFreqs {
		pairList[i] = Pair{word, frequency}
		i++
	}
	sort.Sort(sort.Reverse(pairList))
	return pairList
}

func NewWordsFrequencyManager() *wordsFrequencyManager { return &wordsFrequencyManager{map[string]int{}} }

type wordFrequencyController struct {
	storageManager *dataStorageManager
	stopWordsManager *stopWordManager
	wordFreqManager *wordsFrequencyManager
}

func (wfc wordFrequencyController) Run() {
	for _, word := range wfc.storageManager.Words() {
		if !wfc.stopWordsManager.IsStopWord(word) {
			wfc.wordFreqManager.IncreaseCount(word)
		}
	}

	wordFreqs := wfc.wordFreqManager.Sorted()
	if len(wordFreqs) > 25 { wordFreqs = wordFreqs[0:25] }
	for _, pair := range wordFreqs { fmt.Printf("%s  -  %d\n", pair.Key, pair.Value) }
}

func NewWordFrequencyController(pathToFile, pathToStopWords string) *wordFrequencyController {
	return &wordFrequencyController{
		NewDataStorageManager(pathToFile),
		NewStopWordManager(pathToStopWords),
		NewWordsFrequencyManager(),
	}
}