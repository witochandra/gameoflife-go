package gameoflife

type Universe struct {
	cellPositions []Point
}

func (u Universe) CellPositions() []Point {
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
		cellByCoordinate[pos.String()] = 1
	}
	coordinatesToCheckUniq := make(map[string]bool)
	coordinatesToCheck := make([]Point, 0)
	for _, pos := range u.cellPositions {
		newCoordinates := append([]Point{pos}, pos.Neighbors()...)

		for _, nCoordinate := range newCoordinates {
			if coordinatesToCheckUniq[nCoordinate.String()] {
				continue
			}
			coordinatesToCheckUniq[nCoordinate.String()] = true
			coordinatesToCheck = append(coordinatesToCheck, nCoordinate)
		}
	}
	newCellCoordinates := make([]Point, 0)
	for _, coordinate := range coordinatesToCheck {
		neighbors := coordinate.Neighbors()

		neighborCellsCount := 0
		for _, nCoordinate := range neighbors {
			neighborCellsCount += cellByCoordinate[nCoordinate.String()]
		}

		if (cellByCoordinate[coordinate.String()] == 1 && (neighborCellsCount == 2 || neighborCellsCount == 3)) ||
			(cellByCoordinate[coordinate.String()] == 0 && neighborCellsCount == 3) {
			newCellCoordinates = append(newCellCoordinates, coordinate)
		}
	}

	return NewUniverse(newCellCoordinates)
}

func NewUniverse(positions []Point) Universe {
	// removes duplicate cells in the input
	cellPositionsUniq := make(map[string]bool)
	cellPositions := make([]Point, 0)
	for _, pos := range positions {
		if _, ok := cellPositionsUniq[pos.String()]; ok {
			continue
		}
		cellPositionsUniq[pos.String()] = true
		cellPositions = append(cellPositions, pos)
	}

	return Universe{
		cellPositions: cellPositions,
	}
}
