package util

import (
	"gameoflife/gameoflife"
	"log"
	"os"
)

const (
	DeadCell  = "_"
	AliveCell = "#"
)

func ReadStateFromFile(path string) gameoflife.Points {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return gameoflife.NewPoints(string(content), AliveCell)
}

func WriteStateToFile(path string, cells gameoflife.Points) {
	str := cells.String(2, DeadCell, AliveCell)
	err := os.WriteFile(path, []byte(str), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
