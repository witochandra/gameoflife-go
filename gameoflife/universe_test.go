package gameoflife

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUniverse(t *testing.T) {
	u := NewUniverse([]Point{})
	assert.Equal(t, len(u.CellPositions()), 0)

	u = NewUniverse([]Point{
		{1, 2},
		{1, 2},
		{2, 3},
	})
	assert.Equal(t, len(u.CellPositions()), 2)

	u = NewUniverse([]Point{
		{1, 2},
		{2, 3},
	})
	assert.Equal(t, len(u.CellPositions()), 2)
}

func TestUniverse_Next(t *testing.T) {
	t.Run("Any live cell with fewer than two live neighbors dies, as if by underpopulation.", func(t *testing.T) {
		u := NewUniverse([]Point{
			{0, 0},
		})

		assert.Empty(t, u.Next().CellPositions())

		u = NewUniverse([]Point{
			{0, 0}, {0, 1},
		})

		assert.Empty(t, u.Next().CellPositions())
	})

	t.Run("Any live cell with more than three live neighbors dies, as if by overpopulation.", func(t *testing.T) {
		/*
			x x      x
			xx   ->  x
			 xx      xxx
		*/
		u := NewUniverse([]Point{
			{0, 0}, {2, 0},
			{0, -1}, {1, -1},
			{1, -2}, {2, -2},
		})
		assert.ElementsMatch(t,
			[]Point{
				{0, 0},
				{0, -1},
				{0, -2}, {1, -2}, {2, -2},
			},
			u.Next().CellPositions(),
		)
	})
}
