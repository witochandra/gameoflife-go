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

func TestUniverse_Next(t *testing.T) {
	t.Run("Any live cell with fewer than two live neighbors dies, as if by underpopulation.", func(t *testing.T) {
		u := NewUniverse([][]int{
			{0, 0},
		})

		assert.Empty(t, u.Next().CellPositions())

		u = NewUniverse([][]int{
			{0, 0}, {0, 1},
		})

		assert.Empty(t, u.Next().CellPositions())
	})

	t.Run("Any live cell with two or three live neighbors lives on to the next generation.", func(t *testing.T) {
		u := NewUniverse([][]int{
			{0, 0}, {0, 1}, {1, 1},
		})
		assert.ElementsMatch(t,
			[][]int{
				{0, 0}, {0, 1}, {1, 1},
			},
			u.Next().CellPositions(),
		)

		u = NewUniverse([][]int{
			{0, 0}, {0, 1}, {1, 1}, {1, 0},
		})
		assert.ElementsMatch(t,
			[][]int{
				{0, 0}, {0, 1}, {1, 1}, {1, 0},
			},
			u.Next().CellPositions(),
		)
	})
}
