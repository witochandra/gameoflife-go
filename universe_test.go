package gameoflife

import "testing"

func TestNewUniverse(t *testing.T) {
	positions := [][]int{
		{1, 2},
		{1, 2},
		{2, 3},
	}
	u := NewUniverse(positions)
	if len(u.CellPositions()) != 2 {
		t.Errorf("Expected 2 cells, got %d", len(u.CellPositions()))
	}
}
