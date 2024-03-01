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
	cellByPoint := make(map[string]int)
	for _, pos := range u.cellPositions {
		cellByPoint[pos.String()] = 1
	}
	pointsToCheckUniq := make(map[string]bool)
	pointsToCheck := make([]Point, 0)
	for _, pos := range u.cellPositions {
		newPoints := append([]Point{pos}, pos.Neighbors()...)

		for _, nPoint := range newPoints {
			if pointsToCheckUniq[nPoint.String()] {
				continue
			}
			pointsToCheckUniq[nPoint.String()] = true
			pointsToCheck = append(pointsToCheck, nPoint)
		}
	}
	newCellPosition := make([]Point, 0)
	for _, point := range pointsToCheck {
		neighbors := point.Neighbors()

		neighborCellsCount := 0
		for _, nPoint := range neighbors {
			neighborCellsCount += cellByPoint[nPoint.String()]
		}

		if (cellByPoint[point.String()] == 1 && (neighborCellsCount == 2 || neighborCellsCount == 3)) ||
			(cellByPoint[point.String()] == 0 && neighborCellsCount == 3) {
			newCellPosition = append(newCellPosition, point)
		}
	}

	return NewUniverse(newCellPosition)
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
