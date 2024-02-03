package gameoflife

import "fmt"

type Universe struct {
	cellPositions [][]int
}

func (u Universe) CellPositions() [][]int {
	return u.cellPositions
}

/*
Any live cell with fewer than two live neighbors dies, as if by underpopulation.
Any live cell with two or three live neighbors lives on to the next generation.
Any live cell with more than three live neighbors dies, as if by overpopulation.
Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
*/
func (u Universe) Next() Universe {
	cellByCoordinate := make(map[string]int)
	for _, pos := range u.cellPositions {
		coordinateStr := fmt.Sprintf("%d,%d", pos[0], pos[1])
		cellByCoordinate[coordinateStr] = 1
	}
	coordinatesToCheckUniq := make(map[string]bool)
	coordinatesToCheck := make([][]int, 0)
	for _, pos := range u.cellPositions {
		newCoordinates := [][]int{
			pos,
			{pos[0], pos[0] + 1},
			{pos[0] + 1, pos[0] + 1},
			{pos[0] + 1, pos[0]},
			{pos[0] + 1, pos[0] - 1},
			{pos[0], pos[0] - 1},
			{pos[0] - 1, pos[0] - 1},
			{pos[0] - 1, pos[0]},
			{pos[0] - 1, pos[0] + 1},
		}

		for _, nCoordinate := range neighborCoordinates {
			neighborsCount += cellByCoordinate[fmt.Sprintf("%d,%d", nCoordinate[0], nCoordinate[1])]
		}
		if neighborsCount == 3 {

		}
	}
	newCellCoordinates := make([][]int, 0)

	return NewUniverse(newCellCoordinates)
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
