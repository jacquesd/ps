package letterbox

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

type generic interface{}
type LetterBox interface { Dispatch(operation string, args ...string) []generic }

type dataStorageManager struct { data string }

func (dsm dataStorageManager) words() []string { return strings.Fields(dsm.data) }

func (dsm dataStorageManager) Dispatch(operation string, args ...string) []generic {
	if operation == "words" { return []generic{dsm.words()} }
	panic(fmt.Sprintf("Unexpected operation '%s' for 'dataStorageManager'", operation))
}

func NewDataStorageManager(pathToFile string) *dataStorageManager {
	fd, err := ioutil.ReadFile(pathToFile)
	if err != nil { panic(err) }

	pattern := regexp.MustCompile(`[\W_]+`)
	data := strings.ToLower(pattern.ReplaceAllString(string(fd), ` `))
	return &dataStorageManager{data}
}

type stopWordManager struct { stopWords []string }

func (swm stopWordManager) isStopWord(word string) bool { return stringInSlice(word, swm.stopWords) }

func (swm stopWordManager) Dispatch(operation string, args ...string) []generic {
	if operation == "isStopWord" { return []generic{swm.isStopWord(args[0])} }
	panic(fmt.Sprintf("Unexpected operation '%s' for 'stopWordManager'", operation))
}

func NewStopWordManager(stopWordsPath string) *stopWordManager {
	f, err := ioutil.ReadFile(stopWordsPath)
	if err != nil { panic(err) }

	stopWords := strings.Split(string(f), `,`)
	stopWords = append(stopWords, strings.Split(asciiLowercase, "")...)
	return &stopWordManager{stopWords}
}

type wordsFrequencyManager struct { wordFreqs map[string]int }

func (wfm wordsFrequencyManager) increaseCount(word string) { wfm.wordFreqs[word] += 1 }

func (wfm wordsFrequencyManager) sorted() PairList {
	pairList := make(PairList, len(wfm.wordFreqs))
	i := 0
	for word, frequency := range wfm.wordFreqs {
		pairList[i] = Pair{word, frequency}
		i++
	}
	sort.Sort(sort.Reverse(pairList))
	return pairList
}

func (wfm wordsFrequencyManager) Dispatch(operation string, args ...string) []generic {
	if operation == "increaseCount" {
		wfm.increaseCount(args[0])
		return []generic{}
	} else if operation == "sorted" {
		return []generic{wfm.sorted()}
	}
	panic(fmt.Sprintf("Unexpected operation '%s' for 'wordsFrequencyManager'", operation))
}

func NewWordsFrequencyManager() *wordsFrequencyManager { return &wordsFrequencyManager{map[string]int{}} }

type wordFrequencyController struct {
	storageManager   LetterBox
	stopWordsManager LetterBox
	wordFreqManager  LetterBox
}

func (wfc wordFrequencyController) run() {
	for _, word := range wfc.storageManager.Dispatch("words")[0].([]string) {
		if !wfc.stopWordsManager.Dispatch("isStopWord", word)[0].(bool) {
			wfc.wordFreqManager.Dispatch("increaseCount", word)
		}
	}

	wordFreqs := wfc.wordFreqManager.Dispatch("sorted")[0].(PairList)
	if len(wordFreqs) > 25 { wordFreqs = wordFreqs[0:25] }

	for _, pair := range wordFreqs { fmt.Printf("%s  -  %d\n", pair.Key, pair.Value) }
}

func (wfc wordFrequencyController) Dispatch(operation string, args ...string) []generic {
	if operation == "run" {
		wfc.run()
		return []generic{}
	}
	panic(fmt.Sprintf("Unexpected operation '%s' for 'wordFrequencyController'", operation))
}

func NewWordFrequencyController(pathToFile, pathToStopWords string) *wordFrequencyController {
	return &wordFrequencyController{
		NewDataStorageManager(pathToFile),
		NewStopWordManager(pathToStopWords),
		NewWordsFrequencyManager(),
	}
}
