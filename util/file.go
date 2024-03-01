package util

import (
	"bufio"
	"gameoflife/gameoflife"
	"log"
	"os"
)

func ReadStateFromFile(path string) []gameoflife.Point {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cells := make([]gameoflife.Point, 0)
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		for x, c := range scanner.Text() {
			if c != ' ' {
				cells = append(cells, gameoflife.Point{X: x, Y: y})
			}
		}
		y += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return cells
}

func WriteStateToFile(path string, positions []gameoflife.Point) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maxX, maxY := 0, 0
	cellByCoordinate := make(map[string]int)
	for _, pos := range positions {
		cellByCoordinate[pos.String()] = 1
		if pos.X > maxX {
			maxX = pos.X
		}
		if pos.Y > maxY {
			maxY = pos.Y
		}
	}

	buffer := ""
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			pos := gameoflife.Point{X: x, Y: y}
			if cellByCoordinate[pos.String()] == 1 {
				buffer += "O"
			} else {
				buffer += " "
			}
		}
		buffer += "\n"
	}
	file.WriteString(buffer)
}
