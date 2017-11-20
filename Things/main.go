package main

import (
	"os"
	"./things"
)

func main() { things.NewWordFrequencyController(os.Args[1], os.Args[2]).Run() }
