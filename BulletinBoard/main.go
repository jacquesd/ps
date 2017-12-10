package main

import (
	"os"
	bb "./bulletinboard"
)

func main() {
	manager := bb.NewEventManager()
	bb.NewDataStroage(manager)
	bb.NewStopWordFilter(manager)
	bb.NewWordFrequencyCounter(manager)
	bb.NewWordFreuqencyApplication(manager)
	manager.Publish(bb.Event{Name: "run", Value: []string{os.Args[1], os.Args[2]}})
}
