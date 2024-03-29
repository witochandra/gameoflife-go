package main

import (
	"fmt"
	"gameoflife/gameoflife"
	"gameoflife/util"
	"log"
	"os"
	"strconv"
)

func args() (inputPath string, steps int, outputPath string, err error) {
	inputPath = os.Args[1]
	if len(inputPath) == 0 {
		err = fmt.Errorf("input path is required")
		return
	}
	stepsString := os.Args[2]
	steps, e := strconv.Atoi(stepsString)
	if e != nil {
		err = fmt.Errorf("steps must be a positive integer")
		return
	}
	outputPath = "./output.txt"
	if len(os.Args) > 3 {
		if len(os.Args[3]) > 0 {
			outputPath = os.Args[3]
		}
	}
	return
}

func main() {
	inputPath, steps, outputPath, err := args()
	if err != nil {
		log.Fatal(err)
		return
	}

	state := util.ReadStateFromFile(inputPath)
	newState := gameoflife.Next(state, steps)
	util.WriteStateToFile(outputPath, newState)
}
