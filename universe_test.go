package gameoflife

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUniverse(t *testing.T) {
	u := NewUniverse([][]int{})
	assert.Equal(t, len(u.CellPositions()), 0)

	u = NewUniverse([][]int{
		{1, 2},
		{1, 2},
		{2, 3},
	})
	assert.Equal(t, len(u.CellPositions()), 2)

	u = NewUniverse([][]int{
		{1, 2},
		{2, 3},
	})
	assert.Equal(t, len(u.CellPositions()), 2)
}

func TestNeighborsOf(t *testing.T) {
	neighbors := NeighborsOf([]int{0, 0})
	assert.ElementsMatch(t, [][]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}, neighbors)
}
