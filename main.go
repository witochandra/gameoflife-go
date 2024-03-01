package main

import (
	"fmt"
	"gameoflife/gameoflife"
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
	state := make([]gameoflife.Point, 0)
	state = append(state, gameoflife.CreateBlinker(15, 10)...)
	state = append(state, gameoflife.CreateToad(30, 10)...)
	state = append(state, gameoflife.CreateBeacon(45, 10)...)
	state = append(state, gameoflife.CreatePulsar(60, 10)...)
	state = append(state, gameoflife.CreateGlider(0, 0)...)
	state = append(state, gameoflife.CreateGlider(15, 0)...)

	counter := 0
	for counter <= 1000 {
		width, height, err := term.GetSize(0)
		if err != nil {
			return
		}
		// move cursor to top left & hide cursor
		fmt.Print("\033[H\033[?25l")

		fmt.Printf("W: %v, H: %v, Cycle No: %v\n", width, height, counter)
		renderCells(state, width, height)

		time.Sleep(30 * time.Millisecond)
		counter++

		state = gameoflife.Next(state)
	}
}
