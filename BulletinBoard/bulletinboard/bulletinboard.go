package bulletinboard


import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)


type Event struct { Name string; Value []string}

type EventHandler func(args ...string)

type eventHandler struct { manager eventManager }

func (self *eventHandler) publish(event Event) { self.manager.Publish(event) }

type eventManager struct {
	subscriptions map[string][]EventHandler
}

func (self *eventManager) Subscribe(eventType string, handler EventHandler) {
	self.subscriptions[eventType] = append(self.subscriptions[eventType], handler)
}

func (self *eventManager) Publish(event Event) {
	for _, handler := range self.subscriptions[event.Name] { handler(event.Value...) }
}

func NewEventManager() *eventManager {
	return &eventManager{subscriptions:map[string][]EventHandler{}}
}

type dataStorage struct { eventHandler; data string }

func (self *dataStorage) load (args ...string){
	fd, err := ioutil.ReadFile(args[0])
	if err != nil { panic(err) }
	self.data = string(fd)
	pattern := regexp.MustCompile(`[\W_]+`)
	self.data = strings.ToLower(pattern.ReplaceAllString(self.data, ` `))
}

func (self *dataStorage) produceWords(_ ...string) {
	for _, word := range strings.Fields(self.data) {
		self.publish(Event{Name: "word", Value: []string{word}})
	}
	self.publish(Event{Name: "eof"})
}

func NewDataStroage(manager *eventManager) *dataStorage {
	self := dataStorage{}
	self.manager = *manager
	self.manager = *manager
	manager.Subscribe("load", self.load)
	manager.Subscribe("start", self.produceWords)
	return &self
}

type stopWordFilter struct { eventHandler; stopWords []string }

func (self *stopWordFilter) load(args ...string) {
	f, err := ioutil.ReadFile(args[1])
	if err != nil { panic(err) }
	self.stopWords = strings.Split(string(f), `,`)
	self.stopWords = append(self.stopWords, strings.Split(asciiLowercase, "")...)
}

func (self *stopWordFilter) isStopWord(args ...string) {
	if !stringInSlice(args[0], self.stopWords) {
		self.publish(Event{Name: "valid_word", Value: args})
	}
}

func NewStopWordFilter(manager *eventManager) *stopWordFilter {
	self := stopWordFilter{}
	self.manager = *manager
	manager.Subscribe("load", self.load)
	manager.Subscribe("word", self.isStopWord)
	return &self
}

type wordFrequencyCounter struct { eventHandler; wordFreqs map[string]int }

func (self *wordFrequencyCounter) increaseCount(args ...string) { self.wordFreqs[args[0]] += 1 }

func (self *wordFrequencyCounter) print(_ ...string) {
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

func NewWordFrequencyCounter(manager *eventManager) *wordFrequencyCounter {
	self := wordFrequencyCounter{wordFreqs: map[string]int{}}
	self.manager = *manager
	manager.Subscribe("valid_word", self.increaseCount)
	manager.Subscribe("print", self.print)
	return &self
}

type wordFrequencyApplication struct {eventHandler}

func (self *wordFrequencyApplication) run(args ...string) {
	self.publish(Event{Name: "load", Value: args})
	self.publish(Event{Name: "start"})
}

func (self *wordFrequencyApplication) stop(_ ...string) { self.publish(Event{Name: "print"}) }

func NewWordFreuqencyApplication(manager *eventManager) *wordFrequencyApplication {
	self := wordFrequencyApplication{}
	self.manager = *manager
	manager.Subscribe("eof", self.stop)
	manager.Subscribe("run", self.run)
	return &self
}