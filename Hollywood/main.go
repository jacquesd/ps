package main

import (
	"./hollywood"
	"os"
)

func main() {
	wfapp := hollywood.NewWordFrequencyFrameWork()
	stopWordFilter := hollywood.NewStopWordManager(wfapp)
	dataStorage := hollywood.NewDataStorageManager(wfapp, stopWordFilter)
	hollywood.NewWordsFrequencyCounter(wfapp, dataStorage)
	wfapp.Run(os.Args[1], os.Args[2])
}
