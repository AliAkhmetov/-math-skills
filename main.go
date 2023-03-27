package main

import (
	"math-skills/tools"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	fileName := args[0]
	tools.ReadData(fileName)
}
