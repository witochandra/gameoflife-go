package gameoflife

import "fmt"

type Universe struct {
	cellPositions [][]int
}

func (u Universe) CellPositions() [][]int {
	return u.cellPositions
}

func NewUniverse(positions [][]int) Universe {
	// removes duplicate cells in the input
	cellPositionsUniq := make(map[string]bool)
	cellPositions := make([][]int, 0)
	for _, pos := range positions {
		coordinateStr := fmt.Sprintf("%d,%d", pos[0], pos[1])
		if _, ok := cellPositionsUniq[coordinateStr]; ok {
			continue
		}
		cellPositionsUniq[coordinateStr] = true
		cellPositions = append(cellPositions, pos)
	}

	return Universe{
		cellPositions: cellPositions,
	}
}
