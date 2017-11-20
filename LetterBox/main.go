package main

import (
	"os"
	"./letterbox"
)

func main() {
	var wfc letterbox.LetterBox
	wfc = letterbox.NewWordFrequencyController(os.Args[1], os.Args[2])
	wfc.Dispatch("run")
}

