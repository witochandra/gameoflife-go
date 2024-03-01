package main

import (
	"flag"
	"fmt"
	"gameoflife/gameoflife"
	"gameoflife/util"
	"time"

	"golang.org/x/term"
)

func renderCells(positions []gameoflife.Point, width, height int) {
	cellByCoordinate := make(map[string]int)
	for _, pos := range positions {
		cellByCoordinate[pos.String()] = 1
	}
	buffer := ""
	for y := 0; y < height-2; y++ {
		for x := 0; x < width; x++ {
			pos := gameoflife.Point{X: x, Y: y}
			if cellByCoordinate[pos.String()] == 1 {
				buffer += "O"
			} else {
				buffer += " "
			}
		}
		buffer += "\n"
	}
	fmt.Print(buffer)
}

func main() {
	interval := 100 * time.Millisecond

	inputPath := flag.String("input", "examples/input.txt", "Path to input file")
	outputPath := flag.String("output", "", "Path to output file")
	steps := flag.Int("steps", 2000000000, "Number of steps")
	render := flag.Bool("render", true, "Render to terminal")

	flag.Parse()

	state := util.ReadStateFromFile(*inputPath)

	counter := 0
	for counter < *steps {
		if *render {
			width, height, err := term.GetSize(0)
			if err != nil {
				return
			}
			fmt.Print("\033[H\033[?25l") // move cursor to top left & hide cursor
			fmt.Printf("W: %v, H: %v, Cycle No: %v\n", width, height, counter)
			renderCells(state, width, height)
			time.Sleep(interval)
		}
		counter++

		state = gameoflife.Next(state)
	}
	if *outputPath != "" {
		util.WriteStateToFile(*outputPath, state)
	}
}
