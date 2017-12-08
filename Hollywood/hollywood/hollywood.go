package hollywood

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)


type Handler func()
type StringHandler func(string)
type LoadHandler func(string, string)
type StopWordFilter interface { IsStopWord(string) bool }

type dataStorageManager struct {
	data string
	stopWordFilter StopWordFilter
	wordEventHandlers []StringHandler
}

func (self *dataStorageManager) load(pathToFile, _ string) {
	fd, err := ioutil.ReadFile(pathToFile)
	if err != nil { panic(err) }

	pattern := regexp.MustCompile(`[\W_]+`)
	self.data = strings.ToLower(pattern.ReplaceAllString(string(fd), ` `))
}

func (self *dataStorageManager) produceWords() {
	for _, word := range strings.Fields(self.data) {
		if !self.stopWordFilter.IsStopWord(word) {
			for _, handler := range self.wordEventHandlers { handler(word) }
		}
	}
}

func (self *dataStorageManager) registerForWordEvent(handler StringHandler) {
	self.wordEventHandlers = append(self.wordEventHandlers, handler)
}

func NewDataStorageManager(wfapp *wordFrequencyFramework, stopWordFilter StopWordFilter) *dataStorageManager {
	self := dataStorageManager{stopWordFilter: stopWordFilter}
	wfapp.registerForLoadEvent(self.load)
	wfapp.registerForDoworkEvent(self.produceWords)
	return &self
}

type stopWordManager struct { stopWords []string }

func (self *stopWordManager) load(_, stopWordsPath string) {
	f, err := ioutil.ReadFile(stopWordsPath)
	if err != nil { panic(err) }

	self.stopWords = strings.Split(string(f), `,`)
	self.stopWords = append(self.stopWords, strings.Split(asciiLowercase, "")...)
}

func (self stopWordManager) IsStopWord(word string) bool { return stringInSlice(word, self.stopWords) }

func NewStopWordManager(wfapp *wordFrequencyFramework) *stopWordManager {
	self := stopWordManager{}
	wfapp.registerForLoadEvent(self.load)
	return &self
}

type wordsFrequencyCounter struct { wordFreqs map[string]int }

func (self wordsFrequencyCounter) increaseCount(word string) { self.wordFreqs[word] += 1 }

func (self wordsFrequencyCounter) printFreqs() {
	pairList := make(PairList, len(self.wordFreqs))
	i := 0
	for word, frequency := range self.wordFreqs {
		pairList[i] = Pair{word, frequency}
		i++
	}
	sort.Sort(sort.Reverse(pairList))
	if len(pairList) > 25 { pairList = pairList[0:25] }
	for _, pair := range pairList { fmt.Printf("%s  -  %d\n", pair.Key, pair.Value) }
}

func NewWordsFrequencyCounter(wfapp *wordFrequencyFramework, dataStorage *dataStorageManager) *wordsFrequencyCounter {
	self := wordsFrequencyCounter{wordFreqs: map[string]int{}}
	dataStorage.registerForWordEvent(self.increaseCount)
	wfapp.registerForEndEvent(self.printFreqs)
	return &self
}

type wordFrequencyFramework struct {
	loadEventHandlers []LoadHandler
	doworkEventHandlers []Handler
	endEventHandlers []Handler
}

func (self *wordFrequencyFramework) registerForLoadEvent(handler LoadHandler) {
	self.loadEventHandlers = append(self.loadEventHandlers, handler)
}

func (self *wordFrequencyFramework) registerForDoworkEvent(handler Handler) {
	self.doworkEventHandlers = append(self.doworkEventHandlers, handler)
}

func (self *wordFrequencyFramework) registerForEndEvent(handler Handler) {
	self.endEventHandlers = append(self.endEventHandlers, handler)
}

func (self wordFrequencyFramework) Run(pathToFile, stopWordsPath string) {
	for _, handler := range self.loadEventHandlers { handler(pathToFile, stopWordsPath) }
	for _, handler := range self.doworkEventHandlers { handler() }
	for _, handler := range self.endEventHandlers { handler() }
}

func NewWordFrequencyFrameWork() *wordFrequencyFramework { return &wordFrequencyFramework{} }